FROM golang:1.20-alpine

ENV GO111MODULE=on

WORKDIR /go/gen-insert

RUN apk update && apk add --no-cache \
    bash \
    file \
    curl \
    wget \
    vim \
    make

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

ENV PATH="/go/bin:${PATH}"

COPY . .

RUN go mod download
