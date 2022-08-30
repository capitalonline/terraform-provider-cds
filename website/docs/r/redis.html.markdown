---
layout: "cds"
page_title: "cds: redis"
sidebar_current: "docs-cds-resource-redis"
description: |-
  Provide resources to create or delete Redis.
---

# Redis

Provide resources to create, update or delete redis.

## Example Usage

```hcl
# create redis
resource "cds_redis" "redis_example" {
    region_id         = var.region_id
    vdc_id            = var.vdc_id
    base_pipe_id      = var.base_pipe_id
    instance_name     = var.instance_name
    architecture_type = var.architecture_type
    ram               = var.ram
    redis_version     = var.redis_version
    password          = var.password
}
```
## Argument Reference
The following arguments are supported
### Redis
* `region_id` - (Required,Unmodifiable) The Region of the instance, refer to [All Region](https://github.com/capitalonline/openapi/blob/master/Redis%E6%A6%82%E8%A7%88.md#1describeregins).
* `vdc_id` - (Required,Unmodifiable) Instance belongs to the virtual data center.
* `base_pipe_id` - (Required,Unmodifiable) Vdc private network id, the haproxy instance will create id by this [Get PipeId](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1describevdc)
* `instance_name` - (Required,Unmodifiable) The name of the instance.
* `architecture_type` - (Required) Architecture name, Please refer to data.json availableDB  3:Economic master-slave ,2:The cluster,1:master-slave
* `ram` - (Required) Redis ram/G
* `redis_version` - (Required) Redis version, Please refer to data.json availableDB
* `password` - (Required) Redis password   
* `ip` - (Optional) Redis Ip address.

