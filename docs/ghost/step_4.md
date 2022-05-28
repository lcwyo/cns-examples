---
layout: default
title: Cleanup
parent: Ghost
nav_order: 2
---

### Step four: clean up


now we can use helm upgrade to complete the installation.

1. remove the helm installation

```console
$ helm delete cns-ghost
```

2. remove the maria-db PVC

```console
$ kubectl delete pvc data-cns-ghost-mariadb-0
```