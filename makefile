.PHONY:all
all:clean build compile

build:
	protoc --proto_path=../ -I/usr/local/include -I. -I$(GOPATH)/pkg/mod -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --go_out=plugins=grpc:. proto/mastiff.proto
	protoc --proto_path=../ -I/usr/local/include -I. -I$(GOPATH)/pkg/mod -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --grpc-gateway_out=logtostderr=true:. proto/mastiff.proto
	protoc --proto_path=../ -I/usr/local/include -I. -I$(GOPATH)/pkg/mod -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --swagger_out=logtostderr=true:. proto/mastiff.proto

compile:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac_amd64 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64 main.go

clean:
	rm -rf ./proto/*.go ./proto/*.json ./bin