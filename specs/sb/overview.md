---
title: Overview
owner: Core Services
---

## <a id='architecture-terminology'></a>Architecture & Terminology##


<image src="images/managed-services.png">

## <a id='implementation-deployment'></a>Implementation & Deployment ##

How a service is implemented is up to the service provider/developer. Cloud Foundry only requires that the service provider implement the service broker API. A broker can be implemented as a separate application, or by adding the required http endpoints to an existing service.

