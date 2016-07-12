#!/usr/bin/env bash

function teardown() {
    ./clean.sh
}

@test "define two third party resources in the same domain and make sure they both exist" {
    kubectl create -f tpr-definition.yaml
    kubectl create -f tpr2-definition.yaml
    sleep 15 # the create needs some time to be propagated
    kubectl get Abc --output=yaml
    kubectl get Qrs --output=yaml
}

@test "define two third party resources in different domains and make sure they both exist" {
    kubectl create -f tpr-definition.yaml
    kubectl create -f tpr3-definition.yaml
    sleep 15 # the create needs some time to be propagated
    kubectl get Abc --output=yaml
    kubectl get Qrs --output=yaml
}


@test "validate timing" {
    kubectl create -f tpr-definition.yaml
    kubectl create -f tpr2-definition.yaml
    # sleep 15 # I should not need this sleep here
    kubectl get Abc --output=yaml
    kubectl get Qrs --output=yaml
}
