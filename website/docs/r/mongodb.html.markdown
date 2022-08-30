---
layout: "cds"
page_title: "cds: mongodb"
sidebar_current: "docs-cds-resource-mongodb"
description: |-
  Provide resources to create or delete mongodb.
---

# Mongodb

Provide resources to create, update or delete mongodb.

## Example Usage

```hcl
# create mongodb
resource "cds_mongodb" "mongodb_example" {
    region_id         = var.region_id
    vdc_id            = var.vdc_id
    base_pipe_id      = var.base_pipe_id
    instance_name     = var.instance_name
    cpu               = var.cpu
    ram               = var.ram
    disk_type         = var.disk_type
    disk_value        = var.disk_value
    password          = var.password
    mongodb_version   = var.mongodb_version
}
```
## Argument Reference
The following arguments are supported
### Mongodb
* `region_id` - (Required,Unmodifiable) The Region of the instance, refer to [All Region](https://github.com/capitalonline/openapi/blob/master/%E6%96%B0%E7%89%88MongoDB%E6%A6%82%E8%A7%88.md#1describezones).
* `vdc_id` - (Required,Unmodifiable) Instance belongs to the virtual data center.
* `base_pipe_id` - (Required,Unmodifiable) Vdc private network id, the haproxy instance will create id by this [Get PipeId](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1describevdc)
* `instance_name` - (Required,Unmodifiable) The name of the instance.
* `cpu` - (Required) mongodb cpu/C
* `ram` - (Required) mongodb ram/G
* `disk_type` - (Required) (Required,Unmodifiable) Mongodb disk type, Currently supported ssd_disk
* `disk_value` - (Required) Mongodb disk value 
* `password` - (Required) mongodb password   
* `mongodb_version` - (Required) mongodb version, select by 4.0.3 , 3.6.7 , 3.2.21
* `ip` - (Optional) The ip address of the instance. Known after created.


