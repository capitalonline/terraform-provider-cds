---
layout: "cds"
page_title: "cds: public network"
sidebar_current: "docs-cds-resource-public-network"
description: |-
Provide resources to create or delete public Networks.
---

# VDC

Provide resources to create or delete public Networks.

## Example Usage

```hcl
# create public network
resource "cds_public_network" "pb1"{
  ip_num          = 4
  qos            = 10
  name          = "terraform-copy"
  float_bandwidth = 200
  billing_method  = "BandwIdth"
  auto_renew      = 1
  type           = "Bandwidth_Multi_ISP_BGP"
  vdc_id = cds_vdc.my_vdc.id
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required,Unmodifiable) The name of the public network.
* `ipnum` -(Required,Unmodifiable) Available ip quantity.
* `qos` - (Required) The bandwidth(Mbps) of the public subnet.
* `name` - (Required,Unmodifiable) The name of the public subnet.
* `floatbandwidth` - (Required,Unmodifiable) Maximum limit bandwidth.
* `billingmethod` - (Required,Unmodifiable) Billing method.
* `autorenew` - (Required) Whether to automatically renew,1 is the automatic renewal fee (default), 0 is not automatic renewal.
* `type` - (Required,Unmodifiable) The type of the public subnet, refer to [All Type](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#%E5%B8%A6%E5%AE%BD%E7%B1%BB%E5%9E%8B).