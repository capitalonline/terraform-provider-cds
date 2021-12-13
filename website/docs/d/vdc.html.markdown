---
layout: "cds"
page_title: "cds: vdc"
sidebar_current: "docs-cds-resource-vdc"
description: |-
  Provide resources to create or delete VDCs.
---

# VDC

This data source provides a json file of vdc

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

data "cds_data_source_vdc" "cds_data_source_vdc" {
  id = "${cds_vdc.my_vdc.id}"
  result_output_file="my_test_path"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The ID of the VDC.
* `result_output_file` - (Required) Save the VDC information to the path.
* `vdc_list` - (Computed) The list of the VDC info , By calculation,Don't have to fill out.
