---
layout: "cds"
page_title: "cds: public network"
sidebar_current: "docs-cds-resource-public-network"
description: |-
Provide resources to create or delete public Networks.
---

# VDC

This data source provides a json file of public network

## Example Usage

```hcl
data "cds_data_source_public_network" "pbn"{
  vdc_id = cds_vdc.my_vdc.id
  public_id = cds_public_network.pb1.id
}
```

## Argument Reference

The following arguments are supported:
* `vdc_id` - (Required) The ID of the vdc.
* `public_id` - (Required) The ID of the public network.
* `status` - (Optional) The Status of the public network.
* `qos` - (Optional) The Qos of the public network.
* `name` - (Optional) The Name of the public network.
* `unuse_ip_num` - (Optional) The Number of the public network's unused ip.
* `segments` - (Optional) The Segments of the public network.
  * `mask` - (Optional) The Mask of the segment.
  * `gateway` - (Optional) The Gateway of the segment.
  * `segment_id` - (Optional) The ID of the segment.
  * `address` - (Optional) The Address of the segment.
