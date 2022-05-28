---
layout: default
title: Step One
parent: Guestbook GO
nav_order: 2
---



### Step One: Create the Redis master pod<a id="step-one"></a>

Use the `examples/guestbook-go/redis-master-deployment.yaml` file to create a [deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) and Redis master [pod](https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/). The pod runs a Redis key-value server in a container.

1. Use the [redis-master-deployment.yaml](redis-master-deployment.yaml) file to create the Redis master replication controller in your Kubernetes cluster by running the `kubectl create -f` *`filename`* command:

    ```console
    $ kubectl create -f examples/guestbook-go/redis-master-deployment.yaml
   
    ```

2. To verify that the redis-master controller is up, list the replication controllers you created in the cluster with the `kubectl get deployments` command(if you don't specify a `--namespace`, the `default` namespace will be used. The same below):

    ```console
    $ kubectl get deployments -o wide
    NAME            READY   UP-TO-DATE   AVAILABLE   AGE     CONTAINERS   IMAGES                             SELECTOR
    redis-master    1/1     1            1           10m     master       docker.io/bitnami/redis:latest     app=redis,role=master,tier=backend
    ...
    ```

    Result: The the deployment creates the single Redis master pod.

3. To verify that the redis-master pod is running, list the pods you created in cluster with the `kubectl get pods` command:

    ```console
    $ kubectl get pods
    NAME                        READY     STATUS    RESTARTS   AGE
    redis-master-xx4uv          1/1       Running   0          1m
    ...
    ```

    Result: You'll see a single Redis master pod and the machine where the pod is running after the pod gets placed (may take up to thirty seconds).
