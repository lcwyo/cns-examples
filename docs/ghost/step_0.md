---
layout: default
title: Prerequisites
parent: Ghost
nav_order: 2
---


### Step Zero: Prerequisites

This example assumes that you have a namespace with cNS. See the [Getting Started Guides]() for details about ordering a namespace.

Documentation missing 
{: .label .label-yellow }
_Download kubectl from UI, use kubectl_

#### Install Helm

Helm is a tool for managing Kubernetes charts. Charts are packages of pre-configured Kubernetes resources.

To install Helm, refer to the Helm install guide and ensure that the helm binary is in the PATH of your shell.

#### Add Repo

The following command allows you to download and install all the charts from this repository:


using helm, add the bitnami repo

```console
$ helm repo add bitnami https://charts.bitnami.com/bitnami
```


**Tip:** View all the `kubectl` commands, including their options and descriptions in the [kubectl CLI reference](https://kubernetes.io/docs/user-guide/kubectl-overview/).
