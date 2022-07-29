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
# usage of cds_data_source_vdc
data cds_data_source_vdc "my_vdc_data" {
  id = cds_vdc.my_vdc.id
  vdc_name = cds_vdc.my_vdc.vdc_name
  result_output_file = "data.json"
}
output "public_id" {
#  output or use my_vdc's public_id
  value = data.cds_data_source_vdc.my_vdc_data.vdc_list[0].public_network[0].public_id
}

output "segment_id" {
  #  output or use my_vdc's segment_id
  value = data.cds_data_source_vdc.my_vdc_data.vdc_list[0].public_network[0].segments[0].segment_id

}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The ID of the VDC.
* `result_output_file` - (Required) Save the VDC information to the path.
* `vdc_list` - (Computed) The list of the VDC info , By calculation,Don't have to fill out.
  * `vdc_id` - (Optional) The ID of the VDC.
  * `vdc_name` - (Optional) The Name of the VDC.
  * `region_id` - (Optional) The Region of the VDC.
  * `private_network` - (Optional) The private networks of the VDC.
    * `private_id` - (Optional) The ID of the private network.
    * `status` - (Optional) The Status of the private network.
    * `name` - (Optional) The Name of the private network.
    * `unuse_ip_num` - (Optional) The num of the private network's unused ip. 
    * `segments` - (Optional) Private network segments list.
  * `public_network` - (Optional)  The public networks of the VDC.
    * `public_id` - (Optional) The ID of the public network.
    * `status` - (Optional) The Status of the public network.
    * `qos` - (Optional) The Qos of the public network.
    * `name` - (Optional) The Name of the public network.
    * `unuse_ip_num` - (Optional) The Number of the public network's unused ip.
    * `segments` - (Optional) The Segments of the public network.
      * `mask` - (Optional) The Mask of the segment.
      * `gateway` - (Optional) The Gateway of the segment.
      * `segment_id` - (Optional) The ID of the segment.
      * `address` - (Optional) The Address of the segment.
