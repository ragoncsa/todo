# todo

## Overview

This is a sample application for REST service development with golang. Libraries used:

* [gorm](https://gorm.io/) as ORM library
* [viper](https://github.com/spf13/viper) for configuration management
* [gin](https://github.com/gin-gonic/gin) as web framework
* [gin-swagger](https://github.com/swaggo/gin-swagger) to generate OpenAPI spec from go comments

Simple REST webservice 

## Set up database

Also see: <https://hub.docker.com/_/postgres>

`docker run -it --rm -p 5432:5432 --name pg -e POSTGRES_PASSWORD=password postgres`

```shell
$ docker exec -it pg /bin/bash                               
root@187961c81d2e:/# psql -U postgres
psql (14.2 (Debian 14.2-1.pgdg110+1))
Type "help" for help.
```

## Start the server

go run main.go

## Test the service

### Create a task

`curl -X POST -H 'Content-Type: application/json' localhost:8080/tasks/ -d '{"name": "hello1"}'`

`http POST localhost:8080/tasks/ name=hello1`

### Get a task

`curl -X GET localhost:8080/tasks/1`

`http GET localhost:8080/tasks/1`

### Get all tasks

`curl -X GET localhost:8080/tasks/`

`http GET localhost:8080/tasks/`

### Delete a task

`curl -X DELETE localhost:8080/tasks/1`

`http DELETE localhost:8080/tasks/1`

### Delete all tasks

`curl -X DELETE localhost:8080/tasks/`

`http DELETE localhost:8080/tasks/`

### Run tests

go test ./...

## Generate OpenAPI spec

`swag init`

To access swagger UI go to <http://localhost:8080/swagger/index.html>

For more see: <https://github.com/swaggo/gin-swagger>