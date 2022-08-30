---
layout: "cds"
page_title: "cds: mysql readonly"
sidebar_current: "docs-cds-resource-mysql-readonly"
description: |-
  Provide resources to create Mysql read-only instance.
---

# Mysql

Provide resources to create, update or delete Mysql.

## Example Usage

```hcl
# create mysql
resource "cds_mysql_readonly" "mysql_readonly_example" {
    instance_name     = "readonly_example"
    paas_goods_id     = 6707                   // readonly instance product specification, which must be greater than or equal to the primary instance specification.
    test_group_id     = 0                      // whether to use test group billing. 
    disk_type         = "high_disk"            
    disk_value        = 400                    // disk size. The disk size of a read-only instance cannot be smaller than that of the primary instance.
    amount            = 1                      // quantity purchased Up to three purchases at a time
}

```
## Argument Reference
The following arguments are supported

* `instance_uuid` - (Optional,Unmodifiable) The id of the instance. Known after created.
* `instance_name` - (Required,Unmodifiable) The name of the instance.
* `paas_goods_id` - (Required,Unmodifiable) Readonly instance commodity specifications.
* `test_group_id` - (Optional,Unmodifiable) Whether to use test group billing.
* `disk_type`     - (Required,Unmodifiable) Disk type.
* `disk_value`    - (Required,Unmodifiable) Disk size. The disk size of a read-only instance cannot be smaller than that of the primary instance.
* `amount`        - (Optional,Unmodifiable) Quantity purchased Up to three purchases at a time