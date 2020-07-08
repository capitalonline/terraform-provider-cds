package cds

import (
	"context"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/security_group"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCdsSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCdsSecurityGroupRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "security group ID.",
			},
			"security_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "security group name.",
			},
			"security_group_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "security group type.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "used to save results.",
			},
		},
	}
}

func dataSourceCdsSecurityGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.securitygroup.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	sgService := SecurityGroupService{client: meta.(*CdsClient).apiConn}
	descRequest := security_group.NewDescribeSecurityGroupRequest()
	descRuleRequest := security_group.NewDescribeSecurityGroupRuleRequest()
	if v, ok := d.GetOk("id"); ok {
		descRequest.SecurityGroupId = common.StringPtr(v.(string))
		descRuleRequest.SecurityGroupId = common.StringPtr(v.(string))
	}
	if v, ok := d.GetOk("security_group_name"); ok {
		descRequest.Keyword = common.StringPtr(v.(string))
	}
	if v, ok := d.GetOk("security_group_type"); ok {
		descRequest.SecurityGroupType = common.StringPtr(v.(string))
	}

	result, err := sgService.DescribeSecurityGroup(ctx, descRequest)
	if err != nil {
		return err
	}
	ruleResult, err := sgService.DescribeSecurityGroupRule(ctx, descRuleRequest)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["SecurityGroup"] = result.Data
	data["Rules"] = ruleResult.Data

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), data); err != nil {
			return err
		}
	}

	return nil
}
