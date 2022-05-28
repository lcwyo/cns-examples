---
layout: default
title: Step Seven
parent: Guestbook GO
nav_order: 2
---


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