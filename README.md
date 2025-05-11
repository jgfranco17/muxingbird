# Muxingbird

Muxingbird is a lightweight Go-based CLI tool that spins up mock HTTP servers from a JSON spec
file. Designed for testing, prototyping, and mocking RESTful APIs, Muxingbird lets you simulate
real server responses in seconds.

Inspired by the concepts of "mux" (HTTP routing) and the mockingbird, Muxingbird sings back
exactly what you want your mock API to say.

## Local Development

These instructions will get you a copy of the project up and running on your local machine
for development and testing purposes.

### Pre-requisites

- [Go 1.23](https://go.dev/doc/install) or above
- [Just](https://github.com/casey/just) command runner (optional)

### Repository setup

To get started with this project, clone the repository to your local machine and install the
required dependencies.

```bash
git clone https://github.com/jgfranco17/muxingbird.git
cd muxingbird
go mod tidy
```

The project uses `just` as our command runner. You can get a list of available recipes by
executing `just` with no arguments. While not required, the scripts provided will largely
help in quickly and reproducibly setting up your development environment.

### Build with Go

To run the CLI as a locally-built executable, we can build with Go. For convenience, a Just
script has been provided to streamline the flags setup.

```bash
just build
```

### Build with Docker

To run the CLI in a container, the package comes with both a Dockerfile and a Compose YAML
configuration. Run either of the following to get the API launched in a container; by default,
the generated servers will be set to listen on port `8000` for the Compose.

```bash
docker compose build --no-cache
docker compose run --rm --remove-orphans muxingbird-cli <cli-args>
```

### Testing

In order to run diagnostics and unit tests, first install the testing dependencies. We use the
[`testify/assert`](https://github.com/stretchr/testify) library for our tests. To run the full
test suite, we can simply invoke the `go test` command for the whole module.

```bash
go test -cover ./...
```

## Installation

> [!WARNING]
> This CLI is still an alpha prototype.

To download the CLI binary, an install script has been provided. This will install the binary
in your `$HOME/.local/bin`. Assuming that is on your `PATH`, you should be able to directly
access the CLI straight away.

```bash
wget -O - https://raw.githubusercontent.com/jgfranco17/muxingbird/main/install.sh | bash
```

They always say not to just blindly run scripts from the internet, so feel free to examine
the [file](https://github.com/jgfranco17/muxingbird/blob/main/install.sh) first before
running it.

To verify the installation, you can call the CLI `help` flag for a basic usage guide.

```console
$ muxingbird --help
Muxingbird: spin up HTTP servers with a few clicks

Usage:
  muxingbird [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  run         Run the server from the config

Flags:
  -h, --help            help for muxingbird
  -v, --verbose count   Increase verbosity (-v or -vv)
      --version         version for muxingbird

Use "muxingbird [command] --help" for more information about a command.
```

## License

This project is licensed under the BSD-3 License. See the LICENSE file for more details.
