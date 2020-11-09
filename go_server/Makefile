all: build protoc

build:
	go build -o bin/api_server ./src/api_server
	go mod tidy
	go mod vendor

protoc:
	protoc -I proto --go_out=plugins=grpc:. greet.proto