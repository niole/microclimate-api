# dev

```sh
# install stuff

./install.sh

# generate clients

./client-generators/go/generate_client.sh

# run stuff

docker-compose run --service-ports db
go run ./event/main/server.go -serverPort 8080
go run ./main.go

# or

docker-compose up
```

# prod

Production code depends on GKE.

```
./production/install.sh
```
