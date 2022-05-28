---
layout: default
title: Step Nine
parent: Guestbook GO
nav_order: 2
---



### Step nine: Cleanup <a id="step-nine"></a>

After you're done playing with the guestbook, you can cleanup by deleting the guestbook service and removing the associated resources that were created, including load balancers, forwarding rules, target pools, and Kubernetes replication controllers and services.

Delete all the resources by running the following `kubectl delete -f` *`filename`* command:

```console
$ kubectl delete -f examples/guestbook-go
guestbook-controller
guestbook
redid-master-controller
redis-master
redis-replica-controller
redis-replica
```

Tip: To turn down your Kubernetes cluster, follow the corresponding instructions in the version of the
[Getting Started Guides](https://kubernetes.io/docs/getting-started-guides/) that you previously used to create your cluster.

