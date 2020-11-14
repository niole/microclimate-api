# general install

- install go1.15.3
- Clone git project to $GOPATH directory.
- `./install.sh`
- `./generate_client.sh`

# development

## run

- Test stuff out by running the server.go and then
hitting it with the code in the main.go:
```sh
go run src/server/impl/server.go
go run src/main.go
```

# production

## run services

- `docker-compose up`

## generate java client

- `./generate_java_client.sh`
- `cd ./java_api/api && ls`
