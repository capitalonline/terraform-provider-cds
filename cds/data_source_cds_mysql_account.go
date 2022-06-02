package cds

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceCdsMySQLAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCdsMySQLAccountRead,

		Schema: map[string]*schema.Schema{
			"instance_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database_privileges": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_privilege_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"db_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"account_privilege_detail": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCdsMySQLAccountRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}
