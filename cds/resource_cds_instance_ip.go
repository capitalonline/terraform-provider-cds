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
	}
}
