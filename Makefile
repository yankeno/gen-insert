BIN := gen-insert

build:
	@echo "Building the project..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN) .
	@echo "Build completed!"
