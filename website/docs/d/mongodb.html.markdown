---
layout: "cds"
page_title: "cds: Mongodb"
sidebar_current: "docs-cds-resource-mongodb"
description: |-
  Provide data for Mongodb.
---

# Mongodb

This data source provides a json file of instances in a Mongodb

## Example Usage

```hcl
data cds_data_source_mongodb "mongodb_data" {
    region_id           = "CN_Beijing_A"
    instance_uuid       = "XXXXXXX"
    instance_name       = "XXXXXXX"
    ip                  = "XXXXXXX"
    result_output_file  = "data.json"
    }
```

## Argument Reference

The following arguments are supported:

# Mongodb
* `region_id` - (Required) The Mongodb region.
* `instance_uuid` - (Optional) The Mongodb instance uuid to filter.
* `instance_name` - (Optional) The Mongodb instance name to filter.
* `ip` - (Optional) The Mongodb instance ip to filter.
* `result_output_file` - (Required) Save all instance information to the path.
