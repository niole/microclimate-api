# general install

- install go1.15.3
- Clone git project to $GOPATH directory.
- `./install.sh`
- `./generate_client.sh`

# development

## run

- start db
```sh
docker run \
--name dev_db \
-e MYSQL_ROOT_PASSWORD=pw \
-e MYSQL_USER=niole \
-e MYSQL_PASSWORD=pw \
-e MYSQL_DATABASE=default \
-p 3306:3306 \
-it mysql:8.0.22
```
- Test stuff out by running the server.go and then
hitting it with the code in the main.go

run the server
```sh
export DB_HOST="0.0.0.0"
go run src/server/impl/server.go
```

run the client
```sh
go run src/main.go
```

stop the db
```sh
docker stop dev_db
```

# production

## dependencies

- docker 19
- docker-compose 1.27.4

## run services

- `docker-compose up`

## generate java client

- `./generate_java_client.sh`
- `cd ./java_api/api && ls`
