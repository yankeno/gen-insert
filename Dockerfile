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

COPY . .

RUN go mod download
