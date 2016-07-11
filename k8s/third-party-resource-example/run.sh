#!/usr/bin/env bash

set -o xtrace
set -o errexit

kubectl create -f tpr-definition.yaml
kubectl create -f tpr2-definition.yaml
sleep 15 # the create needs some time to be propagated
kubectl get thirdpartyresources
kubectl get thirdpartyresources --output=yaml
kubectl get thirdpartyresources abc.tuv.xyz --output=yaml
# kubectl get thirdpartyresources Abc --output=yaml
sleep 5
kubectl get Abc --output=yaml
kubectl get Qrs --output=yaml
sleep 1
kubectl create -f tpr-instance.yaml
kubectl create -f tpr-instance2.yaml
kubectl create -f tpr2-instance.yaml
kubectl create -f tpr2-instance2.yaml
sleep 4
kubectl get Abc my-new-abc-object --output=yaml
kubectl get Abc my-2-abc-object --output=yaml
kubectl get Qrs my-new-abc-object --output=yaml
kubectl get Qrs my-2-abc-object --output=yaml


sleep 4
kubectl delete Abc my-new-abc-object
kubectl delete Abc my-2-abc-object
sleep 1 
kubectl delete -f tpr-definition.yaml
kubectl delete -f tpr2-definition.yaml
