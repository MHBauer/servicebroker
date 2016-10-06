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

## <a id='possible-errors'></a>Possible Errors ##

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
