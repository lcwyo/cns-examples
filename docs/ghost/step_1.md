---
layout: default
title: Use Helm to Install Ghost
parent: Ghost
nav_order: 2
---



### Step One: Use helm to install ghost

```console
$ helm install cns-ghost bitnami/ghost --set ghostUsername=cns-user --set resources.limits.cpu=1500m --set resources.limits.memory=2Gi

```

get the credintials and set them as variables

```console
$ export APP_HOST=$(kubectl get svc  cns-ghost --template "{{ range (index .status.loadBalancer.ingress 0) }}{{ . }}{{ end }}")
$ export GHOST_PASSWORD=$(kubectl get secret cns-ghost -o jsonpath="{.data.ghost-password}" | base64 --decode)
$ export MARIADB_ROOT_PASSWORD=$(kubectl get secret cns-ghost-mariadb -o jsonpath="{.data.mariadb-root-password}" | base64 --decode)
$ export MARIADB_PASSWORD=$(kubectl get secret cns-ghost-mariadb -o jsonpath="{.data.mariadb-password}" | base64 --decode)
```

check if mariadb is running 

```console
$ kubectl get pods
NAME                  READY   STATUS    RESTARTS   AGE
cns-ghost-mariadb-0   1/1     Running   0          2m34s
```

```console
$ kubectl statefulset -o wide 
NAME                READY   AGE   CONTAINERS   IMAGES
cns-ghost-mariadb   1/1     67s   mariadb      docker.io/bitnami/mariadb:10.5.15-debian-10-r11
```

check if persistent volume claims where created

```console
kubectl get pvc
NAME                       STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
data-cns-ghost-mariadb-0   Bound    pvc-7b0df4d5-d585-4b53-afd4-7ac00ce24977   8Gi        RWO            standard       4m5s

```

