#!/usr/bin/env bash

set -o xtrace
set -o errexit

kubectl create -f tpr-definition.yaml
sleep 15 # the create needs some time to be propagated
kubectl get thirdpartyresources
kubectl get thirdpartyresources --output=yaml
kubectl get thirdpartyresources abc.tuv.xyz --output=yaml
# kubectl get thirdpartyresources Abc --output=yaml
kubectl get Abc --output=yaml
sleep 1
kubectl create -f tpr-instance.yaml
sleep 4
kubectl get Abc my-new-abc-object --output=yaml
sleep 4
kubectl delete Abc my-new-abc-object
sleep 1 
kubectl delete -f tpr-definition.yaml
