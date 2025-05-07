# syntax=docker/dockerfile:1.4

# ========== CLI BUILD STAGE ==========

ARG GO_VERSION=1.23
FROM golang:${GO_VERSION} AS build

COPY . /app
WORKDIR /app

ARG VERSION=0.0.0
RUN go mod download all && \
    CGO_ENABLED=0 GOOS=linux go build \
        -ldflags="-X main.version=${VERSION}" \
        -o ./muxingbird .

# ========== APP RUN STAGE ==========

FROM alpine:3.21 AS app

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add ca-certificates tzdata \
    && update-ca-certificates

COPY --from=build /app/muxingbird /muxingbird

ARG PORT=8000
EXPOSE ${PORT}

ENTRYPOINT [ "/muxingbird" ]
CMD [ "--help" ]
