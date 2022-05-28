---
layout: default
title: Create the Redis replica service
parent: Guestbook GO
nav_order: 2
---


### Step Four: Create the Redis replica service

Just like the master, we want to have a service to proxy connections to the read replicas. In this case, in addition to discovery, the Redis replica service provides transparent load balancing to clients.

1. Use the [redis-replica-service.yaml](redis-replica-service.yaml) file to create the Redis replica service by running the `kubectl create -f` *`filename`* command:

    ```console
    $ kubectl create -f examples/guestbook-go/redis-replica-service.yaml
   
    ```

2. To verify that the redis-replica service is up, list the services you created in the cluster with the `kubectl get services` command:

    ```console
    $ kubectl get services
    NAME              CLUSTER_IP       EXTERNAL_IP       PORT(S)       SELECTOR               AGE
    redis-master      10.0.136.3       <none>            6379/TCP      app=redis,role=master  1h
    redis-replica     10.0.21.92       <none>            6379/TCP      app-redis,role=replica 1h
    ...
    ```

    Result: The service is created with labels `app=redis` and `role=replica` to identify that the pods are running the Redis replicas.

Tip: It is helpful to set labels on your services themselves--as we've done here--to make it easy to locate them later.
