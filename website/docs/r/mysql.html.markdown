---
layout: "cds"
page_title: "cds: mysql"
sidebar_current: "docs-cds-resource-mysql"
description: |-
  Provide resources to create or delete Mysql.
---

# Mysql

Provide resources to create, update or delete Mysql.

## Example Usage

```hcl
# create mysql
resource "cds_mysql" "mysql_example" {
    region_id         =  var.region_id
    vdc_id            = var.vdc_id
    base_pipe_id      = var.base_pipe_id
    instance_name     = var.instance_name
    cpu               = var.cpu
    ram               = var.ram
    disk_type         = var.disk_type
    disk_value        = var.disk_value
    mysql_version     = var.mysql_version 
    architecture_type = var.architecture_type //0:basic edition|1:master-slave edition
    compute_type      = var.compute_type      //0:common type
}

# create mysql param
data cds_data_source_mysql "mysql_data" {
    region_id = "CN_Beijing_A"
    result_output_file = "data.json" // availableDB, instances, regions
}
```
## Argument Reference
The following arguments are supported
### Haproxy
* `instance_name` - (Required,Unmodifiable) The name of the instance.
* `region_id` - (Required,Unmodifiable) The Region of the instance, refer to [All Region](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#1describeregions).
* `vdc_id` - (Required,Unmodifiable) Instance belongs to the virtual data center.
* `base_pipe_id` - (Required,Unmodifiable) Vdc private network id, the haproxy instance will create id by this [Get PipeId](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1describevdc)
* `cpu` - (Required) Mysql cpu/C
* `ram` - (Required) Mysql ram/G
* `disk_type` - (Required,Unmodifiable) Mysql disk type. It can only be the same as the type at the beginning of purchase. You cannot add multiple types of disks to an instance. For example, you can only add high-performance disks later
* `disk_value` - (Required) Mysql disk value, Please refer to data.json availableDB
* `mysql_version` - (Required) Mysql version, Please refer to data.json availableDB
* `architecture_name` - (Required) Architecture name, Please refer to data.json availableDB
* `compute_name` - (Required) Compute name, Please refer to data.json availableDB
