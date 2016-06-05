.PHONY: build build-arm run

build:
	go build

build-arm:
	env GOOS=linux GOARCH=arm go build

run:
	go build
