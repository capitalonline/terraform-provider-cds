package cds

import (
	"context"
	"github.com/hashicorp/terraform/helper/schema"
	"terraform-provider-cds/cds-sdk-go/common"
	"terraform-provider-cds/cds-sdk-go/vdc"
)

func dataSourceCdsVdc() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCdsVdcRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "vdc ID.",
			},
			"vdc_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "vdc name.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceCdsVdcRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.vdc.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	descRequest := vdc.DescribeVdcRequest()
	if v, ok := d.GetOk("id"); ok {
		descRequest.VdcId = common.StringPtr(v.(string))
	}
	if v, ok := d.GetOk("vdc_name"); ok {
		descRequest.Keyword = common.StringPtr(v.(string))
	}

	result, err := vdcService.DescribeVdc(ctx, descRequest)
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
