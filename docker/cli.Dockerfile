# syntax=docker/dockerfile:1.4

# ========== CLI BUILD STAGE ==========

ARG GO_VERSION=1.23
FROM golang:${GO_VERSION} AS build

COPY . /src
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETOS=linux
ARG VERSION=0.0.0
RUN --mount=type=cache,target=/go/pkg/mod/ \
    CGO_ENABLED=0 GOOS=${TARGETOS} go build \
        -ldflags="-X main.version=${VERSION}" \
        -o ./muxingbird .

# ========== APP RUN STAGE ==========

FROM alpine:3.21 AS app

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add ca-certificates tzdata && \
    update-ca-certificates

COPY --from=build /src/muxingbird /muxingbird

ARG PORT=8000
EXPOSE ${PORT}

LABEL title="Muxingbird CLI"
LABEL description="Mock HTTP server for route simulation"
LABEL author="Joaquin Franco"
LABEL source="https://github.com/jgfranco17/muxingbird"
LABEL licenses="BSD 3-Clause"

ENTRYPOINT [ "/muxingbird" ]
CMD [ "--help" ]
