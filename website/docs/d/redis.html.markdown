---
layout: "cds"
page_title: "cds: Redis"
sidebar_current: "docs-cds-resource-redis"
description: |-
  Provide data for Redis.
---

# Redis

This data source provides a json file of instances in a Redis

## Example Usage

```hcl
data cds_data_source_redis "redis_data" {
    region_id           = "CN_Beijing_A"
    instance_uuid       = "XXXXXXX"
    instance_name       = "XXXXXXX"
    ip                  = "XXXXXXX"
    result_output_file  = "data.json" // availableDB, instances, regions
    }
```

## Argument Reference

The following arguments are supported:

# Redis
* `region_id` - (Required) The Redis region.
* `instance_uuid` - (Optional) The Redis instance uuid to filter.
* `instance_name` - (Optional) The Redis instance name to filter.
* `ip` - (Optional) The Redis instance ip to filter.
* `result_output_file` - (Required) Save all instance information to the path.
