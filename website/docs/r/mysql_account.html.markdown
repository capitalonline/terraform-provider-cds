---
layout: "cds"
page_title: "cds: mysql account"
sidebar_current: "docs-cds-resource-mysql-account"
description: |-
  Provide resources to create Mysql account.
---

# Mysql

Provide resources to create, update or delete Mysql account.

## Example Usage

```hcl
# create mysql
resource "cds_mysql_account" "mysql_readonly_example" {
    account_name      = "test-account"
    password          = "*****"
    account_type      = "Super"
    description       = "create a new account"
#    Modify account privilege.
    operations        = [{
      db_name   = "db-example"
      privilege = "DMLOnly"
    }]                    
}

```
## Argument Reference
The following arguments are supported

* `instance_uuid` - (Required,Unmodifiable) The id of the instance. 
* `account_name` - (Required,Unmodifiable) The name of the account.
* `password` - (Required,Unmodifiable) The password of the account.
* `account_type` - (Optional,Unmodifiable) The type of account. One of the list["Super","Normal"].An instance can only have one high-privileged account.
* `description`     - (Required,Unmodifiable) Description of the account.
* `operations`    - (Required) List of db privilege.
  * `db_name` - (Required) Database name to be authorized.
  * `privilege` - (Required) Database corresponding account permissions.One of the list ["ReadWrite","DMLOnly","ReadOnly","DDLOnly"]