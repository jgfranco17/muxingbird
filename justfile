# PROJECT COMMAND RUNNER

PROJECT_NAME := "muxingbird"
INSTALL_PATH := "${HOME}/.local/bin"
VERSION := "$(jq -r .version specs.json)"

# Default command
_default:
    @just --list --unsorted

# ========== WORKSPACE SETUP ==========

# Sync Go modules
tidy:
    go mod tidy

# ========== CLI DEV EXECUTION ==========

# Run the CLI in dev mode
muxingbird *args:
    @go run . {{ args }}

# Execute unit tests
test:
    @echo "Running unit tests!"
    go clean -testcache
    go test -cover ./...

# ========== DEPLOYMENT & BUILD ==========

# Generate a SBOM file for the project
sbom:
    #!/usr/bin/env bash
    mkdir -p sbom
    SBOM_PATH="sbom/muxingbird-sbom-{{ VERSION }}.spdx.json"
    syft dir:. -o cyclonedx-json | jq . > "${SBOM_PATH}"
    echo "Generated SBOM file: ${SBOM_PATH}"

# Build a local binary
build:
    #!/usr/bin/env bash
    echo "Building {{ PROJECT_NAME }} binary..."
    go mod download all
    CGO_ENABLED=0 GOOS=linux go build \
        -ldflags="-X main.version={{ VERSION }}"\
        -o ./{{ PROJECT_NAME }} .
    chmod +x ./{{ PROJECT_NAME }}
    echo "Built binary for {{ PROJECT_NAME }} {{ VERSION }} successfully!"

# Install the CLI locally in bin
install-local: build
    #!/usr/bin/env bash
    echo "Installing local binary at {{ INSTALL_PATH }}"
    cp ./{{ PROJECT_NAME }} {{ INSTALL_PATH }}
    echo $PATH | grep -q {{ INSTALL_PATH }} || exit 1
    echo "Installed Muxingbird in local: {{ INSTALL_PATH }}/{{ PROJECT_NAME }}"

# ========== CONTAINERIZATION ==========

# Build the Docker image
docker-build:
    docker compose build --no-cache

# Run the CLI in Docker
docker-run *args:
    @docker compose run --rm --remove-orphans muxingbird {{ args }}
