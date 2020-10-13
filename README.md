# Setup cluster
J'imagine plusieurs *Deployments* Wordpress, un par site, qui sont à 0 *Replicas* sauf quand
un administrateur veut modifier son site.

Pour réduire les coûts, on peut faire en sorte de mettre tout le stockage sur un seul PersitentVolume, avec 
des *subPaths*.

Pour la base de données, on peut faire tourner un seul StatefulSet MariaDB 
qui contient plusieurs bases de données (ou une seule et on utilise des préfixe de tables différents)

## Déploiement de l'appli
Il lui faut un ServiceAccount monté dans ``/var/run/secrets/kubernetes.io/serviceaccount`` (voir [la doc](https://github.com/kubernetes/client-go/tree/master/examples/in-cluster-client-configuration))

