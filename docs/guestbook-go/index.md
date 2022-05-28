---
layout: default
title: Guestbook GO
nav_order: 2
has_children: true
permalink: /docs/guestbook-go
---


## Guestbook GO Example

This example shows how to build a simple multi-tier web application using Kubernetes and Docker. The application consists of a web front end, Redis master for storage, and replicated set of Redis replicas, all for which we will create Kubernetes deployments, pods, services, and ingress.


{: .no_toc }

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






