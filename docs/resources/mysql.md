---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cds_mysql Resource - terraform-provider-cds"
subcategory: ""
description: |-
  Mysql instance.
---

# cds_mysql (Resource)

Mysql instance.

## Example usage

```hcl

resource "cds_mysql" "mysql_example" {
    region_id         = "CN_Beijing_E"
    vdc_id            = "xxx"
    base_pipe_id      = "xxx"
    instance_name     = "mysql-instance"
    cpu               = 2
    ram               = 4
    disk_type         = "ssd_disk"
    disk_value        = 100
    mysql_version     = "5.7"
    architecture_type = 0
    compute_type      = 0
    # Set mysql instance parameters
    parameters        = [
        {
          name  = "back_log"
          value = "8888"
        }
    ]
    # set mysql instance time_zone
    time_zone = "+08:00"

    #  Set  backup
    backup = {
        backup_type = "logical-backup"
        desc = "backup"
        db_list = "db1,db2"
    }

    #  Set auto backup policy
    data_backups = {
        time_slot="00:00-01:00"
    #   Split databases with ","
        date_list="1,2,3"
        sign = 0
    }
}

```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `architecture_type` (Number) Architecture type :0.basic edition 、1.master-slave edition
- `base_pipe_id` (String) Base pipe id
- `compute_type` (Number) Compute type: 0.common type
- `cpu` (Number) Cpu num
- `disk_type` (String) Disk type: ssd_disk、high_disk
- `disk_value` (Number)
- `instance_name` (String) Instance name
- `mysql_version` (String) Mysql version
- `ram` (Number) Ram num
- `region_id` (String) Region id
- `vdc_id` (String) Vdc id

### Optional

- `backup` (Map of String) Create db instance backup.[View document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#16createbackup)
- `data_backups` (Map of String) Data backup. [View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#20modifydbbackuppolicy)
- `instance_uuid` (String)
- `parameters` (Block List) Mysql instance parameters (see [below for nested schema](#nestedblock--parameters))
- `subject_id` (Number) Subject ID.
- `time_zone` (String) Time zone.

### Read-Only

- `id` (String) The ID of this resource.
- `ip` (String) Instance ip

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `name` (String)
- `value` (String)
