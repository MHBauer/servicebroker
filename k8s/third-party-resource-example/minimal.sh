#!/usr/bin/env bash

set -o xtrace
set -o errexit

kubectl create -f tpr-definition.yaml
#sleep 15 # the create needs some time to be propagated
kubectl get thirdpartyresources abc.tuv.xyz --output=yaml
kubectl get Abc --output=yaml
kubectl delete -f tpr-definition.yaml
