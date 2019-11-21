---
layout: "cds"
page_title: "cds: security group"
sidebar_current: "docs-cds-resource-securitygroup"
description: |-
  Provide resources to create or delete security groups.
---

# Security Group

Provide resources to create, update or delete security groups.

## Example Usage

```hcl
# create security group
resource "cds_security_group" "security_group_example" {
  name = "example"
  description = "New security group"
  type ="public"
  rule  {
    action        = "1"
    description   = "description"
    targetaddress = "120.78.170.188/28;120.78.170.188/28;120.78.170.188/28"
    targetport    = "70;90;8"
    localport     = "800"
    direction     = "all"
    priority      = "11"
    protocol      = "TCP"
    ruletype      = "ip"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the security group.
* `description` - (Required) The Description of the security group.
* `type` - (Required) The Type of the security group, private or public.
* `rule` - (Required) Create a rule for the security group.
  * `action` - (Required) Network connectivity policy within the security group. 0 means prohibited, 1 means allowed.
  * `description` - (Required) Rule description information.
  * `targetaddress` - (Required)Target ip address.
  * `targetport` - (Required)Target port. Where 0 represents a collection of all ports.
  * `localport` - (Required)Source port open or disabled by this unit.
  * `direction` - (Required)Rule access rights. Currently only supports the addition and binding of two-way rules.
  * `priority` - (Required)Set the rule priority in the security group, ranging from 1 to 100.
  * `protocol` - (Required)Transport layer protocol. The optional parameters are: ICMP, TCP, UDP. Note: If the protocol is ICMP, you do not need to pass the TargetPort and LocalPort parameters..
  * `ruletype` - (Required)Set the rule type. Options: mac or ip.
