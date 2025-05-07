# PROJECT COMMAND RUNNER

# Default command
default:
    @just --list

# Run the CLI
muxingbird *args:
    @go run . {{ args }}

# Execute unit tests
test:
    @echo "Running unit tests!"
    go clean -testcache
    go test -cover ./...

# Build Docker image
build:
	@echo "Building Docker image..."
	docker build -t muxingbird:latest -f .docker/api.Dockerfile .
	@echo "Docker image built successfully!"

# Sync Go modules
tidy:
    go mod tidy
    go work sync
