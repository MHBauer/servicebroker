---
title: Managing Service Brokers
owner: Core Services
---

Mostly Irrelevant.

# suggested usage 

## update

Updating a broker is how to ingest changes a broker author has made
into Cloud Foundry. Similar to adding a broker, update causes Cloud
Controller to fetch the catalog from a broker, validate it, and update
the Cloud Controller database with any changes found in the catalog.

## removing

Removing a service broker will remove all services and plans in the
broker's catalog from the Cloud Foundry Marketplace.

Attempting to remove a service broker MUST fail if there are service
instances for any service plan in its catalog. When planning to shut
down or delete a broker, make sure to remove all service instances
first. Failure to do so will leave <a href="api.md#orphans">orphaned
service instances</a> in the Cloud Foundry database. If a service
broker has been shut down without first deleting service instances,
you can remove the instances with the CLI; see <a
href="#purge-service">Purge a Service</a>.

## Purging leftover services

If a service broker has been shut down or removed without first
deleting service instances from Cloud Foundry, you will be unable to
remove the service broker or its services and plans from the
Marketplace. In development environments, broker authors often destroy
their broker deployments and need a way to clean up the Cloud
Controller database.

The following command will delete a service offering, all of its
plans, as well as all associated service instances and bindings from
the Cloud Controller database, without making any API calls to a
service broker. For services from v1 brokers, you must provide a
provider with `-p PROVIDER`. Once all services for a broker have been
purged, the broker can be removed normally.

```
$ cf purge-service-offering v1-test -p pivotal-software
Warning: This operation assumes that the service broker responsible for this
service offering is no longer available, and all service instances have been
deleted, leaving orphan records in Cloud Foundry's database. All knowledge of
the service will be removed from Cloud Foundry, including service instances and
service bindings. No attempt will be made to contact the service broker; running
this command without destroying the service broker will cause orphan service
instances. After running this command you may want to run either
delete-service-auth-token or delete-service-broker to complete the cleanup.

Really purge service offering v1-test from Cloud Foundry? y
OK
```

### <a id='purge-service-instance'></a>Purge a Service Instance###

The following command will delete a single service instance, its service bindings and its service keys from the Cloud Controller database, without making any API calls to a service broker.
This can be helpful in instances a Service Broker is not conforming to the Service Broker API and not returning a 200 or 410 to requests to delete the service instance.

<pre class="terminal">
$ cf purge-service-instance mysql-dev
WARNING: This operation assumes that the service broker responsible for this
service instance is no longer available or is not responding with a 200 or 410,
and the service instance has been deleted, leaving orphan records in Cloud
Foundry's database. All knowledge of the service instance will be removed from
Cloud Foundry, including service bindings and service keys.

Really purge service instance mysql-dev from Cloud Foundry?> y
Purging service mysql-dev...
OK
</pre>

`purge-service-instance` requires cf-release v218 and cf CLI 6.14.0.

## <a id='possible-errors'></a>Possible Errors ##

If incorrect basic auth credentials are provided:

<pre class="terminal">
Server error, status code: 500, error code: 10001, message: Authentication
failed for the service broker API.
Double-check that the username and password are correct:
    http://github-broker.a1-app.example.com/v2/catalog
</pre>

If you receive the following errors, check your broker logs.
You may have an internal error.

<pre class="terminal">
Server error, status code: 500, error code: 10001, message:
    The service broker response was not understood

Server error, status code: 500, error code: 10001, message:
    The service broker API returned an error from
    http://github-broker.a1-app.example.com/v2/catalog: 404 Not Found

Server error, status code: 500, error code: 10001, message:
    The service broker API returned an error from
    http://github-broker.primo.example.com/v2/catalog: 500 Internal Server Error
</pre>

If your broker's catalog of services and plans violates validation of presence,
uniqueness, and type, you will receive meaningful errors.

<pre class="terminal">
Server error, status code: 502, error code: 270012, message: Service broker catalog is invalid:
Service service-name-1
  service id must be unique
  service description is required
  service "bindable" field must be a boolean, but has value "true"
  Plan plan-name-1
    plan metadata must be a hash, but has value [{"bullets"=>["bullet1", "bullet2"]}]
</pre>
