PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc
PROTOC_GEN_GRPC_GATEWAY := $(shell go env GOPATH)/bin/protoc-gen-grpc-gateway
PROTOC_GEN_OPENAPIV2 := $(shell go env GOPATH)/bin/protoc-gen-openapiv2

.PHONY: all clean proto build server client

all: proto build

clean:
	rm -rf internal/api/generated

proto:
	mkdir -p internal/api/generated
	protoc -I. -I $(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway/v2) \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO) \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
		--plugin=protoc-gen-grpc-gateway=$(PROTOC_GEN_GRPC_GATEWAY) \
		--plugin=protoc-gen-openapiv2=$(PROTOC_GEN_OPENAPIV2) \
		--go_out=internal/api/generated --go_opt=paths=source_relative \
		--go-grpc_out=internal/api/generated --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=internal/api/generated --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=internal/api/generated \
		internal/api/proto/host.proto

build:
	go build -o bin/server cmd/server/main.go
	go build -o bin/client cmd/client/main.go

server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go
