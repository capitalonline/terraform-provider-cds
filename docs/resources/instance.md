---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cds_instance Resource - terraform-provider-cds"
subcategory: ""
description: |-
  Instance of vm. View documentation https://github.com/capitalonline/openapi/blob/master/%E4%BA%91%E4%B8%BB%E6%9C%BA%E6%A6%82%E8%A7%88.md

---

# cds_instance (Resource)

Instance of vm. [View documentation](https://github.com/capitalonline/openapi/blob/master/%E4%BA%91%E4%B8%BB%E6%9C%BA%E6%A6%82%E8%A7%88.md)

## Example Usage

```hcl

resource "cds_instance" "my_instance" {
  instance_name = "test_zz_04"
  region_id     = "CN_Beijing_A"
  image_id      = "Ubuntu_16.04_64"
  instance_type = "high_ccs"
  # In v1.4.5 and later, changing the instance specification will automatically shut down and start up
  cpu           = 4
  ram           = 4
  vdc_id        = cds_vdc.my_vdc.id
  # public_key = file("/home/guest/.ssh/test.pub")
  # password is required after v1.3.1
  password  = "123abc,.;"
  # image_password is optional
  image_password = "123abc,.;"
  # operate_instance_status required value 'run' or 'stop' or 'reboot'
  operate_instance_status = "run"
  # user self-defined data,must be encoded by base64
  user_data = ["IyEvYmluL3NoCmVjaG8gIkhlbGxvIFdvcmxkIg==",#!/bin/sh echo "Hello World"
    "IyEvYmluL3NoCmVjaG8gIm5hbWVzZXJ2ZXIgOC44LjguOCIgfCB0ZWUgL2V0Yy9yZXNvbHYuY29uZg==",]#!/bin/sh echo "nameserver 8.8.8.8" | tee /etc/resolv.conf
  public_ip = "auto"
  private_ip {
    private_id = cds_private_subnet.my_private_subnet_1.id
    address    = "auto"
  }


  # type  system_disk | ssd_system_disk
  # if type = system_disk ,you can not set and modify iops ,iops must set 0
  # if type = ssd_system_disk ,you can set and modify iops
  # if you not set system_disk , Default when created size = 20 type = system_disk ,iops = 0
  system_disk = {
    type = var.system_disk_type
    size = 100
    iops = 5
  }
  
  # you can set data_disks at create instance ,or after append data_disks
  
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

  # you can modify data disk iops and size 
  # if you want modify iops ,the type must be ssd_disk,
  # if type = high_disk ,the iops must equal 0
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

  # you can delete data disks by disk_id 
  # disk_id from data cds_data_source_instance
  #delete_data_disks = [
  #  {
  #    disk_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  #  }, 
  #  {
  #    disk_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  #  },    
  #]
  
  

  security_group_binding {
    type              = "private"
    subnet_id         = cds_private_subnet.my_private_subnet_1.id
    security_group_id = cds_security_group.security_group_2.id
  }
  #utc = true
}

```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cpu` (Number) The number of cpu, the unit (a) can only be selected [1, 2, 4, 8, 10, 16, 32] The default selection can be purchased the smallest.
- `instance_name` (String) The name of the instance
- `instance_type` (String) The type of the instance.[All valid type](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#%E4%B8%BB%E6%9C%BA%E7%B1%BB%E5%9E%8B)
- `ram` (Number) The amount of memory, the unit (GB) can only be selected [1, 2, 4, 8, 12, 16, 24, 32, 48, 64, 96, 128] The default selection can be purchased the smallest.
- `vdc_id` (String) Vdc Id

### Optional

- `amount` (Number) Number of instances created in batch, maximum 50
- `auto_renew` (Number) Whether the subscription-based instance will automatically renew. 1 indicates automatic renewal (default), 0 indicates no automatic renewal.
- `data_disks` (Block List, Max: 15) Data Disks. Add at creation time and append after creation. The quantity limit can be set between 1 and 15. (see [below for nested schema](#nestedblock--data_disks))
- `delete_data_disks` (Block List, Max: 15) Data disks to delete. (see [below for nested schema](#nestedblock--delete_data_disks))
- `host_name` (String) Host name.
- `image_id` (String) The image of the operating system, default value is Ubuntu_16.04_64. [All valid images ](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#%E5%85%AC%E5%85%B1%E6%A8%A1%E6%9D%BF)
- `image_password` (String) When using a public image, this field is optional; when using a custom image, this field is required.
- `instance_charge_type` (String) The payment methods for instance are as follows: 1. PrePaid: Prepaid, monthly or yearly subscription.  2. PostPaid (default): Pay-as-you-go.
- `operate_instance_status` (String) Operate instance . Allow values: `reboot`, `stop`, `run`
- `password` (String, Sensitive) Password to an instance is a string of 8 to 30 characters. It must contain uppercase/lowercase letters, numerals and special symbols.
- `prepaid_month` (Number) The duration of the subscription for the instance, where 0 indicates a purchase until the end of the current month, 1 indicates one full calendar month. The default value is 0.
- `private_ip` (Block List, Max: 15) Private ip (see [below for nested schema](#nestedblock--private_ip))
- `public_ip` (String) The public ip of the instance
- `public_key` (String) Public key
- `region_id` (String) The Region of the instance, refer to [All Regions](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#%E5%8F%AF%E7%94%A8%E5%8C%BA%E5%90%8D%E7%A7%B0)
- `security_group_binding` (Block Set, Max: 15) Instance binding security group (see [below for nested schema](#nestedblock--security_group_binding))
- `subject_id` (Number) Subject id.
- `dedicated_host_id` (String) Dedicated host id.
- `system_disk` (Map of String) System Disk . If not set , default: size = 20 , type = system_disk ,iops = 0
- `update_data_disks` (Block List, Max: 15) Modify data disks (see [below for nested schema](#nestedblock--update_data_disks))
- `user_data` (List of String) User-defined data must be in base64 encoding format.
- `utc` (Boolean) Whether to set the time zone to UTC

### Read-Only

- `id` (String) The ID of this resource.
- `instance_status` (String) Status of the instance

<a id="nestedblock--data_disks"></a>
### Nested Schema for `data_disks`

Required:

- `size` (Number)

Optional:

- `iops` (Number) The size of the disk iops
- `type` (String) The type of the disk. Allow values: `big_disk`, `high_disk`, `ssd_disk`


<a id="nestedblock--delete_data_disks"></a>
### Nested Schema for `delete_data_disks`

Optional:

- `disk_id` (String) Disk id


<a id="nestedblock--private_ip"></a>
### Nested Schema for `private_ip`

Required:

- `private_id` (String) Private subnet ID.

Optional:

- `address` (String) Ip address. Automatically assign input: auto, the default is not written as not assigning private network ip.
- `interface_id` (String) Network interface id


<a id="nestedblock--security_group_binding"></a>
### Nested Schema for `security_group_binding`

Required:

- `security_group_id` (String) Security group ID
- `subnet_id` (String) Subnet ID
- `type` (String) Specify a public or private network binding security group. Allow values: `private`, `public`.


<a id="nestedblock--update_data_disks"></a>
### Nested Schema for `update_data_disks`

Required:

- `size` (Number) The size of the disk in GiBs.

Optional:

- `disk_id` (String) The id of data disk , from data source instance
- `iops` (Number) The size of the disk iops int,type equal ssd_disk can modify iops.
- `type` (String) The type of the disk. Allow values: `big_disk``, `high_disk`, `ssd_disk`
