package cds

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceCdsMySQLAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCdsMySQLAccountRead,

		Schema: map[string]*schema.Schema{
			"instance_uuid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance uuid.",
			},
			"account_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Account name.",
			},
			"account_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Account type.",
			},
			"account_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Account description.",
			},
			"service_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Service id.",
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
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Account privilege type.",
						},
						"db_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Database name.",
						},
						"account_privilege_detail": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Account privilege detail",
						},
					},
				},
				Description: "Database privileges.",
			},
		},
	}
}

func dataSourceCdsMySQLAccountRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}
