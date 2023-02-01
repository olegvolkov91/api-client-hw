.SILENT:

BINARY_NAME=apiclient

.PHONY:build
build:
	go build -o ./.bin/${BINARY_NAME} cmd/main.go

.PHONY: run
run: build
	./.bin/${BINARY_NAME}