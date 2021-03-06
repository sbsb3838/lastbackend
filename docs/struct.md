Structure
=========

### Default work-flow 

Namespace management (namespace name should be unique per cluster)
- Create namespace 
- Get namespaces list
- Get namespace by name
- Get namespace info by name
- Remove namespace by name

Service management (service name should be unique per namespace)
- Create service 
- Get services list by namespace
- Get service info by namespace
- Update service info by name
- Remove service by name


### ETCD v3 storage tree

```generic
/lastbackend

integrations:
/lastbackend/vendors/<vendor>: <vendor (github|bitbucket|gitlab) object>

namespace:
/lastbackend/namespace/<namespace name>/meta: <namespace meta object> 
/lastbackend/namespace/<namespace name>/spec: <namespace spec object>
/lastbackend/namespace/<namespace name>/stat: <namespace stat object>

service:
/lastbackend/service/<service name>/meta: <service meta object>
/lastbackend/service/<service name>/spec: <service spec object>
/lastbackend/service/<service name>/stat: <service stat object>

deployment:
/lastbackend/deployment/<deployment name>/meta: <deployment meta object>
/lastbackend/deployment/<deployment name>/spec: <deployment spec object>
/lastbackend/deployment/<deployment name>/stat: <deployment stat object>

pod:
/lastbackend/pod/<pod name>/meta: <pod meta object>
/lastbackend/pod/<pod name>/spec: <pod spec object>
/lastbackend/pod/<pod name>/stat: <pod stat object>

volume:
/lastbackend/volume/<volume name>/meta: <volume meta object>
/lastbackend/volume/<volume name>/spec: <volume spec object>
/lastbackend/volume/<volume name>/stat: <volume stat object>

secret:
/lastbackend/secret/<secret name>/meta: <secret meta object>
/lastbackend/secret/<secret name>/data: <secret data object>

route:
/lastbackend/route/<route name>/meta: <route meta object>
/lastbackend/route/<route name>/spec: <route spec object>
/lastbackend/route/<route name>/stat: <route stat object>

discovery:
/lastbackend/discovery/<endpoint>: <endpoint spec object>


cluster:
/lastbackend/cluster/meta: <cluster meta object>
/lastbackend/cluster/stat: <cluster stat object>

subnets:
/lastbackend/cluster/snet/<node hostname>/meta: <subnet meta object>
/lastbackend/cluster/snet/<node hostname>/spec: <subnet spec object>

nodes:
/lastbackend/cluster/node/<node hostname>/meta: <node meta object>
/lastbackend/cluster/node/<node hostname>/stat: <node stat object>
/lastbackend/cluster/node/<node hostname>/beat: <node beat object> 
/lastbackend/cluster/node/<node hostname>/spec/pod/<pod name>: <pod spec object>
/lastbackend/cluster/node/<node hostname>/spec/route/<route name>: <route spec object>
/lastbackend/cluster/node/<node hostname>/spec/volume/<volume name>: <volume spec object>

system:

/lastbackend/system
/lastbackend/system/controller
/lastbackend/system/controller/lead
/lastbackend/system/controller/process/<system hostname>
/lastbackend/system/controller/queue/service/<service name>: <service stat object>
/lastbackend/system/controller/queue/deployment/<deployment name>: <deployment stat object>
/lastbackend/system/controller/queue/volume/<volume name>: <volume stat object>
/lastbackend/system/controller/queue/route/<route name>: <route stat object>

/lastbackend/system/scheduler
/lastbackend/system/scheduler/lead
/lastbackend/system/scheduler/process/<system hostname>
/lastbackend/system/scheduler/pod/<pod name>: <pod stat object>
/lastbackend/system/scheduler/volume/<volume name>: <volume stat object>
/lastbackend/system/scheduler/route/<route name>: <volume stat object>
```

### Structures


Namespace info object
```json
{
  "name": "demo",
  "created": "Wed Mar 01 2017 17:13:08 GMT+03:00",
  "updated": "Wed Mar 01 2017 17:13:08 GMT+03:00"
}
```

Service info object
```json
{
  "name": "mysql",
  "created": "Wed Mar 01 2017 17:13:08 GMT+03:00",
  "updated": "Wed Mar 01 2017 17:13:08 GMT+03:00"
}
```

Service config object
```json
{
  "image": "lastbackend/proxy:latest",
  "name": "",
  "replicas": 2,
  "memory": 32,
  "ports": {},
  "env": {},
  "volumes": {}
}
```

Service domains object
```json
{
  "service.lbapp.in": true,
  "service.domain.com": false
}
```

Service container object
```json
{
  "id": "59e8bce5a3032034dd84339c64fec42a8084bba90cdac6115f9456e29f646015",
  "status": "running",
  "ports" : {
    "3306/TCP": 44536
  },
  "updated": "Wed Mar 01 2017 17:13:39 GMT+03:00",
  "created": "Wed Mar 01 2017 17:13:39 GMT+03:00"
}
```
