#!/usr/bin/env bash

set -o xtrace

kubectl delete -f tpr-definition.yaml
kubectl delete -f tpr2-definition.yaml
kubectl delete -f tpr3-definition.yaml
