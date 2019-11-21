---
layout: "cds"
page_title: "Provider: CDS"
sidebar_current: "docs-cds-index"
description: |-
  The CDS provider is used to interact with many resources supported by cds. The provider needs to be configured with the proper credentials before it can be used.
---

# CDS Provider

The CDS provider is used to interact with the
many resources supported by [CDS](https://www.capitalonline.net/zh-cn/). The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
## Parameter preparation
variable "vdc_name" {
  default = "test_vdc"
}
variable "availability_zone" {
  default = "CN_Beijing_A"
}

variable "image_id" {
  default = "Ubuntu_16.04_64"
}

variable "instance_type" {
  default = "high_ccs"
}

variable "instance_name" {
  default = "test_instance"
}

variable "cpu" {
  default = 4
}

variable "ram" {
  default = 4
}

variable "password" {
  default = "123abc!@#"
}

variable "public_address" {
  default = "auto"
}

variable "group_name" {
  default = "test_sg"
}
variable "group_type" {
  default = "public"
}


## Create VDC
resource "cds_vdc" "my_vdc" {
  vdc_name = var.vdc_name
  public_network = {
    "IPNum" = 4
    "Qos" = 20
    "Name" = "test"
    "FloatBandwidth" = 200
    "BillingMethod" = "BandwIdth"
    "AutoRenew" = 1
    "Type" = "Bandwidth_BGP"
  }
}

## Create Private Subnet for VDC
resource "cds_private_subnet" "my_private_subnet_1" {
  vdc_id = cds_vdc.my_vdc.id
  name = "private_1"
  type = "private"
  addres = ""
  mask = "26"
}

## Create Security Group
resource "cds_security_group" "security_group_1" {
  name = var.group_name
  description = "New security group"
  type =var.group_type
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

## Create Instance
resource "cds_instance" "my_instance" {
  instance_name       = var.instance_name
  region_id           = "CN_Beijing_A"
  image_id            = var.image_id
  instance_type       = var.instance_type
  cpu                 = var.cpu
  ram                 = var.ram
  vdc_id              = cds_vdc.my_vdc.id
  password            = var.password
  public_ip           = var.public_address
  //if you want to stop or reboot instance, please open it
//    operate_instance_status = var.operate_instance_status
  private_ip          = {
    "private_id" = cds_private_subnet.my_private_subnet_1.id,
    "address" = "auto"
  }
}
```

## Authentication

The cds provider offers a flexible means of providing credentials for authentication.
The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

### Static credentials ###

Static credentials can be provided by adding an `secret_id` `secret_key` and `region` in-line in the
cds provider block:

Usage:

```hcl
provider "cds" {
  secret_id  = "${var.secret_id}"
  secret_key = "${var.secret_key}"
  region     = "${var.region}"
}
```


### Environment variables

You can provide your credentials via `CDS_SECRET_ID` and `CDS_SECRET_KEY`,
environment variables, representing your cds Access Key and Secret Key, respectively.
`CDS_REGION` is also used, if applicable:

```hcl
provider "cds" {}
```

Usage:

```shell
$ export CDS_SECRET_ID="your_fancy_accesskey"
$ export CDS_SECRET_KEY="your_fancy_secretkey"
$ export CDS_REGION="ap-guangzhou"
$ terraform plan
```


## Argument Reference

The following arguments are supported:

* `secret_id` - (Optional) This is the cds access key. It must be provided, but
  it can also be sourced from the `CDS_SECRET_ID` environment variable.

* `secret_key` - (Optional) This is the cds secret key. It must be provided, but
  it can also be sourced from the `CDS_SECRET_KEY` environment variable.

* `region` - (Required) This is the cds region. It must be provided, but
  it can also be sourced from the `CDS_REGION` environment variables.
  The default input value is ap-guangzhou.


## Testing

Credentials must be provided via the `CDS_SECRET_ID`, and `CDS_SECRET_KEY` environment variables in order to run acceptance tests.
