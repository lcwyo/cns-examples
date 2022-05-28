---
layout: default
title: Step Six
parent: Guestbook GO
nav_order: 2
---


### Step Six: Create the guestbook service <a id="step-six"></a>

Just like the others, we create a service to group the guestbook pods but this time, to make the guestbook front end externally visible, we specify `"type": "LoadBalancer"`.

1. Use the [guestbook-service.yaml](guestbook-service.yaml) file to create the guestbook service by running the `kubectl create -f` *`filename`* command:

    ```console
    $ kubectl create -f examples/guestbook-go/guestbook-service.yaml
    ```


2. To verify that the guestbook service is up, list the services you created in the cluster with the `kubectl get services` command:

    ```console
    $ kubectl get services
    NAME              CLUSTER_IP       EXTERNAL_IP       PORT(S)       SELECTOR               AGE
    guestbook         10.0.217.218     146.148.81.8      3000/TCP      app=guestbook          1h
    redis-master      10.0.136.3       <none>            6379/TCP      app=redis,role=master  1h
    redis-replica     10.0.21.92       <none>            6379/TCP      app-redis,role=replica 1h
    ...
    ```

    Result: The service is created with label `app=guestbook`.
