Этот документ доступен на русском языке ---> [Русская документация](README_ru.md)

# gRPC Gateway Example

This project demonstrates the use of gRPC and gRPC-Gateway to create a microservice application that manages host settings and DNS servers.
The application includes a server part, a CLI client part, and a REST API.
The project was written as a test assignment for YADRO.

## Installation

1. Install Go: https://golang.org/doc/install
2. Install protoc: https://grpc.io/docs/protoc-installation/
3. Install the necessary protoc plugins:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.10.0
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.10.0
```

## Compilation and launch

1. Generate protobuf files and compile the project:
```sh
make all
```
2. Start the server in one terminal:
```sh
./bin/server
```
3. Execute commands through the client:
```sh
./bin/client <command> <flag> <arg>
```
P.S.
(Examples of commands are provided below)

## Usage

### CLI client

The CLI supports 4 commands:

1. `set-hostname` -> sets a new hostname;
2. `get-dns-servers` -> retrieves the list of DNS servers;
3. `add-dns-server` -> adds a new DNS server;
4. `remove-dns-server` -> removes the specified DNS server.

P. S.
To add or remove a server, you may need to change file permissions, for example:
```sh
sudo chmod a+w /etc/resolv.conf
```

Then you can revert the permissions:
```sh
sudo chmod a-w /etc/resolv.conf
```

### CLI Usage Examples

```sh
# Set a hostname
./bin/client set-hostname --hostname=my-new-host

# Get DNS servers list
./bin/client get-dns-servers

# Add DNS server
./bin/client add-dns-server --server=8.8.8.8

# Remove DNS server
./bin/client remove-dns-server --server=8.8.8.8
```

### REST API

The REST API is available on port `8080`. Examples of requests through the terminal:

1. set-hostname:
```sh
curl -X POST -d '{"hostname": "my-new-host"}' http://localhost:8080/v1/hostname
```
2. get-dns-servers:
```sh
curl http://localhost:8080/v1/dns-servers
```
3. add-dns-server:
```sh
curl -X POST -d '{"server": "8.8.8.8"}' http://localhost:8080/v1/dns-servers
```
4. remove-dns-server:
```sh
curl -X DELETE http://localhost:8080/v1/dns-servers/8.8.8.8
```

## Описание API

### gRPC API

The `host.proto` file describes the gRPC API:

1. `SetHostname(SetHostnameRequest) returns (SetHostnameResponse)`
2. `GetDNSServers(GetDNSServersRequest) returns (GetDNSServersResponse)`
3. `AddDNSServer(AddDNSServerRequest) returns (AddDNSServerResponse)`
4. `RemoveDNSServer(RemoveDNSServerRequest) returns (RemoveDNSServerResponse)`

### REST API

gRPC-Gateway allows using the same methods through REST:

1. `POST /v1/hostname`
2. `GET /v1/dns-servers`
3. `POST /v1/dns-servers`
4. `DELETE /v1/dns-servers/{server}`

### Makefile commands

1. `make all`: Generate protobuf files and compile the project
2. `make clean`: Clean the generated files
3. `make proto`: Generate protobuf files
4. `make build`: Compile server and client
5. `make server`: Run the server
6. `make client`: Run the client

## Conclusion
This README file describes the project structure, installation and launch process, as well as the main commands and API methods.
It is intended to help you and other developers quickly understand and start using the project.
