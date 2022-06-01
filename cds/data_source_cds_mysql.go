package cds

import (
	"context"
	"fmt"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/mysql"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCdsMySQL() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCdsMySQLRead,

		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Region ID.",
			},
			"instance_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance uuid",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance name",
			},
			"ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ip",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Used to save results.",
			},
			"readonly_instances": {
				Type:     schema.TypeList,
				Optional: true,
				//Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Description: "create readonly instances ",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"paas_goods_id": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"test_group_id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"disk_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"disk_value": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"amount": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCdsMySQLRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.mysql.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	result := map[string]interface{}{}

	mySQLService := MySQLService{client: meta.(*CdsClient).apiConn}

	regionsRequest := mysql.NewDescribeRegionsRequest()
	regionsResponse, err := mySQLService.DescribeRegions(ctx, regionsRequest)
	if err != nil {
		return err
	}
	if *regionsResponse.Code != "Success" {
		return fmt.Errorf("describe regions response error: %s", *regionsResponse.Message)
	}

	result["regions"] = regionsResponse.Data

	availableDBRequest := mysql.NewDescribeAvailableDBConfigRequest()
	availableDBRequest.RegionId = common.StringPtr(d.Get("region_id").(string))
	availableDBResponse, err := mySQLService.DescribeAvailableDBConfig(ctx, availableDBRequest)
	if err != nil {
		return err
	}

	if *availableDBResponse.Code != "Success" {
		return fmt.Errorf("describe available db config error: %s", *availableDBResponse.Message)
	}

	result["availableDB"] = availableDBResponse.Data

	instancesRequest := mysql.NewDescribeDBInstancesRequest()
	if inter, ok := d.GetOk("instance_uuid"); ok {
		instancesRequest.InstanceUuid = common.StringPtr(inter.(string))
	}
	if inter, ok := d.GetOk("instance_name"); ok {
		instancesRequest.InstanceName = common.StringPtr(inter.(string))
	}
	if inter, ok := d.GetOk("ip"); ok {
		instancesRequest.IP = common.StringPtr(inter.(string))
	}

	instancesResponse, err := mySQLService.GetMySQLList(ctx, instancesRequest)
	if err != nil {
		return err
	}

	if *instancesResponse.Code != "Success" {
		return fmt.Errorf("get mysql instance list failed, error: %s", *instancesResponse.Message)
	}

	result["instances"] = instancesResponse.Data

	if inter, ok := d.GetOk("instance_uuid"); ok {
		avarilableModifyInstanceRequest := mysql.NewDescribeModifiableDBSpecRequest()
		avarilableModifyInstanceRequest.InstanceUuid = common.StringPtr(inter.(string))

		avarilableModifyInstanceResponse, err := mySQLService.DescribeModifiableDBSpec(ctx, avarilableModifyInstanceRequest)
		if err != nil {
			return err
		}

		if *avarilableModifyInstanceResponse.Code != "Success" {
			return fmt.Errorf("read avariable modify instance spec error: %s", *avarilableModifyInstanceResponse.Message)
		}

		result["modify_instance_spec"] = avarilableModifyInstanceResponse.Data
	}

	req := mysql.NewDescribeAvailableReadOnlyConfigRequest()
	instanceUuid, _ := d.GetOk("instance_uuid")
	req.InstanceUuid = common.StringPtr(instanceUuid.(string))
	resp, err := mySQLService.GetAvailableReadOnlyConfig(ctx, req)
	if err == nil && *resp.Code == "Success" {
		result["available_read_only_config"] = resp.Data
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), result); err != nil {
			return err
		}
	}

	return nil
}
