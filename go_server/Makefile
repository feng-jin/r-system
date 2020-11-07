all: install

install:
	go build -o bin/api_server ./src/api_server
	go mod tidy
	go mod vendor