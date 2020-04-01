# Variable
ARG GO_VERSION=1.14.1
ARG ALPINE_VERSION=3.11

# Build Stage
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS build-stage
ENV CGO_ENABLED 0

WORKDIR /go/go_dev
COPY . .
RUN set -ex && \
    apk update && \
    apk add --no-cache git && \
    go get golang.org/x/tools/cmd/goimports && \
    go get golang.org/x/lint/golint
    # go get gopkg.in/urfave/cli.v2@master && \
    # go get github.com/oxequa/realize && \
    # go get github.com/go-delve/delve/cmd/dlv && \
    # go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv && \
    # go build -o /app/go_dev
