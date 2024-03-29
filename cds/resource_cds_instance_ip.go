package cds

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceCdsInstanceIp() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCdsHaproxy,
		Read:   readResourceCdsHaproxy,
		Update: updateResourceCdsHaproxy,
		Delete: deleteRresourceCdsHaproxy,
		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "regon id.",
			},
			"vdc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "vdc id.",
			},
			"base_pipe_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "base pipe id.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance name.",
			},
			"cpu": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "instance cpu num",
			},
			"ram": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "instance ram size",
			},
			"ips": {
				Type:       schema.TypeList,
				ConfigMode: schema.SchemaConfigModeAttr,
				Required:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pipe_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"pipe_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"segment_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
		Description: "Instance of vm. [View documentation](https://github.com/capitalonline/openapi/blob/master/%E4%BA%91%E4%B8%BB%E6%9C%BA%E6%A6%82%E8%A7%88.md)\n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
// create instance
resource "cds_instance" "my_instance2" {
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
` +
			"\n```",
	}
}
