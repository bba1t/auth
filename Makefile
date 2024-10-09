LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get github.com/golang/protobuf/ptypes/wrappers
	go get google.golang.org/protobuf/types/known/timestamppb
	go get github.com/golang/protobuf/ptypes/empty
	go get google.golang.org/protobuf/runtime/protoimpl
	go get google.golang.org/protobuf/reflect/protoreflect
	go get google.golang.org/grpc/status
	go get google.golang.org/grpc
	go get google.golang.org/grpc/codes

generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=./bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=./bin/protoc-gen-go-grpc \
	api/user_v1/user.proto