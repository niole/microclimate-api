#!/bin/bash

echo -e "\nThis script assumes that you have already applied the ManagedCertificate resource, the gcp endpoints already exist, and all referenced domains already exist\n"

if [ "$1" = "" ]
then
    echo "Checking for gcloud"
    if [ "$(which gcloud)" = "" ]
    then
        echo "Please install gcloud"
        exit 1
    fi
    echo "gcloud exists"

    echo -e "\nChecking for kubectl"
    if [ "$(which kubectl)" = "" ]
    then
        echo "Please install kubectl"
        exit 1
    fi
    echo "kubectl exists"

    echo -e "\nCurrent kubectl context"
    kubectl config current-context

    echo -e "\nCurrent gloud configuration"
    gcloud config list

    echo -e "\nListing available endpoint services"
    gcloud endpoints services list

    echo -e "\nListing active container clusters"
    gcloud container clusters list

    while true; do
        read -p "Verify that context is correct. Do you wish to continue?" yn
        case $yn in
            [Yy]* ) break;;
            [Nn]* ) exit;;
        esac
    done

    echo -e "\nContinuing...\n"
fi

echo -e "\nApplying secrets"
kubectl apply -f ./production/k8s/secrets/

echo -e "\nApplying database"
kubectl apply -f ./production/k8s/database/

echo -e "\nApplying deployment"
kubectl apply -f ./production/k8s/deployment/

echo -e "\nApplying neg ingress"
kubectl apply -f ./production/k8s/neg-ingress/
