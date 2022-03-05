# todo

## Overview

This is a sample application for REST service development with golang. Libraries used:

* [gorm](https://gorm.io/) as ORM library
* [viper](https://github.com/spf13/viper) for configuration management
* [gin](https://github.com/gin-gonic/gin) as web framework
* [gin-swagger](https://github.com/swaggo/gin-swagger) to generate OpenAPI spec from go comments

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

### Set up the database

Also see: <https://hub.docker.com/_/postgres>

`docker run -it --rm -p 5432:5432 --name pg -e POSTGRES_PASSWORD=password postgres`

```shell
$ docker exec -it pg /bin/bash                               
root@187961c81d2e:/# psql -U postgres
psql (14.2 (Debian 14.2-1.pgdg110+1))
Type "help" for help.
```

### Start the server

`go run main.go`

Go to Swagger UI <http://localhost:8080/swagger/index.html>

## Test the service

`go test ./...`

## Generate OpenAPI spec

`swag init`

For more see: <https://github.com/swaggo/gin-swagger>

