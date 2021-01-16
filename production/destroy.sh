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

# TODO destroy NEGs
