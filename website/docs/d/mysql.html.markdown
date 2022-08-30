---
layout: "cds"
page_title: "cds: Mysql"
sidebar_current: "docs-cds-resource-mysql"
description: |-
  Provide data for Mysql.
---

# Mysql

This data source provides a json file of instances in a Mysql

## Example Usage

```hcl
data cds_data_source_mysql "mysql_data" {
    region_id = "CN_Beijing_A"
    result_output_file = "data.json" // availableDB, instances, regions
}
```

## Argument Reference

The following arguments are supported:

# MySQL
* `region_id` - (Required) The Mysql region.
* `instance_uuid` - (Optional) The Mysql instance uuid to filter.
* `instance_name` - (Optional) The Mysql instance name to filter.
* `ip` - (Optional) The Mysql instance ip to filter.
* `result_output_file` - (Required) Save all instance information to the path.
* `readonly_instances` - (Optional)  List of readonly instances.
  * `instance_name` - (Optional) Name of the readonly instance.
  * `disk_type` - (Optional) Disk type of the readonly instance.
  * `disk_value` - (Optional) Disk size of the readonly instance.
