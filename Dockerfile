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
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
#RUN chmod +x main
