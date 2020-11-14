# development

## install and run

- install go
- Clone git project to $GOPATH directory.
- `./install.sh`
- `./generate_client.sh`
- Test stuff out by running the server.go and then
hitting it with the code in the main.go:
```sh
go run src/server/impl/server.go
go run src/main.go
```

# production

## install

- install go
- Clone git project to $GOPATH directory.
- `./install.sh`

## run services

- `docker-compose up`

## generate java client

- `./generate_java_client.sh`
- `cd ./java_api/api && ls`
