# Service Broker Specification

**Table of Contents**

- [Service Broker Specification](#service-broker-specification)
    - [Overview](#overview)
    - [Notation](#notation)
    - [Definitions](#definitions)
    - [Scope](#scope)
        - [in scope](#in-scope)
        - [out of scope](#out-of-scope)
    - [Use Case](#use-case)
- [API](#api)
    - [description of stuff generalized over the whole api](#description-of-stuff-generalized-over-the-whole-api)
        - [errors](#errors)
            - [body](#body)
    - [Resource Definitions](#resource-definitions)
        - [catalog](#catalog)
            - [/v2/catalog](#v2catalog)
        - [service instances](#service-instances)
        - [service bindings](#service-bindings)

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

## Scope 

client <X> platform <-> broker <X> backend

### in scope

communication to a broker from a platform

### out of scope

communication to whatever platform talks to the broker

communication from broker to backend




## Use Case
The only requirement for a service to be available to end users is
that a service implements the broker API.

Thus a service can be entirely provided and managed by a third party,
or contained as an internal deployment and managed wholly within a
platform.

(?maybe take this out, too specific?)  Some examples from Cloud
Foundry:
* Entire service packaged and deployed by BOSH alongside Cloud Foundry
* Broker packaged and deployed by BOSH alongside Cloud Foundry, rest
  of the service deployed and maintained by other means
* Broker (and optionally service) pushed as an application to Cloud
  Foundry user space
* Entire service, including broker, deployed and maintained outside of
  Cloud Foundry by other means


# API

    * this section describes the api and how to access it
    
notation in this section consisting of a common format with all
sections being optional

```
request
route
paramters
curl example
response
body
```

## description of stuff generalized over the whole api

  * version header
  * authentication
  * format
  * asynchronous operations and polling the last operation
  * blocking operations
  * errors
  * orphans
  * etc

### format

All messages with an http message body shall have the body formatted as json (rfc7159).
<!-- I'm not aware of any messages we're sending that are not formatted as json. -->

### errors 

##### response

Broker failures beyond the scope of the well-defined HTTP response
codes listed in the individual resource definitions should return an
appropriate HTTP response code (chosen to accurately reflect the
nature of the failure) and a body containing a valid JSON Object (not
an array).

##### body 

All response bodies must be valid JSON. This is for future
compatibility; it will be easier to add fields in the future if JSON
is expected rather than to support the cases when a JSON body may or
may not be returned.

For error responses, the following fields are valid. Others will be
ignored. If an empty JSON object is returned in the body `{}`, a
generic message containing the HTTP response code returned by the
broker will be displayed to the requestor.

|  response |type   | description  |
|---|---|---|
|  description | string  | An error message explaining why the request failed. This message will be displayed to the user who initiated the request.  |


```json
{
  "description": "Something went wrong. Please contact support at http://support.example.com."
}
```




## Resource Definitions

A conforming implementation shall expose the following rest resources.
  * /v2/catalog
  * /v2/service_instances

### catalog

#### /v2/catalog

### service instances

### service bindings

