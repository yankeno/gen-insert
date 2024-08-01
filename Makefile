BIN := gen-insert

.PHONY: build
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN) .

.PHONY: test
test: build
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run
