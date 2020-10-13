#!/bin/bash
kubectl create namespace wordpresses
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install -n wordpresses big-mariadb --values values-mariadb.yaml bitnami/mariadb