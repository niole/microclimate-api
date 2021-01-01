#!/bin/bash

echo -e "\nThis script destroys everything k8s except the ManagedCertificate resource\n"

echo -e "\nDestroying neg ingress"
kubectl delete -f ./production/k8s/neg-ingress/

echo -e "\nDestroying deployment"
kubectl delete -f ./production/k8s/deployment/

echo -e "\nDestroying user service"
kubectl delete -f ./production/k8s/users/

echo -e "\nDestroying peripheral service"
kubectl delete -f ./production/k8s/peripherals/

echo -e "\nDestroying event service"
kubectl delete -f ./production/k8s/events/

echo -e "\nDestroying database"
kubectl delete -f ./production/k8s/database/

echo -e "\nDestroying secrets"
kubectl delete -f ./production/k8s/secrets/
