package cds

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/redis"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCdsRedis() *schema.Resource {
	return &schema.Resource{
		Create: resourceCdsRedisCreate,
		Read:   resourceCdsRedisRead,
		Update: resourceCdsRedisUpdate,
		Delete: resourceCdsRedisDelete,
		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Region id. ",
			},
			"vdc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vdc id. ",
			},
			"base_pipe_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Base pipe id. ",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance name. ",
			},
			"architecture_type": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Architecture type. [View Document](https://github.com/capitalonline/openapi/blob/master/Redis%E6%A6%82%E8%A7%88.md#2describeavailabledbconfig)",
			},
			"ram": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Ram.[View Document](https://github.com/capitalonline/openapi/blob/master/Redis%E6%A6%82%E8%A7%88.md#2describeavailabledbconfig)",
			},
			"redis_version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Redis version.[View Document](https://github.com/capitalonline/openapi/blob/master/Redis%E6%A6%82%E8%A7%88.md#2describeavailabledbconfig)",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password.",
			},
			"ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Ip.",
			},
		},
		Description: "Redis instance. \n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
resource "cds_redis" "redis_example" {
    region_id         = "CN_Beijing_A"
    vdc_id            = "xxx"
    base_pipe_id      = "xxx"
    instance_name     = "redis_test"
    architecture_type = 3
    ram               = 4
    redis_version     = "2.8"
    password          = "password"
}
` +
			"\n```",
	}
}

func resourceCdsRedisCreate(data *schema.ResourceData, meta interface{}) error {
	log.Println("create redis instance")
	defer logElapsed("resource.cds_redis.create")()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), "logId", logId)

	redisService := RedisService{client: meta.(*CdsClient).apiConn}

	request := redis.NewCreateDBInstanceRequest()
	request.RegionId = common.StringPtr(data.Get("region_id").(string))
	request.VdcId = common.StringPtr(data.Get("vdc_id").(string))
	request.BasePipeId = common.StringPtr(data.Get("base_pipe_id").(string))
	request.Password = common.StringPtr(data.Get("password").(string))
	request.InstanceName = common.StringPtr(data.Get("instance_name").(string))

	architecture_type := common.IntPtr(data.Get("architecture_type").(int))
	ram := common.IntPtr(data.Get("ram").(int))
	version := common.StringPtr(data.Get("redis_version").(string))

	passGoodsId, err := matchRedisPassGoodsId(ctx, redisService, *architecture_type, *ram, *version, *request.RegionId)

	if err != nil {
		return err
	}

	request.PaasGoodsId = &passGoodsId

	amount := 1
	request.Amount = common.IntPtr(amount)
	response, err := redisService.CreateRedis(ctx, request)

	if err != nil {
		return err
	}
	if *response.Code != "Success" {
		return fmt.Errorf("create redis db instance failed,error: %v", err)
	}

	if len(response.Data.InstancesUuid) == 0 {
		return fmt.Errorf("create db failed")
	}
	data.SetId(response.Data.InstancesUuid[0])
	time.Sleep(100 * time.Second)
	if err := waitRedisRunning(ctx, redisService, response.Data.InstancesUuid[0]); err != nil {
		return err
	}

	return resourceCdsRedisRead(data, meta)
}

func resourceCdsRedisRead(data *schema.ResourceData, meta interface{}) error {
	log.Println("read redis instance")
	defer logElapsed("resource.cds_redis.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	redisService := RedisService{client: meta.(*CdsClient).apiConn}

	request := redis.NewDescribeDBInstancesRequest()
	request.SetHttpMethod("GET")
	request.InstanceUuid = common.StringPtr(data.Id())
	request.InstanceName = common.StringPtr(data.Get("instance_name").(string))
	// request.IP = common.StringPtr(data.Get("ip").(string))

	response, err := redisService.DescribeRedis(ctx, request)

	if err != nil {
		return err
	}

	if *response.Code != "Success" {
		return errors.New("not found")
	}
	log.Printf("read redis request:%v,response:%v", request.ToJsonString(), response.ToJsonString())
	//设置redis info ip 的值
	data.Set("ip", response.Data[0].IP)
	return nil
}

func resourceCdsRedisUpdate(d *schema.ResourceData, meta interface{}) error {
	//暂时不支持修改redis的参数
	return fmt.Errorf("no support modify redis conf")
}

func resourceCdsRedisDelete(data *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_redis.delete")
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	request := redis.NewDeleteDBInstanceRequest()
	request.InstanceUuid = common.StringPtr(data.Id())

	redisService := RedisService{client: meta.(*CdsClient).apiConn}
	response, err := redisService.DeleteRedis(ctx, request)

	if err != nil {
		return err
	}

	if *response.Code != "Success" {
		return errors.New(*response.Message)
	}

	if err != waitRedisDeleted(ctx, redisService, *request.InstanceUuid) {
		return err
	}

	return nil
}

func matchRedisPassGoodsId(ctx context.Context, service RedisService, architecture_type int, ram int, version string, regionId string) (int, error) {
	request := redis.NewDescribeAvailableDBConfigRequest()
	request.RegionId = common.StringPtr(regionId)

	response, err := service.DescribeAvailableDBConfig(ctx, request)

	if err != nil {
		return -1, err
	}

	for _, product := range *response.Data.Products {
		if *product.Version == version {
			for _, arch := range product.Architectures {
				if *arch.ArchitectureType == architecture_type {
					for _, role := range arch.ComputeRoles {
						for _, cpuRam := range role.Standards.CpuRam {
							if *cpuRam.RAM == ram {
								return *cpuRam.PaasGoodsId, nil
							}
						}
					}
				}
			}
		}
	}

	return -1, fmt.Errorf("RegionId %v,architecture_type %d ,ram %d ,redis_version %v can not found paasGoodsId", regionId, architecture_type, ram, version)
}

func waitRedisRunning(ctx context.Context, service RedisService, instanceUuid string) error {
	request := redis.NewDescribeDBInstancesRequest()
	request.SetHttpMethod("GET")
	request.InstanceUuid = &instanceUuid

	for {
		time.Sleep(time.Second * 15)
		response, err := service.DescribeRedis(ctx, request)
		if err != nil {
			return err
		}

		if *response.Code != "Success" {
			return errors.New(*response.Message)
		}

		for _, entry := range response.Data {
			if *entry.Status == "RUNNING" {
				return nil
			}
		}
	}
}

func waitRedisDeleted(ctx context.Context, service RedisService, instanceUuid string) error {
	request := redis.NewDescribeDBInstancesRequest()
	request.InstanceUuid = &instanceUuid
	request.SetHttpMethod("GET")
	for {
		time.Sleep(time.Second * 15)
		_, err := service.DescribeRedis(ctx, request)
		if err != nil {
			return nil
		}
	}
}
