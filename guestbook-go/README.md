## Guestbook Example

This example shows how to build a simple multi-tier web application using Kubernetes and Docker. The application consists of a web front end, Redis master for storage, and replicated set of Redis replicas, all for which we will create Kubernetes deployments, pods, services, and ingress.

##### Table of Contents

 * [Step Zero: Prerequisites](#step-zero)
 * [Step One: Create the Redis master pod](#step-one)
 * [Step Two: Create the Redis master service](#step-two)
 * [Step Three: Create the Redis replica pods](#step-three)
 * [Step Four: Create the Redis replica service](#step-four)
 * [Step Five: Create the guestbook pods](#step-five)
 * [Step Six: Create the guestbook service](#step-six)
 * [Step Seven: Create the guestbook ingress](#step-seven)
 * [Step Eight: View the guestbook](#step-eight)
 * [Step Nine: Cleanup](#step-nine)

### Step Zero: Prerequisites <a id="step-zero"></a>

This example assumes that you have a working cluster. See the [Getting Started Guides]() for details about creating a cluster.

**Tip:** View all the `kubectl` commands, including their options and descriptions in the [kubectl CLI reference](https://kubernetes.io/docs/user-guide/kubectl-overview/).

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

### Step Two: Create the Redis master service <a id="step-two"></a>

A Kubernetes [service](https://kubernetes.io/docs/concepts/services-networking/service/) is a named load balancer that proxies traffic to one or more pods. The services in a Kubernetes cluster are discoverable inside other pods via environment variables or DNS.

Services find the pods to load balance based on pod labels. The pod that you created in Step One has the label `app=redis` and `role=master`. The selector field of the service determines which pods will receive the traffic sent to the service.

1. Use the [redis-master-service.yaml](redis-master-service.yaml) file to create the service in your Kubernetes cluster by running the `kubectl create -f` *`filename`* command:

    ```console
    $ kubectl create -f examples/guestbook-go/redis-master-service.yaml
   
    ```

2. To verify that the redis-master service is up, list the services you created in the cluster with the `kubectl get services` command:

    ```console
    $ kubectl get services
    NAME              CLUSTER_IP       EXTERNAL_IP       PORT(S)       SELECTOR               AGE
    redis-master      10.0.136.3       <none>            6379/TCP      app=redis,role=master  1h
    ...
    ```

    Result: All new pods will see the `redis-master` service running on the host (`$REDIS_MASTER_SERVICE_HOST` environment variable) at port 6379, or running on `redis-master:6379`. After the service is created, the service proxy on each node is configured to set up a proxy on the specified port (in our example, that's port 6379).


### Step Three: Create the Redis replica pods <a id="step-three"></a>

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

### Step Four: Create the Redis replica service <a id="step-four"></a>

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

### Step Five: Create the guestbook pods <a id="step-five"></a>

This is a simple Go `net/http` ([negroni](https://github.com/codegangsta/negroni) based) server that is configured to talk to either the replica or master services depending on whether the request is a read or a write. The pods we are creating expose a simple JSON interface and serves a jQuery-Ajax based UI. Like the Redis replica pods, these pods are also managed by a deployment controller.

1. Use the [guestbook-deployment.yaml](guestbook-deployment.yaml) file to create the guestbook replication controller by running the `kubectl create -f` *`filename`* command:

    ```console
    $ kubectl create -f examples/guestbook-go/guestbook-deployment.yaml
    
    ```

 Tip: If you want to modify the guestbook code open the `app` of this example and read the README.md and the Makefile. If you have pushed your custom image be sure to update the `image` accordingly in the guestbook-deployment.yaml.

2. To verify that the guestbook replication controller is running, run the `kubectl get deployments` command:

    ```console
    $ kubectl get deployments
    NAME            READY   UP-TO-DATE   AVAILABLE   AGE
    guestbook       3/3     3            3           49m
    redis-master    1/1     1            1           18m
    redis-replica   2/2     2            2           16m
    ...
    ```

3. To verify that the guestbook pods are running (it might take up to thirty seconds to create the pods), list the pods you created in cluster with the `kubectl get pods` command:

    ```console
    $ kubectl get pods
    NAME                           READY     STATUS    RESTARTS   AGE
    guestbook-3crgn                1/1       Running   0          2m
    guestbook-gv7i6                1/1       Running   0          2m
    guestbook-x405a                1/1       Running   0          2m
    redis-master-xx4uv             1/1       Running   0          23m
    redis-replica-b6wj4              1/1       Running   0          6m
    redis-replica-iai40              1/1       Running   0          6m
    ...
    ```

    Result: You see a single Redis master, two Redis replicas, and three guestbook pods.

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

### Step seven: create the guestbook ingress <a id="step-seven"></a>

In order to make the guesstbook front end externally visiable, we create an ingress using the shared ingress controller.


1. Use the [guestbook-ingress.yaml](guestbook-ingress.yaml) file to create the guestbook ingress by running the `kubectl create -f` *`filename`* command:

    ```console
    $ kubectl create -f examples/guestbook-go/guestbook-ingress.yaml
    ```

2. To verify that the guestbook ingress is up, list the ingress you created in the cluster with the `kubectl get ingresses.v1.networking.k8s.io` command:

    ```console
    $ kubectl get ingresses.v1.networking.k8s.io
    NAME                CLASS        HOSTS                                  ADDRESS          PORTS     AGE
    guestbook-ingress   cmy-shared   guestbook.cloud.eu1.cloudmobility.io   62.153.212.225   80, 443   37m
    ...
    ```

### Step eight: View the guestbook <a id="step-eight"></a>

You can now play with the guestbook that you just created by opening it in a browser (it might take a few moments for the guestbook to come up).

 * **Local Host:**
    If you are running Kubernetes locally, to view the guestbook, navigate to `http://localhost:3000` in your browser.

 * **Remote Host:**
    1. To view the guestbook on a remote host, locate the external IP of the load balancer in the **IP** column of the `kubectl get services` output. In our example, the internal IP address is `10.0.217.218` and the external IP address is `146.148.81.8` (*Note: you might need to scroll to see the IP column*).

    2. Append port `3000` to the IP address (for example `http://146.148.81.8:3000`), and then navigate to that address in your browser.

    Result: The guestbook displays in your browser:

    ![Guestbook](guestbook-page.png)

    **Further Reading:**
    If you're using Google Compute Engine, see the details about limiting traffic to specific sources at [Google Compute Engine firewall documentation][gce-firewall-docs].

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

