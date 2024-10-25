package cds

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/mongodb"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCdsMongodb() *schema.Resource {
	return &schema.Resource{
		Create: resourceCdsMongodbCreate,
		Read:   resourceCdsMongodbRead,
		Update: resourceCdsMongodbUpdate,
		Delete: resourceCdsMongodbDelete,
		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Region id.",
			},
			"vdc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vdc id.",
			},
			"base_pipe_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Base pipe id.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance name.",
			},
			"cpu": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Cpu count.",
			},
			"ram": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Ram info.",
			},
			"disk_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Disk type.",
			},
			"disk_value": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Disk value.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password.",
			},
			"mongodb_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mongodb version. Available version: 4.0.3(default)、3.6.7、3.2.21.",
				Default:     "4.0.3",
			},
			"ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Mongodb ip.",
			},
			"subject_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Subject ID.",
			},
		},
		Description: "Mongodb instance.[View documentation](https://github.com/capitalonline/openapi/blob/master/%E6%96%B0%E7%89%88MongoDB%E6%A6%82%E8%A7%88.md#3createdbinstance)" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
resource "cds_mongodb" "mongodb_example" {
    region_id         = "CN_Beijing_A"
    vdc_id            = "xxx"
    base_pipe_id      = "xxx"
    instance_name     = "mongo_a"
    cpu               = 1
    ram               = 2
    disk_type         = "ssd_disk"'
    disk_value        = 100
    password          = "password"
    mongodb_version   = "4.0.3"
}

` +
			"\n```",
	}

}

func resourceCdsMongodbCreate(data *schema.ResourceData, meta interface{}) error {
	log.Println("create mongodb instance")
	defer logElapsed("resource.cds_mongodb.create")

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	mongodbService := MongodbService{client: meta.(*CdsClient).apiConn}
	request := mongodb.NewCreateDBInstanceRequest()
	region_id := data.Get("region_id").(string)
	request.RegionId = common.StringPtr(region_id)
	request.VdcId = common.StringPtr(data.Get("vdc_id").(string))
	request.BasePipeId = common.StringPtr(data.Get("base_pipe_id").(string))
	request.InstanceName = common.StringPtr(data.Get("instance_name").(string))
	request.DiskType = common.StringPtr(data.Get("disk_type").(string))
	request.DiskValue = common.IntPtr(data.Get("disk_value").(int))
	request.Password = common.StringPtr(data.Get("password").(string))
	if subject, ok := data.GetOk("subject_id"); ok {
		subjectId, ok := subject.(int)
		if !ok {
			return errors.New("subject_id must be int")
		}
		request.SubjectId = common.IntPtr(subjectId)
	}

	mongodb_version := data.Get("mongodb_version").(string)
	request.Version = common.StringPtr(mongodb_version)

	cpu := data.Get("cpu").(int)
	ram := data.Get("ram").(int)

	passGoodsId, err := matchMongodbPassGoodsId(ctx, mongodbService, region_id, mongodb_version, cpu, ram)
	if err != nil {
		return err
	}
	request.PaasGoodsId = common.IntPtr(passGoodsId)

	response, err := mongodbService.CreateMongodb(ctx, request)

	if err != nil {
		return err
	}

	if *response.Code != "Success" {
		return fmt.Errorf("create mongodb instance failed,error: %v", err)
	}
	if instance_uuid := *response.Data.InstanceUuid; instance_uuid != "" {
		data.SetId(instance_uuid)
	}
	time.Sleep(100 * time.Second)

	if err := waitMongodbRunning(ctx, mongodbService, *response.Data.InstanceUuid); err != nil {
		return err
	}

	return resourceCdsMongodbRead(data, meta)
}

func resourceCdsMongodbRead(data *schema.ResourceData, meta interface{}) error {
	log.Println("read mongodb instance")
	defer logElapsed("resource.cds_mongodb.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	service := MongodbService{client: meta.(*CdsClient).apiConn}
	request := mongodb.NewDescribeDBInstancesRequest()
	request.InstanceUuid = common.StringPtr(data.Id())
	request.InstanceName = common.StringPtr(data.Get("instance_name").(string))
	response, err := service.DescribeDBInstances(ctx, request)

	if err != nil {
		return err
	}

	if *response.Code != "Success" {
		return errors.New("not found")
	}
	log.Printf("read redis request:%v,response:%v", request.ToJsonString(), response.ToJsonString())
	//设置 mongodb的 ip
	if len(response.Data) > 0 {
		data.Set("ip", response.Data[0].IP)
	}

	return nil
}

func resourceCdsMongodbUpdate(data *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("no support modify Mongodb conf")
}

func resourceCdsMongodbDelete(data *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_mongodb.delete")
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	request := mongodb.NewDeleteDBInstanceRequest()
	request.InstanceUuid = common.StringPtr(data.Id())
	service := MongodbService{client: meta.(*CdsClient).apiConn}
	response, err := service.DeleteDBInstance(ctx, request)
	if err != nil {
		return err
	}

	if *response.Code != "Success" {
		return errors.New(*response.Message)
	}

	err = waitMongodbDeleted(ctx, service, *request.InstanceUuid)
	return err

}

func matchMongodbPassGoodsId(ctx context.Context, service MongodbService,
	region_id string, mongodb_version string, cpu int, ram int) (int, error) {
	request := mongodb.NewDescribeSpecInfoRequest()
	request.RegionId = &region_id
	response, err := service.DescribeSpecInfo(ctx, request)
	if err != nil {
		return -1, err
	}

	for _, product := range response.Data.Products {
		if *product.Version == mongodb_version {
			for _, arch := range product.Architectures {
				fmt.Printf("arch:%v", arch)
				//目前mongodb接口只提供 副本集 这个类型，后续接口有其他的再进行对接
				if *arch.ArchitectureName != "副本集" {
					continue
				}
				for _, role := range arch.ComputeRoles {
					for _, cpuRam := range role.Standards.CpuRam {
						if *cpuRam.CPU == cpu && *cpuRam.RAM == ram {
							log.Printf("cpu:%v,ram:%v", cpu, ram)
							return *cpuRam.PaasGoodsId, nil
						}
					}

				}

			}
		}
	}
	return -1, fmt.Errorf("RegionId %v, cpu %d,ram %d ,mongodb_version %v can not found paasGoodsId", region_id, cpu, ram, mongodb_version)

}

func waitMongodbRunning(ctx context.Context, service MongodbService, instanceUuid string) error {
	request := mongodb.NewDescribeDBInstancesRequest()
	request.InstanceUuid = &instanceUuid

	for {
		time.Sleep(time.Second * 15)
		response, err := service.DescribeDBInstances(ctx, request)
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

func waitMongodbDeleted(ctx context.Context, service MongodbService, instanceUuid string) error {
	request := mongodb.NewDescribeDBInstancesRequest()
	request.InstanceUuid = &instanceUuid

	for {
		time.Sleep(15 * time.Second)
		response, err := service.DescribeDBInstances(ctx, request)
		if err != nil {
			return err
		}
		if response.Data == nil {
			return nil
		}
	}
}
