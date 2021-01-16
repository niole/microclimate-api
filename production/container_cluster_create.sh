#!/bin/bash

zone=${1:-us-central1-c}

gcloud container clusters create mic-cluster \
    --preemptible \
    --enable-ip-alias \
    --create-subnetwork="" \
    --network=default \
    --zone=$zone
