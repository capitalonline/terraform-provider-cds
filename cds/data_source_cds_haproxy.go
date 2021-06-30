package cds

import (
	"context"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/haproxy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHaproxy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyRead,

		Schema: map[string]*schema.Schema{
			"ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "haproxy ip.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "instance name.",
			},
			"start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "start time.",
			},
			"end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "end time.",
			},
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Region ID.",
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceHaproxyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.haproxy.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	request := haproxy.NewDescribeLoadBalancersRequest()
	if inter, ok := d.GetOk("ip"); ok {
		ip, exist := inter.(string)
		if exist {
			request.IP = common.StringPtr(ip)
		}
	}
	if inter, ok := d.GetOk("instance_name"); ok {
		instanceName, exist := inter.(string)
		if exist {
			request.InstanceName = common.StringPtr(instanceName)
		}
	}
	if inter, ok := d.GetOk("start_time"); ok {
		startTime, exist := inter.(string)
		if exist {
			request.StartTime = common.StringPtr(startTime)
		}
	}
	if inter, ok := d.GetOk("end_time"); ok {
		endTime, exist := inter.(string)
		if exist {
			request.EndTime = common.StringPtr(endTime)
		}
	}

	response, err := haproxyService.DescribeHaproxy(ctx, request)
	if err != nil {
		return err
	}

	datas := []map[string]interface{}{}
	for _, ha := range response.Data {
		strategyRequest := haproxy.NewDescribeLoadBalancerStrategysRequest()
		strategyRequest.InstanceUuid = ha.InstanceUuid
		strategtResponse, err := haproxyService.DescribeLoadBalancerStrategys(ctx, strategyRequest)
		if err != nil {
			return nil
		}

		datas = append(datas, map[string]interface{}{
			"instance": ha,
			"strategy": strategtResponse.Data,
		})
	}

	goodsRequest := haproxy.NewDescribeLoadBalancersSpecRequest()

	if inter, ok := d.GetOk("region_id"); ok {
		regionId, exist := inter.(string)
		if exist {
			goodsRequest.RegionId = common.StringPtr(regionId)
		}
	}

	goodsResponse, err := haproxyService.DescribeLoadBalancersSpec(ctx, goodsRequest)
	if err != nil {
		return err
	}

	zonesResponse, err := haproxyService.DescribeZones(ctx, haproxy.NewDescribeZonesRequest())
	if err != nil {
		return err
	}

	result := map[string]interface{}{
		"haproxy": datas,
		"goods":   goodsResponse.Data,
		"zones":   zonesResponse.Data,
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), result); err != nil {
			return err
		}
	}

	return nil
}
