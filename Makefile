.PHONY: myrun

build:
	go build -o bin/cmd/main /cmd/main.go

myrun:
	go run ./...