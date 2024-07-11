This document is available on English ---> [English documentation](README.md)

# gRPC Gateway Example

Этот проект демонстрирует использование gRPC и gRPC-Gateway для создания микросервисного приложения, которое управляет настройками хоста и DNS-серверами. 
Приложение включает серверную часть, клиентскую часть CLI и REST API. 
Проект был написан в качестве тестового задания для YADRO.

## Установка и запуск

1. Установите Go: https://golang.org/doc/install
2. Установите protoc: https://grpc.io/docs/protoc-installation/
3. Установите необходимые плагины для protoc:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.10.0
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.10.0
```

## Компиляция и запуск

1. Сгенерируйте protobuf файлы и скомпилируйте проект:
```sh
make all
```
2. Запустите сервер в одном терминале:
```sh
./bin/server
```
3. Выполняйте команды через клиента:
```sh
./bin/client <command> <flag> <arg>
```
(Примеры команд приведены ниже)

## Использование

### CLI client

CLI поддерживает 4 команды:

1. `set-hostname` -> устанавливает новое имя хоста;
2. `get-dns-servers` -> получает список dns-серверов;
3. `add-dns-server` -> добавляет новый dns-сервер;
4. `remove-dns-server` -> удаляет указанный dns-сервер.

P. S. 
Чтобы была возможность добавить сервер или удалить, нужно, к примеру, изменить права доступа к файлу:
```sh
sudo chmod a+w /etc/resolv.conf
```

После чего можно будет вернуть права доступа прежние:
```sh
sudo chmod a-w /etc/resolv.conf
```

### Примеры использования CLI

```sh
# Установить имя хоста
./bin/client set-hostname --hostname=my-new-host

# Получить список DNS-серверов
./bin/client get-dns-servers

# Добавить DNS-сервер
./bin/client add-dns-server --server=8.8.8.8

# Удалить DNS-сервер
./bin/client remove-dns-server --server=8.8.8.8
```

### REST API

REST API доступен на порту `8080`. Примеры запросов через терминал:

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

Файл `host.proto` описывает gRPC API:

1. `SetHostname(SetHostnameRequest) returns (SetHostnameResponse)`
2. `GetDNSServers(GetDNSServersRequest) returns (GetDNSServersResponse)`
3. `AddDNSServer(AddDNSServerRequest) returns (AddDNSServerResponse)`
4. `RemoveDNSServer(RemoveDNSServerRequest) returns (RemoveDNSServerResponse)`

### REST API

gRPC-Gateway позволяет использовать те же методы через REST:

1. `POST /v1/hostname`
2. `GET /v1/dns-servers`
3. `POST /v1/dns-servers`
4. `DELETE /v1/dns-servers/{server}`

### Makefile команды

1. `make all`: Сгенерировать protobuf файлы и скомпилировать проект
2. `make clean`: Очистить сгенерированные файлы
3. `make proto`: Сгенерировать protobuf файлы
4. `make build`: Скомпилировать сервер и клиент
5. `make server`: Запустить сервер
6. `make client`: Запустить клиент

## Итог
Этот README_ru.md файл описывает структуру проекта, процесс установки и запуска, а также основные команды и методы API. 
Он предназначен для того, чтобы помочь вам и другим разработчикам быстро понять и начать использовать проект.
