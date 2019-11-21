---
layout: "cds"
page_title: "cds: instance"
sidebar_current: "docs-cds-resource-instance"
description: |-
  Provide resources to create or delete VDCs.
---

# Instance

This data source provides a json file of instances in a vdc

## Example Usage

```hcl
resource "cds_vdc" "my_vdc" {
  vdc_name = "Terraform(using)"
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
  addres = "192.168.0.0"
  mask = "16"
}

resource "cds_security_group" "security_group_1" {
  name = "test_tf_new_zz"
  description = "New security group"
  type ="private"
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

resource "cds_instance" "my_instance" {
  instance_name       = "test_zz_002"
  region_id           = "CN_Beijing_A"
  image_id            = "Ubuntu_16.04_64"
  instance_type       = "high_ccs"
  cpu                 = 4
  ram                 = 4
  vdc_id              = cds_vdc.my_vdc.id
  password            = "123abc,.;"
  public_ip           = "auto"
  private_ip          = {
    "private_id" = cds_private_subnet.my_private_subnet_1.id,
    "address" = "auto"
  }
  data_disks {
    size  =  100
    type  =  "high_disk"
  }
  security_group_binding {
    type = "private"
    subnet_id = cds_private_subnet.my_private_subnet_1.id,
    security_group_id = cds_security_group.security_group_1.id
  }
}
data "cds_data_source_instance" "cds_data_source" {
  vdc_id = cds_vdc.my_vdc.id
  result_output_file="my_test_path"
}
```

## Argument Reference

The following arguments are supported:

* `vdc_id` - (Required) The ID of the VDC.
* `result_output_file` - (Required) Save all instance information under the VDC to the path.