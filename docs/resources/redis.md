---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cds_redis Resource - terraform-provider-cds"
subcategory: ""
description: |-
  Redis instance.
---

# cds_redis (Resource)

Redis instance. 

## Example usage

```hcl

resource "cds_redis" "redis_example" {
    region_id         = "CN_Beijing_A"
    vdc_id            = "xxx"
    base_pipe_id      = "xxx"
    instance_name     = "redis_test"
    architecture_type = 3
    ram               = 4
    redis_version     = "2.8"
    password          = "password"
}

```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `architecture_type` (Number) Architecture type. [View Document](https://github.com/capitalonline/openapi/blob/master/Redis%E6%A6%82%E8%A7%88.md#2describeavailabledbconfig)
- `base_pipe_id` (String) Base pipe id.
- `instance_name` (String) Instance name.
- `password` (String) Password.
- `ram` (Number) Ram.[View Document](https://github.com/capitalonline/openapi/blob/master/Redis%E6%A6%82%E8%A7%88.md#2describeavailabledbconfig)
- `redis_version` (String) Redis version.[View Document](https://github.com/capitalonline/openapi/blob/master/Redis%E6%A6%82%E8%A7%88.md#2describeavailabledbconfig)
- `region_id` (String) Region id.
- `vdc_id` (String) Vdc id.

### Optional

- `subject_id` (Number) Subject id.

### Read-Only

- `id` (String) The ID of this resource.
- `ip` (String) Ip.
