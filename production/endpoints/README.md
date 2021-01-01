# ESP gRPC endpoints

## Deploy new version of endpoints service for an ESP

```sh
export M_ROOT="$HOME/go/microclimate-api"

$M_ROOT/client-generators/go/generate_client.sh

gcloud endpoints services deploy $M_ROOT/protobuf/api_descriptor.pb  $M_ROOT/production/endpoints/api_config_auth.yaml

gcloud services enable servicemanagement.googleapis.com
gcloud services enable servicecontrol.googleapis.com
gcloud services enable endpoints.googleapis.com

gcloud services enable <service.name>

gcloud endpoints services list
```
