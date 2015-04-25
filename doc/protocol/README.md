zeroctl protocol v0.0.1

# Basic concepts

## Service providers

are providing defined serice to components of the system. Service provider **have** to acknowledge accepting task to client if client requests it.

Shortened to **SP** for rest of the document.

Each service have name and optional *instances* and *sub-instances*

### Instances
Instances are for dividing services of same type into logical instances.

For example service.nmap might be divided into service.nmap.dc1 and service.nmap.dc2 to give client ability to choose a location without sending request to specific node

Or job manager could have `service.jobs.bulk` and `service.jobs.important` to divide work between slower and faster workers

It can be also used to separate resources of different projects and then futher subdivide it inside a given project for example multiple projects with multiple varnish cache pools: `service.cache-invalidator.project1.varnish-gfx`, `service.cache-invalidator.project2.varnish-mm`

### Accepting a task

SP **must** generate name of the task when accepting, it **must** be unique

task **should** be named based on client node name and uniquely generated ID, for example

`web1_example_com-parser.1382439c-d46b-425f-88a1-98e66d462da1`

If client sends `on-start` header with return path, SP **must** send an ack with task ID to client, or nak if job is invalid.

If client sends `on-finish` header with return path, SP **may** send an ack when job is finished.

If client sends `on-finish` header with return path and SP does not support sending `on-finish` ack, it **must** respond with `no-on-finish: 1` in `on-start` ack to inform client it wont be sending `on-finish` even if it was requested

## Nodes

ther are 2 types of nodes:

### Persistent

Emit heartbeat and are discoverable via discovery.* namespace.

### Transient

Nodes not registered via discovery service, they can generate/consume events and use existing services but not provide them

# Protocol

Message is divided in 3 parts:

* Routing key - used in low-level routing and rough qualification of message type
* Header - passing basic parameters, and in advanced filtering
* Body - container for service requests/responses

## Routing key

String of 255 ascii chars used to address components of the system. It is formatted by parts separated by dots.
First part is a type of [endpoint](endpoints.md)
Second part **should** be a name of underlying service and each subsequent part **should** further subdivide service.

So for example if we need to address 6 clusters of service distibuted into 2 DCs, routing key should look like `service.batch-jobs.dc1.cluster3`

### Filtering

each dotted part of routing key can be filtered using 2 filters:

* `*`(star) - means "substiture one word"
* `#` (hash) - means "substitute more than one word"

so `service.img.*.png` will match `service.img.crop.png` but not `service.img.png` or `service.img.dc1.crop.png`

and `service.#.png` will match all of those

### Routing

Routing engine **must** be able to route using prefixes (so `prefix.#` filter equivalent) and **must** be able to filter (at minimum, receive and ignore not matching) via it

## Header

A set of key=>value pairs with key being 250 ascii string and value limited by header size.

Minimum header size is 64KiB of json-encoded data (even if underlying transport have to use less/more data to accomplish it).

It should be encoded on transport level (so client only passes a hash of data).

Transport **may** add new headers but **must not** modify passed ones. If transport need thos headers for it's own use it should prefix received headers on send, and remove that prefix on receive.

Transport **can** add keys on receive but they **must** be contained under `_transport-` prefix

### required keys

* `node-name` - 1-256 byte UTF-8 node name. Human readable, preferably in fqdn-appname form
* `node-uuid` - 32 byte node UUID
* `ts` - unixtime, can be s/ms/us accuracy
* `sha256` - checksum of data part

### reserved keys prefix

* `_transport-` - transport specific info (like used auth/id)
* `_sec-` - security-related parts of protocol like signature **should not** be generated/verified by client/service directly but by library handling the communication

## Body

Body **should** be json encoded message. Encoding **must** be done before put into transport, transport **must not** encode it and should treat it as binary blob. Blob **must** be checksummed with checksum in `sha256` header

Minmum size of body transport **must** transfer is 4MiB

# Implementation details

* [endpoints](endpoints.md)
* [protocol -> wire protocol mappings](mapping.md)
