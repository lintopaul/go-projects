.DEFAULT_GOAL := build
BIN_FILE=bookstore
build:
	GOOS=darwin GOARCH=arm64 go build -o ${BIN_FILE} ./cmd/...
buildlinux:
	GOOS=linux GOARCH=arm64 go build -o ${BIN_FILE} ./cmd/...	
buildmac:
	GOOS=darwin GOARCH=amd64 go build -o ${BIN_FILE} ./cmd/...
buildmacm1:
	GOOS=darwin GOARCH=arm64 go build -o ${BIN_FILE} ./cmd/...
clean:
	go clean
test:
	go test -v ./cmd/... ./pkg/...
cover:
	go test -coverprofile cover.out ./cmd/... ./pkg/...
	go tool cover -html=cover.out
run:
	./${BIN_FILE}
lint:
	golangci-lint run --enable-all