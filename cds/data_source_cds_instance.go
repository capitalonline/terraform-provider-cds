package cds

import (
	"context"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/instance"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCdsInstance() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCdsInstanceRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "instance ID.",
			},
			"vdc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "vdc id.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceCdsInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.vdc.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	instanceService := InstanceService{client: meta.(*CdsClient).apiConn}
	descRequest := instance.NewDescribeInstanceRequest()
	if v, ok := d.GetOk("instance_id"); ok {
		descRequest.InstanceId = common.StringPtr(v.(string))
	}
	if v, ok := d.GetOk("vdc_id"); ok {
		descRequest.VdcId = common.StringPtr(v.(string))
	}

	result, err := instanceService.DescribeInstance(ctx, descRequest)
	if err != nil {
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), result.Data); err != nil {
			return err
		}
	}

	return nil
}
