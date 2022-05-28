---
layout: default
title: Helm upgrade
parent: Ghost
nav_order: 2
---

### Step Two: Helm upgrade


now we can use helm upgrade to complete the installation.


```console
$ helm upgrade --namespace lances-test-24a46 cns-ghost bitnami/ghost \
    --set service.type=LoadBalancer,ghostHost=$APP_HOST,ghostPassword=$GHOST_PASSWORD,mariadb.auth.rootPassword=$MARIADB_ROOT_PASSWORD,mariadb.auth.password=$MARIADB_PASSWORD \
    --set resources.limits.memory=2Gi \
    --set resources.limits.cpu=1500m \
    --set resources.requests.cpu=1000m \
    --set resources.requests.memory=512Mi
```

** Please be patient while the chart is being deployed **


1. Get your Ghost login credentials by running:

```console
  echo Email:    user@example.com
  echo Password: $(kubectl get secret cns-ghost -o jsonpath="{.data.ghost-password}" | base64 --decode)
  ```

2. use the URL and username and password to login

```console
  echo Blog URL  : http://$APP_HOST/
  echo Admin URL : http://$APP_HOST/ghost
```