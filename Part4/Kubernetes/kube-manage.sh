#!/bin/bash
ACTION=$1

case $ACTION in
  "create")
    kubectl create -f adminer-deployment.yaml,adminer-service.yaml,mysql-deployment.yaml,mysql-service.yaml,rabbitmq-deployment.yaml,rabbitmq-service.yaml,webservice-deployment.yaml,webservice-service.yaml
    ;;

  "apply")
    kubectl apply -f adminer-deployment.yaml,adminer-service.yaml,mysql-deployment.yaml,mysql-service.yaml,rabbitmq-deployment.yaml,rabbitmq-service.yaml,webservice-deployment.yaml,webservice-service.yaml
    ;;

  "delete")
    kubectl delete -f adminer-deployment.yaml,adminer-service.yaml,mysql-deployment.yaml,mysql-service.yaml,rabbitmq-deployment.yaml,rabbitmq-service.yaml,webservice-deployment.yaml,webservice-service.yaml
    ;;

esac