---
layout: "cds"
page_title: "cds: security group"
sidebar_current: "docs-cds-resource-securitygroup"
description: |-
  Provide resources to create or delete security groups.
---

# Security Group

This data source provides a json file of security group

## Example Usage

```hcl
resource "cds_security_group" "security_group_1" {
  name = "test_tf_new_zz"
  description = "New security group"
  type ="public"
  rule  {
    action        = "1"
    description   = "tf_rule_test"
    targetaddress = "120.78.170.188/28;120.78.170.188/28;120.78.170.188/28"
    targetport    = "70;90;8"
    localport     = "800"
    direction     = "all"
    priority      = "11"
    protocol      = "TCP"
    ruletype      = "ip"
  }
}

data "cds_data_source_security_group" "cds_data_source" {
  id = "${cds_security_group.security_group_1.id}"
  result_output_file="my_test_path"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The ID of the security group.
* `result_output_file` - (Required) Save all information for security group and security group rules to this path.