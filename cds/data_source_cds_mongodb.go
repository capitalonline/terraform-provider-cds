package cds

import (
	"context"
	"fmt"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/mongodb"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCdsMongodb() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCdsMongodnRead,

		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Region id.",
			},
			"instance_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance uuid.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance name.",
			},
			"ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Ip.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Used to save results.",
			},
		},
		Description: "Data source mongodb.\n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
data cds_data_source_mongodb "mongodb_data" {
    region_id           = "CN_Beijing_A"
    instance_uuid       = "xxx"
    instance_name       = "xxx"
    ip                  = ""
    result_output_file  = "data.json" // availableDB, instances, regions
}
` +
			"\n```",
	}
}

func dataSourceCdsMongodnRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.mongodb.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	result := map[string]interface{}{}
	mongodbService := MongodbService{client: meta.(*CdsClient).apiConn}

	//DescribeZones
	zonesRequest := mongodb.NewDescribeZonesRequest()

	zonesResponse, err := mongodbService.DescribeZones(ctx, zonesRequest)

	if err != nil {
		return err
	}

	if *zonesResponse.Code != "Success" {
		return fmt.Errorf("describe zones response errors :%s", *zonesResponse.Message)
	}

	result["zones"] = zonesResponse.Data

	//DescribeSpecInfo
	describeSpecInfoRequest := mongodb.NewDescribeSpecInfoRequest()
	describeSpecInfoRequest.RegionId = common.StringPtr(d.Get("region_id").(string))
	describeSpecInfoResponse, err := mongodbService.DescribeSpecInfo(ctx, describeSpecInfoRequest)
	if err != nil {
		return err
	}
	if *describeSpecInfoResponse.Code != "Success" {
		return fmt.Errorf("describe Spec Info is error:%s", *describeSpecInfoResponse.Message)
	}

	result["availableMongodb"] = describeSpecInfoResponse.Data
	//DescribeDBInstances

	instancesRequest := mongodb.NewDescribeDBInstancesRequest()

	if inter, ok := d.GetOk("instance_uuid"); ok {
		instancesRequest.InstanceUuid = common.StringPtr(inter.(string))
	}
	if inter, ok := d.GetOk("instance_name"); ok {
		instancesRequest.InstanceName = common.StringPtr(inter.(string))
	}
	if inter, ok := d.GetOk("ip"); ok {
		instancesRequest.IP = common.StringPtr(inter.(string))
	}

	instancesResponse, err := mongodbService.DescribeDBInstances(ctx, instancesRequest)

	if err != nil {
		return err
	}

	if *instancesResponse.Code != "Success" {
		return fmt.Errorf("get mongodb instance list failed, error: %s", *instancesResponse.Message)
	}

	result["instances"] = instancesResponse.Data
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), result); err != nil {
			return err
		}
	}

	return nil
}
