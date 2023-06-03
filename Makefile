.PHONY: all build run go tool clean help

BINARY="bluebell"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/${BINARY}


run:
	@go run ./main.go config/config.yaml


gotool:
	go fmt ./
	go vet ./

clean:

help: