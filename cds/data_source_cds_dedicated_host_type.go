package cds

import (
	"context"
	"errors"
	"fmt"
	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/instance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceDedicatedHostType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedHostTypeRead,
		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Region id.",
			},
			"dedicated_host_types": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Dedicated host types.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bill_scheme_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Bill scheme id.",
						},
						"cpu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Cpu.",
						},
						"goods_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Goods id.",
						},
						"ram": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Ram.",
						},
						"vm_family_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vm family id.",
						},
						"vm_rule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vm rule name.",
						},
						"vm_spec_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vm spec id.",
						},
						"vm_type_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vm type description.",
						},
						"vm_type_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vm type id.",
						},
						"vm_type_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vm type name.",
						},
						"vm_type_sort": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Vm type sort.",
						},
					},
				},
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Result output file. Used to save results.",
			},
		},
		Description: "Data source dedicated host.\n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
data "cds_data_source_dedicated_host_types" "host_type" {
	region_id 		= "CN_Beijing_H"
}
` +
			"\n```",
	}
}

func dataSourceDedicatedHostTypeRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.dedicated_host_type.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	instanceService := InstanceService{client: meta.(*CdsClient).apiConn}
	request := instance.NewDescribeDedicatedHostTypesRequest()
	regionId, ok := d.Get("region_id").(string)
	if !ok {
		return errors.New("region_id is required")
	}
	request.RegionId = common.StringPtr(regionId)
	response, err := instanceService.DescribeDedicatedHostTypes(ctx, request)
	if err != nil {
		return fmt.Errorf("describe dedicated hosts err:%v", err)
	}
	if *response.Code != success {
		return fmt.Errorf("describe dedicated hosts failed. with code:%s, with message:%s", *response.Code, *response.Message)
	}
	dedicatedHostTypes := make([]map[string]interface{}, 0, len(response.Data))
	for _, item := range response.Data {
		var element = make(map[string]interface{})
		if item.BillSchemeId != nil {
			element["bill_scheme_id"] = *item.BillSchemeId
		}
		if item.Cpu != nil {
			element["cpu"] = *item.Cpu
		}
		if item.Ram != nil {
			element["ram"] = *item.Ram
		}
		if item.GoodsId != nil {
			element["goods_id"] = *item.GoodsId
		}
		if item.VmFamilyId != nil {
			element["vm_family_id"] = *item.VmFamilyId
		}
		if item.VmRuleName != nil {
			element["vm_rule_name"] = *item.VmRuleName
		}
		if item.VmSpecId != nil {
			element["vm_spec_id"] = *item.VmSpecId
		}
		if item.VmTypeDescription != nil {
			element["vm_type_description"] = *item.VmTypeDescription
		}
		if item.VmTypeId != nil {
			element["vm_type_id"] = *item.VmTypeId
		}
		if item.VmTypeName != nil {
			element["vm_type_name"] = *item.VmTypeName
		}
		if item.VmTypeSort != nil {
			element["vm_type_sort"] = *item.VmTypeSort
		}
		dedicatedHostTypes = append(dedicatedHostTypes, element)
	}
	if err = d.Set("dedicated_host_types", dedicatedHostTypes); err != nil {
		return err
	}
	id := fmt.Sprintf("dedicated_host_types-%s", regionId)
	d.SetId(id)
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), map[string]interface{}{
			"dedicated_host_types": dedicatedHostTypes,
		}); err != nil {
			return err
		}
	}
	return nil
}
