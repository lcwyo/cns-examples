---
layout: default
title: Create the guestbook deployment
parent: Guestbook PHP
nav_order: 2
---


### Step Five: Create the guestbook deployment

This is a simple php-fpm / nginx server that is configured to talk to either the replica or master services depending on whether the request is a read or a write. Like the Redis replica pods, these pods are also managed by a deployment.

#### Deploy the php-fpm

1. Use the [guestbook-php-deployment.yaml](guestbook-php-deployment.yaml) file to create the guestbook replication controller by running the `kubectl create -f` *`filename`* command:

    ```console
    $ kubectl create -f examples/guestbook-php/guestbook-php-deployment.yaml
    
    ```

 Tip: If you want to modify the guestbook code open the `app` of this example and read the README.md and the Makefile. If you have pushed your custom image be sure to update the `image` accordingly in the guestbook-deployment.yaml.

2. To verify that the guestbook replication controller is running, run the `kubectl get deployments` command:

    ```console
    $ kubectl get deployments
    NAME            READY   UP-TO-DATE   AVAILABLE   AGE
    guestbook-php       3/3     3            3           49m
    redis-master    1/1     1            1           18m
    redis-replica   2/2     2            2           16m
    ...
    ```

3. To verify that the guestbook pods are running (it might take up to thirty seconds to create the pods), list the pods you created in cluster with the `kubectl get pods` command:

    ```console
    $ kubectl get pods
    NAME                           READY     STATUS    RESTARTS   AGE
    guestbook-php-3crgn                1/1       Running   0          2m
    guestbook-php-gv7i6                1/1       Running   0          2m
    guestbook-php-x405a                1/1       Running   0          2m
    redis-master-xx4uv             1/1       Running   0          23m
    redis-replica-b6wj4              1/1       Running   0          6m
    redis-replica-iai40              1/1       Running   0          6m
    ...
    ```

    Result: You see a single Redis master, two Redis replicas, and three guestbook pods.

#### Deploy the nginx