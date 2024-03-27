package cds

import (
	"context"
	"fmt"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/redis"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCdsRedis() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCdsRedisRead,

		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Region ID.",
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
		Description: "Data source dedicated host.\n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
data cds_data_source_redis "redis_data" {
    region_id           = ""
    instance_uuid       = ""
    instance_name       = ""
    ip                  = ""
    result_output_file  = "data.json" // availableDB, instances, regions
}
` +
			"\n```",
	}
}

func dataSourceCdsRedisRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.redis.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	result := map[string]interface{}{}
	redisService := RedisService{client: meta.(*CdsClient).apiConn}

	//DescribeRegion
	regionsRequest := redis.NewDescribeRegionsRequest()
	regionsRequest.SetHttpMethod("GET")
	regionsResponse, err := redisService.DescribeRegions(ctx, regionsRequest)

	if err != nil {
		return err
	}

	if *regionsResponse.Code != "Success" {
		return fmt.Errorf("describe region response errors :%s", *regionsResponse.Message)
	}

	result["regions"] = regionsResponse.Data

	//DescribeAvailableDBConfig
	availableDBRequest := redis.NewDescribeAvailableDBConfigRequest()
	availableDBRequest.RegionId = common.StringPtr(d.Get("region_id").(string))
	availableDBResponse, err := redisService.DescribeAvailableDBConfig(ctx, availableDBRequest)
	if err != nil {
		return err
	}

	if *availableDBResponse.Code != "Success" {
		return fmt.Errorf("describe available db config error: %s", *availableDBResponse.Message)
	}
	result["availableDB"] = availableDBResponse.Data

	//DescribeRedisDescribeDBInstance
	instancesRequest := redis.NewDescribeDBInstancesRequest()
	instancesRequest.SetHttpMethod("GET")
	if inter, ok := d.GetOk("instance_uuid"); ok {
		instancesRequest.InstanceUuid = common.StringPtr(inter.(string))
	}
	if inter, ok := d.GetOk("instance_name"); ok {
		instancesRequest.InstanceName = common.StringPtr(inter.(string))
	}
	if inter, ok := d.GetOk("ip"); ok {
		instancesRequest.IP = common.StringPtr(inter.(string))
	}

	instancesResponse, err := redisService.DescribeRedis(ctx, instancesRequest)

	if err != nil {
		return err
	}

	if *instancesResponse.Code != "Success" {
		return fmt.Errorf("get redis instance list failed, error: %s", *instancesResponse.Message)
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
