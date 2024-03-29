---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cds_mysql_readonly Resource - terraform-provider-cds"
subcategory: ""
description: |-
  Mysql read-only instance
---

# cds_mysql_readonly (Resource)

Mysql read-only instance

## Example usage

```hcl

resource "cds_mysql_readonly" "readonly1" {
    instance_uuid = cds_mysql.mysql_example.id
    instance_name = "readonly"
#    You can find paas_goods_id in data.json.
#    The field name is available_read_only_config
    paas_goods_id = 1680
#    test_group_id = 0
    disk_type = "high_disk"
    disk_value = "500"
}

```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `disk_type` (String) Disk type. [View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#15createreadonlydbinstance)
- `disk_value` (Number) Disk value. The size of disk. [View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#15createreadonlydbinstance)
- `instance_name` (String) Instance name. Read only instance name.[View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#15createreadonlydbinstance)
- `instance_uuid` (String) Instance uuid. Mysql instance uuid. [View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#15createreadonlydbinstance)

### Optional

- `paas_goods_id` (Number) Paas goods id.[View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#15createreadonlydbinstance)
- `subject_id` (Number) Subject ID.
- `test_group_id` (Number) Test group id. [View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#15createreadonlydbinstance)

### Read-Only

- `id` (String) The ID of this resource.
