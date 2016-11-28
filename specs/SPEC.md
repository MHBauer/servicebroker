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

 * Application - program desiring to use a service. Is in need of
   access to a service.
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

## managing brokers

### lifecycle

  1. register broker
  2. 
  
### multiple brokers

# API

    * this section describes the api and how to access it
    
    
???
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

### version header
Requests from the Cloud Controller to the broker contain a header that defines
the version number of the Broker API that Cloud Controller will use.
This header will be useful in future minor revisions of the API to allow
brokers to reject requests from Cloud Controllers that they do not understand.
While minor API revisions will always be additive, it is possible that brokers
will come to depend on a feature that was added after 2.0, so they may use this
header to reject the request.
Error messages from the broker in this situation should inform the operator of
what the required and actual version numbers are so that an operator can go
upgrade Cloud Controller and resolve the issue.
A broker should respond with a `412 Precondition Failed` message when rejecting
a request.

The version numbers are in the format `MAJOR.MINOR`, using semantic versioning
such that 2.9 comes before 2.10.
An example of this header as of publication time is:

`X-Broker-Api-Version: 2.9`

### authentication

A client authenticates with a broker using HTTP basic authentication
(the `Authorization:` header) on every request and will reject any
broker registrations that do not contain a username and password.

The broker is responsible for checking the username and password and
returning a `401 Unauthorized` message if credentials are invalid.

Connecting to a broker using SSL is supported. OPTIONAL / RECOMMENDED. 
<!-- This should probably be REQUIRED -->

Your service broker should validate the username and password sent in every
request; otherwise, anyone could curl your broker to delete service instances.


### format

All messages with an http message body shall have the body formatted
as json (rfc7159).
<!-- I'm not aware of any messages we're sending that are not formatted as json. -->

For example, responses should return an empty object '{}' even when
there is no content to return. This is not to be interpreted as being
in contravention to the HTTP spec for responses that MUST NOT have a
message body.

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

## user facing metadata
<!-- I am not sure this needs it's own section -->

Services may come from one or many service brokers. Clients may need
to distinguish among them.

All clients are not expected to have the same requirements for
information to expose about services and plans. The broker API enables
broker authors to provide metadata required by different clients.

Services and plans support `metadata` field. The `metadata` field is
optional. The contents of the `metadata` field may be validated by
clients. 

### suggested fields 

### example 

The example below contains a catalog of one service, having one
service plan. Of course, a broker can offering a catalog of many
services, each having many plans.

```
{
   "services":[
      {
      "id":"766fa866-a950-4b12-adff-c11fa4cf8fdc",
         "name":"cloudamqp",
         "description":"Managed HA RabbitMQ servers in the cloud",
         "requires":[

         ],
         "tags":[
            "amqp",
            "rabbitmq",
            "messaging"
         ],
         "metadata":{
            "displayName":"CloudAMQP",
            "imageUrl":"https://d33na3ni6eqf5j.cloudfront.net/app_resources/18492/thumbs_112/img9069612145282015279.png",
            "longDescription":"Managed, highly available, RabbitMQ clusters in the cloud",
            "providerDisplayName":"84codes AB",
            "documentationUrl":"http://docs.cloudfoundry.com/docs/dotcom/marketplace/services/cloudamqp.html",
            "supportUrl":"http://www.cloudamqp.com/support.html"
         },
         "dashboard_client":{
            "id": "p-mysql-client",
            "secret": "p-mysql-secret",
            "redirect_uri": "http://p-mysql.example.com/auth/create"
         },
         "plans":[
            {
               "id":"024f3452-67f8-40bc-a724-a20c4ea24b1c",
               "name":"bunny",
               "description":"A mid-sided plan",
               "metadata":{
                  "bullets":[
                     "20 GB of messages",
                     "20 connections"
                  ],
                  "costs":[
                     {
                        "amount":{
                           "usd":99.0
                        },
                        "unit":"MONTHLY"
                     },
                     {
                        "amount":{
                           "usd":0.99
                        },
                        "unit":"1GB of messages over 20GB"
                     }
                  ],
                  "displayName":"Big Bunny"
               }
            }
         ]
      }
   ]
}
```



## Resource Definitions

???swagger definition as ground truth. in the event that the text and
swagger are in conflict, swagger has precedence.???

A conforming implementation shall expose the following rest resources
with the specified HTTP methods.
  * /v2/catalog
      * `GET /v2/catalog`
  * /v2/service_instances
      * `PUT /v2/service_instances/{instance_id}`
      * `DELETE /v2/service_instances/{instance_id}`
      * `PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id}`
      * `DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}`

### catalog

#### /v2/catalog

### service instances

### service bindings



## Example flow

 1. get catalog
 2. create instance
 3. bind instance, create binding
 4. unbind instance, delete binding
 5. delete instance
