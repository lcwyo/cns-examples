---
layout: default
title: Create the Redis replica deployment
parent: Guestbook GO
nav_order: 2
---


### Step Three: Create the Redis replica deployment

The Redis master we created earlier is a single pod (REPLICAS = 1), while the Redis read replicas we are creating here are 'replicated' pods. In Kubernetes, a replication controller is responsible for managing the multiple instances of a replicated pod.

1. Use the file [redis-replica-deployment.yaml](redis-replica-deployment.yaml) to create the replication controller by running the `kubectl create -f` *`filename`* command:

    ```console
    $ kubectl create -f examples/guestbook-go/redis-replica-deployment.yaml
    
    ```

2. To verify that the redis-replica controller is running, run the `kubectl get deployments` command:

    ```console
    $ kubectl get deployments -o wide
    NAME            READY   UP-TO-DATE   AVAILABLE   AGE     CONTAINERS   IMAGES                             SELECTOR    
    redis-master    1/1     1            1           10m     master       docker.io/bitnami/redis:latest     app=redis,role=master,tier=backend
    redis-replica   2/2     2            2           8m29s   replica      docker.io/lcwyo/redis-replica:v3   app=redis,role=replica,tier=backend
    ...
    ```

    Result: The deployment controller creates and configures the Redis replica pods through the redis-master service (name:port pair, in our example that's `redis-master:6379`).

    Example:
    The Redis replicas get started by the deployment controller with the following command:

    ```console
    redis-server --replicaof redis-master 6379
    ```

3. To verify that the Redis master and replicas pods are running, run the `kubectl get pods` command:

    ```console
    $ kubectl get pods
    NAME                          READY     STATUS    RESTARTS   AGE
    redis-master-xx4uv            1/1       Running   0          18m
    redis-replica-b6wj4           1/1       Running   0          1m
    redis-replica-iai40           1/1       Running   0          1m
    ...
    ```

    Result: You see the single Redis master and two Redis replica pods.
