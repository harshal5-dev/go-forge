.DEFAULT_GOAL := run

.PHONY:fmt vet build run clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

run: vet
	go run hello.go

clean:
	go clean
