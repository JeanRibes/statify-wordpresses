helm repo add bitnami https://charts.bitnami.com/bitnami
helm install -n wordpresses wp1 --values values-wordpress.yaml bitnami/wordpress