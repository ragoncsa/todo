# todo

## Overview

This is a sample application for REST service development with golang. Libraries used:

* [gorm](https://gorm.io/) as ORM library
* [viper](https://github.com/spf13/viper) for configuration management
* [gin](https://github.com/gin-gonic/gin) as web framework
* [gin-swagger](https://github.com/swaggo/gin-swagger) to generate OpenAPI spec from go comments
* [resty](https://github.com/go-resty/resty) for REST client implementation (for example to talk to OPA)
* [Open Policy Agent (OPA)](https://www.openpolicyagent.org/) for authorization decisions

![Alt text](assets/screenshot-swagger-ui.png?raw=true "Screenshot")

## Run from container

To build

`docker build -t todo .`

Docker-compose starts the built container with a database

`docker-compose up`

Go to Swagger UI <http://localhost:8080/swagger/index.html>


### Reset the database

`docker-compose down --volumes`

## Run without container

### Start dependencies

```shell
docker-compose up db
opa build authz -o authz/bundle.tar.gz
docker-compose up bundle_server
docker-compose up opa
```

To access the database:

```shell
$ docker exec -it todo_db_1 /bin/bash                               
root@187961c81d2e:/# psql -U postgres
psql (14.2 (Debian 14.2-1.pgdg110+1))
Type "help" for help.
```

### Start the server

`go run main.go`

Go to Swagger UI <http://localhost:8080/swagger/index.html>

## Testing

### Test the application

`go test ./...`

### Test the authorization rules

Run unit tests

`opa test authz -v --ignore '*.tar.gz'`

Test rules on the server

```shell
echo "{\"input\": {\"method\":\"POST\",\"owner\":\"johndoe\",\"path\":[\"tasks\"],\"user\":\"johndoe\"}}" \
| http -v POST http://127.0.0.1:8181/v1/data/authz
```

## Generate OpenAPI spec

`swag init`

For more see: <https://github.com/swaggo/gin-swagger>

