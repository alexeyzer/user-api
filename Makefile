DST_DIR = ./internal
SRC_DIR = ./api

generate:
	mkdir -p "pb"
	protoc -I/usr/local/include -I. \
    			-I${GOPATH}/src \
    			-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    			-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
    			--grpc-gateway_out=logtostderr=true:./api_pb \
    			--swagger_out=allow_merge=true,merge_file_name=api:. \
    			--go_out=plugins=grpc:./pb ./api/user-api/v1/*.proto
lint:
	golangci-lint run