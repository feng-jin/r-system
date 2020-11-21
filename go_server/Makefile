all: build p

build:
	go build -o bin/api_server ./src/api_server
	go mod tidy
	go mod vendor

p:
	rm -rf src/lib/proto
	protoc -I proto --go_out=plugins=grpc:. greet.proto
	protoc -I proto --go_out=plugins=grpc:. errmsg.proto
	protoc -I proto --go_out=plugins=grpc:. api.proto
	protoc -I proto --go_out=plugins=grpc:. recall.proto
	protoc -I proto --go_out=plugins=grpc:. rank.proto
	protoc -I proto --go_out=plugins=grpc:. filter.proto
	cp -r go_server/src/lib/proto src/lib/
	rm -rf go_server