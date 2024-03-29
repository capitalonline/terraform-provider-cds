package cds

import (
	"context"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/vdc"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCdsPrivateSubnet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCdsPrivateSubnetRead,

		Schema: map[string]*schema.Schema{
			"vdc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "Vdc id.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Used to save results.",
			},
		},
		Description: "Data source private subnet.\n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`

` +
			"\n```",
	}
}

func dataSourceCdsPrivateSubnetRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.subnet.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	descRequest := vdc.DescribeVdcRequest()
	if v, ok := d.GetOk("vdc_id"); ok {
		descRequest.VdcId = common.StringPtr(v.(string))
	}

	result, err := vdcService.DescribeVdc(ctx, descRequest)
	if err != nil {
		return err
	}
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), result.Data[0].PrivateNetwork); err != nil {
			return err
		}
	}

	return nil
}
