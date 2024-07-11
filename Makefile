PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc
PROTOC_GEN_GRPC_GATEWAY := $(shell go env GOPATH)/bin/protoc-gen-grpc-gateway
PROTOC_GEN_OPENAPIV2 := $(shell go env GOPATH)/bin/protoc-gen-openapiv2

.PHONY: run all clean proto build server client

all: proto build

clean:
	rm -rf api/generated

proto:
	mkdir -p api/generated
	protoc -I. \
		-I $(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway/v2) \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--plugin=protoc-gen-grpc-gateway=$(PROTOC_GEN_GRPC_GATEWAY) \
		--plugin=protoc-gen-openapiv2=$(PROTOC_GEN_OPENAPIV2) \
		--go_out=api/generated --go_opt=module=grpc-gateway/api/generated \
		--go-grpc_out=api/generated --go-grpc_opt=module=grpc-gateway/api/generated \
		--grpc-gateway_out=api/generated --grpc-gateway_opt=module=grpc-gateway/api/generated \
		--openapiv2_out=. \
		api/proto/host.proto

build:
	go build -o bin/server cmd/server/main.go
	go build -o bin/client cmd/client/cli.go

server:
	go run cmd/server/main.go

client:
	go run cmd/client/cli.go
