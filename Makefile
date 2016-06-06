.PHONY: build build-arm run

build:
	env GO15VENDOREXPERIMENT=1 go build

build-arm:
	env GO15VENDOREXPERIMENT=1 GOOS=linux GOARCH=arm go build

run:
	go run
