BINARY_NAME=rpsg
.DEFAULT_GOAL := build

build:
	GOARCH=amd64 GOOS=linux go build -o ./build/$(BINARY_NAME)-linux_amd64 main.go
	GOARCH=amd64 GOOS=darwin go build -o ./build/$(BINARY_NAME)-darwin_amd64 main.go
	GOARCH=amd64 GOOS=windows go build -o ./build/$(BINARY_NAME)-windows_amd64 main.go

test:
	go test ./test/*

install:
	go build
	cp ./rpsg /usr/local/bin/
	rm -f ./rpsg
