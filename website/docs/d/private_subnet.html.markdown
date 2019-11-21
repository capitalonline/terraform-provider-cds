---
layout: "cds"
page_title: "CapitalCloud: capitalcloud_security_group_rule"
sidebar_current: "docs-capitalcloud-resource-security_group_rule"
description: |-
  Provides a resource to create security group rule.
---

# CDS Private Subnet For VDC

This data source provides a json file of private subnet in a vdc

## Example Usage

```hcl
resource "cds_vdc" "my_vdc" {
  vdc_name = "test_tf_bew_zz"
  region_id = "CN_Beijing_A"
  public_network = {
    "ipnum" = 4
    "qos" = 20
    "name" = "test-accPubNet002"
    "floatbandwidth" = 200
    "billingmethod" = "BandwIdth"
    "autorenew" = 1
    "type" = "Bandwidth_BGP"
  }
}
resource "cds_private_subnet" "my_private_subnet_1" {
  vdc_id = cds_vdc.my_vdc.id
  name = "private_1"
  type = "auto"
  address = "192.168.0.0"
  mask = "16"
}
data "cds_data_source_private_subnet" "cds_data_source" {
  vdc_id = "9e53b1e0-c49a-4827-bea0-af8cb4857b30"
  result_output_file="my_test_path"
}

```

## Argument Reference

The following arguments are supported:

* `vdc_id` - (Required) The ID of the VDC.
* `result_output_file` - (Required) Save all private subnets under vdc to this path.