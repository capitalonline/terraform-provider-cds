---
layout: "cds"
page_title: "CapitalCloud: capitalcloud_security_group_rule"
sidebar_current: "docs-capitalcloud-resource-security_group_rule"
description: |-
  Provides a resource to create security group rule.
---

# CDS Private Subnet For VDC

Provides a resource to create Private Subnet for VDC.

## Example Usage

```hcl
# create private subnet
resource "cds_private_subnet" "my_private_subnet_1" {
  vdc_id = "vdc_id"
  name = "private_name"
  type = "auto"
  address = "192.168.0.0"
  mask = "16"
}

```

## Argument Reference

The following arguments are supported:

* `vdc_id` - (Required,Unmodifiable) The ID of the VDC.
* `name` - (Required,Unmodifiable) The Name of the private subnet.
* `type` - (Required,Unmodifiable) The Type of the private subnet, allowed value: auto, mu.
* `address` - (Required,Unmodifiable) The address of the private subnet.
* `mask` - (Required,Unmodifiable) The mask of the private subnet.