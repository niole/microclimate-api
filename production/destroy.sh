#!/bin/bash

echo -e "\nThis script destroys everything k8s except the ManagedCertificate resource\n"

echo -e "\nDestroying neg ingress"
kubectl delete -f ./production/k8s/neg-ingress/ingress.yaml

echo -e "\nDestroying all services"
kubectl delete -Rf ./production/k8s/services/

echo -e "\nDestroying database"
kubectl delete -Rf ./production/k8s/database/

echo -e "\nDestroying storage"
kubectl delete -Rf ./production/k8s/storage/

echo -e "\nDestroying secrets"
kubectl delete -f ./production/k8s/secrets/

echo -e "\nDestroying NEGs"
echo -e "\nDestroying deployment-service NEG"
gcloud compute network-endpoint-groups list | \
grep deployment-service | \
awk '{print $1}' | \
while IFS= read -r line; do; gcloud compute network-endpoint-groups delete $line ; done

echo "Destroying event-service NEG"
gcloud compute network-endpoint-groups list | \
grep event-service | \
awk '{print $1}' | \
while IFS= read -r line; do; gcloud compute network-endpoint-groups delete $line ; done

echo "Destroying peripheral-service NEG"
gcloud compute network-endpoint-groups list | \
grep peripheral-service | \
awk '{print $1}' | \
while IFS= read -r line; do; gcloud compute network-endpoint-groups delete $line ; done

echo "Destroying user-service NEG"
gcloud compute network-endpoint-groups list | \
grep user-service | \
awk '{print $1}' | \
while IFS= read -r line; do; gcloud compute network-endpoint-groups delete $line ; done
