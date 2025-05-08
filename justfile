# PROJECT COMMAND RUNNER

PROJECT_NAME := "muxingbird"
INSTALL_PATH := "${HOME}/.local/bin"

# Default command
_default:
    @just --list --unsorted

# Sync Go modules
tidy:
    go mod tidy
    go work sync

# Run the CLI in dev mode
muxingbird *args:
    @go run . {{ args }}

# Execute unit tests
test:
    @echo "Running unit tests!"
    go clean -testcache
    go test -cover ./...

# Build a local binary
build:
    #!/usr/bin/env bash
    echo "Building {{ PROJECT_NAME }} binary..."
    go mod download all
    VERSION=$(jq -r .version specs.json)
    CGO_ENABLED=0 GOOS=linux go build \
        -ldflags="-X main.version=${VERSION}"\
        -o ./{{ PROJECT_NAME }} .
    chmod +x ./{{ PROJECT_NAME }}
    echo "Built binary for {{ PROJECT_NAME }} ${VERSION} successfully!"

# Install the CLI locally in bin
install-local: build
    #!/usr/bin/env bash
    echo "Installing local binary at {{ INSTALL_PATH }}"
    cp ./{{ PROJECT_NAME }} {{ INSTALL_PATH }}
    echo $PATH | grep -q {{ INSTALL_PATH }} || exit 1
    echo "Installed Muxingbird in local: {{ INSTALL_PATH }}/{{ PROJECT_NAME }}"

# Build the Docker image
docker-build:
    docker compose -f docker-compose.yaml build --no-cache

# Run the CLI in Docker
docker-run *args:
    docker compose \
        -f docker-compose.yaml \
        run --rm --remove-orphans muxingbird-cli {{ args }}
