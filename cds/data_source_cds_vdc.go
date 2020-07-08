package cds

import (
	"context"
	"time"

	"terraform-provider-cds/cds-sdk-go/common"
	"terraform-provider-cds/cds-sdk-go/vdc"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceCdsVdc() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCdsVdcRead,

		Schema: map[string]*schema.Schema{
			"vdc_id": {
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
				Optional:    true,
				Description: "Used to save results.",
			},
			"vdc": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vdc_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vdc_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_network": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"private_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"unuse_ip_num": {
										// using string instead of int to prevent 0 for misleading
										Type:     schema.TypeString,
										Computed: true,
									},
									"segments": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"public_network": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"public_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"qos": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"unuse_ip_num": {
										// using string instead of int to prevent 0 for misleading
										Type:     schema.TypeString,
										Computed: true,
									},
									"segments": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mask": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"gateway": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"segment_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"address": {
													Type:     schema.TypeString,
													Computed: true,
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
	if v, ok := d.GetOk("vdc_id"); ok {
		descRequest.VdcId = common.StringPtr(v.(string))
	}
	if v, ok := d.GetOk("vdc_name"); ok {
		descRequest.Keyword = common.StringPtr(v.(string))
	}

	result, err := vdcService.DescribeVdc(ctx, descRequest)
	if err != nil {
		return err
	}

	return vdcDescriptionAttributes(d, result)
}

func vdcDescriptionAttributes(d *schema.ResourceData, result vdc.DescVdcResponse) error {
	var names []string
	var out []map[string]interface{}

	for _, vdc := range result.Data {
		mapping := map[string]interface{}{
			"vdc_id":          vdc.VdcId,
			"vdc_name":        vdc.VdcName,
			"region_id":       vdc.RegionId,
			"private_network": flattenPrivateNetworks(vdc.PrivateNetwork),
			"public_network":  flattenPublicNetworks(vdc.PublicNetwork),
		}
		names = append(names, *vdc.VdcName)
		out = append(out, mapping)
	}

	d.SetId(time.Now().UTC().String())
	// TODO: sort by region id
	if err := d.Set("vdc", out); err != nil {
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), result.Data); err != nil {
			return err
		}
	}

	return nil
}

func flattenPublicSegments(publicSegments []*vdc.PublicSegment) []map[string]interface{} {

	publicSegmentSchema := make([]map[string]interface{}, 0, len(publicSegments))

	for _, publicSegment := range publicSegments {
		data := map[string]interface{}{
			"mask":       publicSegment.Mask,
			"gateway":    publicSegment.Gateway,
			"segment_id": publicSegment.SegmentId,
			"address":    publicSegment.Address,
		}

		publicSegmentSchema = append(publicSegmentSchema, data)
	}

	return publicSegmentSchema
}

func flattenPublicNetworks(publicNetworks []*vdc.PublicNetworkInfo) []map[string]interface{} {
	publicNetworkSchema := make([]map[string]interface{}, 0, len(publicNetworks))

	for _, publicNetwork := range publicNetworks {
		data := map[string]interface{}{
			"public_id":    publicNetwork.PublicId,
			"status":       publicNetwork.Status,
			"qos":          publicNetwork.Qos,
			"name":         publicNetwork.Name,
			"unuse_ip_num": intToString(publicNetwork.UnuseIpNum),
			"segments":     flattenPublicSegments(publicNetwork.Segments),
		}

		publicNetworkSchema = append(publicNetworkSchema, data)
	}

	return publicNetworkSchema
}

func flattenPrivateNetworks(privateNetworks []*vdc.PrivateNetwork) []map[string]interface{} {
	privateNetworkSchema := make([]map[string]interface{}, 0, len(privateNetworks))

	for _, privateNetwork := range privateNetworks {
		data := map[string]interface{}{
			"private_id":   privateNetwork.PrivateId,
			"status":       privateNetwork.Status,
			"name":         privateNetwork.Name,
			"unuse_ip_num": intToString(privateNetwork.UnuseIpNum),
			"segments":     privateNetwork.Segments,
		}

		privateNetworkSchema = append(privateNetworkSchema, data)
	}

	return privateNetworkSchema
}
