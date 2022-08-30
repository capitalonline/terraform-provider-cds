---
layout: "cds"
page_title: "cds: vdc"
sidebar_current: "docs-cds-resource-vdc"
description: |-
  Provide resources to create or delete VDCs.
---

# VDC

Provide resources to create or delete VDCs. 

## Example Usage

```hcl
# create vdc
resource "cds_vdc" "example" {
  vdc_name = "example"
  region_id = "CN_Beijing_A"
  public_network = {
    "ipnum" = 4
    "qos" = 20
    "name" = "example-public-network"
    "floatbandwidth" = 200
    "billingmethod" = "BandwIdth"
    "autorenew" = 1
    "type" = "Bandwidth_BGP"
  }
}

# update vdc
# The following fields can be modified: vdc_name、add_public_ip、delete_public_ip、public_network
resource "cds_vdc" "example" {
#  Modified field vdc_name will rename the vdc.
  vdc_name = "new-vdc-name"

#  Modified field 'add_public_ip' will add a public IP segment to the purchased public network.
#  The value is the number of Ip addresses to be purchased.
  add_public_ip = 4
#  Modified field 'delete_public_ip' will delete an IP segment under the Vdc public network.
#  The value is the IP segment id will be deleted.
  delete_public_ip = "xxx"
}


```

## Argument Reference

The following arguments are supported:

* `vdc_name` - (Required,Unmodifiable) The name of the VDC.
* `region_id` - (Required,Unmodifiable) The Region of the VDC, refer to [All Region](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#%E8%8A%82%E7%82%B9%E5%90%8D%E7%A7%B0).
* `public_network` - (Optional) Create a public subnet for the VDC.
  * `ipnum` -(Required,Unmodifiable) Available ip quantity.
  * `qos` - (Required) The bandwidth(Mbps) of the public subnet, .
  * `name` - (Required,Unmodifiable) The name of the public subnet.
  * `floatbandwidth` - (Required,Unmodifiable) Maximum limit bandwidth.
  * `billingmethod` - (Required,Unmodifiable) Billing method.
  * `autorenew` - (Required) Whether to automatically renew,1 is the automatic renewal fee (default), 0 is not automatic renewal.
  * `type` - (Required,Unmodifiable) The type of the public subnet, refer to [All Type](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#%E5%B8%A6%E5%AE%BD%E7%B1%BB%E5%9E%8B).
* `add_public_ip` - (Optional) Add public ip num. Valid value: in [4,8,16,32,64].
* `delete_public_ip` - (Optional) Delete public Ip segment id.
* `public_id` - (Optional,Unmodifiable) Public Network id.
