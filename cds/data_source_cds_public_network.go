package cds

import (
	"context"
	"errors"
	"fmt"
	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/vdc"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCdsPublicNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCdsPublicNetworkRead,

		Schema: map[string]*schema.Schema{
			"vdc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vdc id.",
			},
			"public_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Public network id.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status.",
			},
			"qos": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Qos.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name.",
			},
			"unuse_ip_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Un use ip num.",
			},
			"segments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Public network segments.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mask": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Mask.",
						},
						"gateway": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway.",
						},
						"segment_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Segment id.",
						},
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Address",
						},
					},
				},
			},
		},
		Description: "Data source public network.\n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
data "cds_data_source_public_network" "pbn" {
  vdc_id    = ""
  public_id = ""
}
` +
			"\n```",
	}
}

func dataSourceCdsPublicNetworkRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.public_network.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	descRequest := vdc.DescribeVdcRequest()
	vdcId := d.Get("vdc_id")
	descRequest.VdcId = common.StringPtr(vdcId.(string))
	resp, err := vdcService.DescribeVdc(ctx, descRequest)
	if err != nil {
		return err
	}
	if *resp.Code != "Success" || len(resp.Data) == 0 {
		return errors.New(fmt.Sprintf("request public network failed: %v,message: %s", *resp.Code, *resp.Message))
	}
	publicId := d.Get("public_id")
	for _, publicNetwork := range resp.Data[0].PublicNetwork {
		if *publicNetwork.PublicId == publicId.(string) {
			d.SetId(publicId.(string))
			if publicNetwork.Status != nil {
				d.Set("status", *publicNetwork.Status)
			}
			if publicNetwork.Qos != nil {
				d.Set("qos", *publicNetwork.Qos)
			}
			if publicNetwork.Name != nil {
				d.Set("name", *publicNetwork.Name)
			}
			if publicNetwork.UnuseIpNum != nil {
				d.Set("unuse_ip_num", *publicNetwork.UnuseIpNum)
			}
			var segments = make([]map[string]interface{}, 0)
			if publicNetwork.Segments != nil {
				for _, item := range *publicNetwork.Segments {
					segment := make(map[string]interface{})
					if item.Mask != nil {
						segment["mask"] = *item.Mask
					}
					if item.Gateway != nil {
						segment["gateway"] = *item.Gateway
					}
					if item.SegmentId != nil {
						segment["segment_id"] = *item.SegmentId
					}
					if item.Address != nil {
						segment["address"] = *item.Address
					}
					segments = append(segments, segment)
				}
			}
			d.Set("segments", segments)
			return nil
		}
	}

	return errors.New("the public network is not available")
}
