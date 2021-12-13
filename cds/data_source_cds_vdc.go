package cds

import (
	"context"
	"log"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/vdc"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
			"vdc_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of vdc configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vdc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "vdc id",
						},
						"vdc_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "vdc name",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "region id",
						},
						"private_network": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "PrivateNetwork list",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"private_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Private network Id",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "status",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "name",
									},
									"unuse_ip_num": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "unuse ip num",
									},
									"segments": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "private network segments list",
									},
								},
							},
						},
						"public_network": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Public Network list",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"public_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "public id",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "status",
									},
									"qos": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "qos",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "name",
									},
									"unuse_ip_num": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "un use ip num",
									},
									"segments": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "public network segments",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mask": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "mask",
												},
												"gateway": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "gateway",
												},
												"segment_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "segmetId",
												},
												"address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "address",
												},
											},
										},
									},
								},
							},
						},
					},
				},
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
	vdc_id := ""
	if v, ok := d.GetOk("id"); ok {
		vdc_id = v.(string)
		descRequest.VdcId = common.StringPtr(v.(string))
	}
	if v, ok := d.GetOk("vdc_name"); ok {
		descRequest.Keyword = common.StringPtr(v.(string))
	}

	result, err := vdcService.DescribeVdc(ctx, descRequest)
	if err != nil {
		return err
	}
	vdcList := make([]map[string]interface{}, 0, len(result.Data))
	for _, vdcInfo := range result.Data {
		mapping := map[string]interface{}{
			"vdc_id":          *vdcInfo.VdcId,
			"vdc_name":        *vdcInfo.VdcName,
			"region_id":       *vdcInfo.RegionId,
			"private_network": flattenPrivateNetworkMappings(vdcInfo.PrivateNetwork),
			"public_network":  flattenPublicNetworkMappings(vdcInfo.PublicNetwork),
		}
		vdcList = append(vdcList, mapping)
	}
	d.SetId(vdc_id)
	err = d.Set("vdc_list", vdcList)
	log.Printf("vdc_list:%v", vdcList)
	if err != nil {
		log.Printf("[CRITAL]%s provider set configuration list fail, reason:%s\n ", logId, err.Error())
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), result.Data); err != nil {
			return err
		}
	}

	return nil
}
