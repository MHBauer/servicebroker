# Service Broker Specification

## Overview

This specification defines an API that provides access to services in
a generic way that is not tied to a specific platform or cloud
provider. It is based on the Cloud Foundry API.


## Notation

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT",
      "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in
      this document are to be interpreted as described in RFC 2119.

## Definitions

 * Application - 
 * Service - a product or API that is provided by a third party.
 * Broker is the component which implements the service broker API.
   - advertises a catalog of service offerings and service
     plans
     - interprets calls for provision (create), bind, unbind, and
deprovision (delete).
        - What a broker does with each call can vary between services; in
     general,
     - 'provision' reserves resources on a service
     - 'bind' delivers information to an application necessary for
       accessing the resource.
   - reserved resources are a Service Instance.
     - What a service instance represents can vary by service; it
       could be a single database on a multi-tenant server, a
       dedicated cluster, or even just an account on a web
       application.
 * Backend - opaque and platform/provider specific implementation that
   can provision services.
 * Catalog - a list of service definitions.
 * Plan - a possible set of qualities that a service may have. (usually
        this is in the form of a cost to use the service. Different levels
        of service quality or other restrictions.)
 * Instance - a deployment of a service by a backend. Can be bound to
        with a bind call.
    - a provisionment [find a better word] of the service. 
    - resources are set aside for an instance.
    - an instance existing implies that something was done to make a
      service ready for access.
 * Binding - a dependency of an application on a service that ties the
   application to the service. A binding contains the relevant
   connection information for an application to connect. This is
   information such as the location of the resource, as well as
   authorization information like username and password. 
    - A registration of use of a service by an external (outside of
      the service backend) application.
    - 'accessed with' relationship. ServiceInstance X is accessed with
      Binding Y.
    - 


## Use Case
The only requirement for a service to be available to end users is that a service implements the broker API.

Thus a service can be entirely provided and managed by a third party, or contained as an internal deployment and managed wholly within a platform.

(?maybe take this out, too specific?)
Some examples from Cloud Foundry:
* Entire service packaged and deployed by BOSH alongside Cloud Foundry
* Broker packaged and deployed by BOSH alongside Cloud Foundry, rest of the service deployed and maintained by other means
* Broker (and optionally service) pushed as an application to Cloud Foundry user space
* Entire service, including broker, deployed and maintained outside of Cloud Foundry by other means


## 
