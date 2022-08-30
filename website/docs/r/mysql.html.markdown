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
    // The UTC time zone of the instance, the default value varies according to the region.  
    // This field is valid only at creation.
    time_zone = "+08:00"
  
    // Modify parameters of the instance.
    parameters = {
      name="connect_timeout"          // parameter name
      value="100"                     // parameter value
    }
    // Create a backup.
    backup = {
      backup_type = "logical-backup" // backup type,'physical-backup' or 'logical-backup.'
      desc = "test"
      db_list = ["db1","db2"]        // databases to be backed up.
    }
    // Set auto backup policy.
    data_backups = {
      time_slot = "00:00-01:00"      // the backup time period starts and ends on the hour with an interval of one hour.
      date_list = [0,1,2,3,4]        // 0 is Sunday, 1 is Monday, and so on
      sign = 0                       // auto backup switch, off: 0, on: 1
    }
}

# create mysql param
data cds_data_source_mysql "mysql_data" {
    region_id = "CN_Beijing_A"
    result_output_file = "data.json" // availableDB, instances, regions
}


# output instances ip
output "instance_ip" {
  value = cds_mysql.mysql_example.ip
}
```
## Argument Reference
The following arguments are supported
### Haproxy
* `instance_uuid` - (Optional,Unmodifiable) The id of the instance. Known after created.
* `instance_name` - (Required,Unmodifiable) The name of the instance.
* `region_id` - (Required,Unmodifiable) The Region of the instance, refer to [All Region](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#1describeregions).
* `vdc_id` - (Required,Unmodifiable) Instance belongs to the virtual data center.
* `base_pipe_id` - (Required,Unmodifiable) Vdc private network id, the haproxy instance will create id by this [Get PipeId](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1describevdc)
* `cpu` - (Required) Mysql cpu/C
* `ram` - (Required) Mysql ram/G
* `disk_type` - (Required,Unmodifiable) Mysql disk type. It can only be the same as the type at the beginning of purchase. You cannot add multiple types of disks to an instance. For example, you can only add high-performance disks later
* `disk_value` - (Required) Mysql disk value, Please refer to data.json availableDB
* `mysql_version` - (Required) Mysql version, Please refer to data.json availableDB
* `architecture_type` - (Required) Architecture name, Please refer to data.json availableDB 0:basic edition|1:master-slave edition
* `compute_type` - (Required) Compute name, Please refer to data.json availableDB 0:common type
* `ip` - (Optional) The IP of the instance. Known after created.
* `parameters` - (Optional) The settings of the DB instance.
  * `name` - (Optional) Parameter name.
  * `value` - (Optional) Parameter value. 
* `time_zone` - (Optional) The UTC time zone of the instance.Valid values: one of ["-12:00","-11:00","-10:00","-09:00","-08:00","-07:00","-06:00","-05:00","-04:00","-03:00","-02:00","-01:00","+00:00","+01:00","+02:00","+03:00","+04:00","+05:00","+05:30","+06:00","+07:00","+08:00","+09:00","+10:00","+11:00","+12:00","+13:00",]
* `backup` - (Optional) Create backup.
  * `backup_type` - (Optional) Backup type. Valid values: one of ["physical-backup","logical-backup"]
  * `desc` - (Optional) Description of backup.
  * `db_list` - (Optional)  Databases which will be backup.
* `data_backups` - (Optional) Backup policy of the instance.
  * `time_slot` - (Optional) Backup time period. Valid values: one of ["00:00-01:00","01:00-02:00","02:00-03:00","03:00-04:00","04:00-05:00","05:00-06:00","06:00-07:00","07:00-08:00","08:00-09:00","09:00-10:00","10:00-11:00","11:00-12:00","12:00-13:00","13:00-14:00","14:00-15:00","15:00-16:00","16:00-17:00","17:00-18:00","18:00-19:00","19:00-20:00","20:00-21:00","21:00-22:00","22:00-23:00","23:00-24:00"]
  * `date_list` - (Optional) The backup cycle.Valid values: some of ["0","1","2","3","4","5","6"]
  * `sign` - (Optional) Automatic backup switch. Value 0: off,value 1:on.