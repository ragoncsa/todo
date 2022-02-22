# todo

### Set up database

Also see: <https://hub.docker.com/_/postgres>

`docker run -it --rm -p 5432:5432 --name pg -e POSTGRES_PASSWORD=password postgres`

```shell
$ docker exec -it pg /bin/bash                               
root@187961c81d2e:/# psql -U postgres
psql (14.2 (Debian 14.2-1.pgdg110+1))
Type "help" for help.
```