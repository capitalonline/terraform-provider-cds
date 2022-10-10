---
layout: "cds"
page_title: "cds: instance"
sidebar_current: "docs-cds-resource-instance"
description: |-
  Provide resources to create or delete VDCs.
---

# Instance

Provide resources to create, update or delete instances.

## Example Usage

```hcl
# create instance
resource "cds_instance" "instance_example" {
  instance_name       = "instance_example"
  region_id           = var.region_id
  image_id            = var.image_id
  instance_type       = var.instance_type
  cpu                 = var.cpu
  ram                 = var.ram
  vdc_id              = cds_vdc.my_vdc.id
  password            = var.password
  public_ip           = var.public_address
  user_data = ["IyEvYmluL3NoCmVjaG8gIkhlbGxvIFdvcmxkIg==",
     "IyEvYmluL3NoCmVjaG8gIm5hbWVzZXJ2ZXIgOC44LjguOCIgfCB0ZWUgL2V0Yy9yZXNvbHYuY29uZg==",]
  operate_instance_status = var.operate_instance_status
  private_ip          = {
    "private_id" = cds_private_subnet.my_private_subnet_1.id,
    "address" = "auto"
  }
  system_disk = {
    type = "ssd_system_disk"
    size = 200
    iops = 20
  }

  #data_disks = [{   
  #    iops = 5
  #    size = 100
  #    type = "ssd_disk"
  # },
  # {
  #    iops = 0
  #    size = 120
  #    type = "high_disk"
  # }
  #]

  #update_data_disks = [
  #  {   disk_id = "xxxxxxxxxxxxxxxxxxxxxxxxx"
  #      iops = 22
  #      size = 122
  #      type = "ssd_disk"
  #  },
  #  {   disk_id = "xxxxxxxxxxxxxxxxxxxxxxxxx"
  #      iops = 33
  #      size = 133
  #      type = "ssd_disk"
  #  },
  #]

  #delete_data_disks = [
  #  {
  #    disk_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  #  }, 
  #  {
  #    disk_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  #  },    
  #]
  
  security_group_binding {
    type = "private"
    subnet_id = "private 1"
    security_group_id = "tf_SecurityGroup_xx"
  }
}
```

## Argument Reference

The following arguments are supported
> NOTE: The cds_instance resource supports batch creation by setting the amount parameter, but does not allow any modification in the batch creation mode. Please use it with caution.

* `instance_name` - (Required,Unmodifiable) The name of the instance.
* `region_id` - (Required,Unmodifiable) The Region of the instance, refer to [All Region](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#%E8%8A%82%E7%82%B9%E5%90%8D%E7%A7%B0).
* `image_id` - (Required,Unmodifiable)Template ID, if not specified, centos7.4 is selected by default (the interface displays the first one), refer to [All Image](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#%E5%85%AC%E5%85%B1%E6%A8%A1%E6%9D%BF). 
* `instance_type` - (Required,Unmodifiable) The type of the instance, refer to [All Instance Type](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#%E4%B8%BB%E6%9C%BA%E7%B1%BB%E5%9E%8B).
* `cpu` - (Required,Unmodifiable) The number of cpu, the unit (a) can only be selected [1, 2, 4, 8, 10, 16, 32] The default selection can be purchased the smallest.
* `ram` - (Required,Unmodifiable) The amount of memory, the unit (GB) can only be selected [1, 2, 4, 8, 12, 16, 24, 32, 48, 64, 96, 128] The default selection can be purchased the smallest.
* `vdc_id` - (Required,Unmodifiable) Instance belongs to the virtual data center.
* `password` - (Required,Unmodifiable) Password to an instance is a string of 8 to 30 characters. It must contain uppercase/lowercase letters, numerals and special symbols.
* `public_ip` - (Required,Unmodifiable) The public ip of the instance.
* `amount` - (Optional,Unmodifiable) Number of instances created in batch, maximum 50.
* `user_data` - (Optional) User defined data, the format must be base64 encoding
* `operate_instance_status` - (Optional) The status of the instance. Allow values: reboot, stop, run.
* `private_ip` - (Optional) Private ip.
  * `private_id` - (Required)Private subnet ID.
  * `address` - (Required)ip address. Automatically assign input: auto, the default is not written as not assigning private network ip.
* `system_disk` - (Optional) System Disk . If you not set system_disk , Default when created size = 20 type = system_disk ,iops = 0
  * `size` - The size of the disk in GiBs.
  * `type` - The type of the disk. Allow values: system_disk, ssd_system_disk.
  * `iops` - The size of the disk iops int, type equal ssd_system_disk can set iops, type equal system_disk can not set iops (iops must equal 0)
* `data_disks` - (Optional) Data Disks. Add at creation time and append after creation 
  * `size` - The size of the disk in GiBs.
  * `type` - The type of the disk. Allow values: big_disk, high_disk, ssd_disk.
  * `iops` - The size of the disk iops int. 
* `update_data_disks` - (Optional) modify Data Disks
  * `disk_id` - The id of data disk , from data source instance
  * `type` - The type of the disk. Allow values: big_disk, high_disk, ssd_disk.   
  * `size` - The size of the disk in GiBs.
  * `iops` - The size of the disk iops int,type equal ssd_disk can modify iops.
* `security_group_binding` - (Optional) Instance binding security group.
  * `type` -(Required) Specify a public or private network binding security group. Allow values: private, public.
  * `subnet_id` - (Required)Subnet ID.
  * `security_group_id` - (Required) Security group ID.
* `image_password` - The password of image