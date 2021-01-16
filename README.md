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
./production/container_cluster_create.sh us-central1-c
kubectl apply -f production/k8s/certificate/certificate.yaml
./production/install.sh
```
