package cds

import (
	"context"
	"errors"
	"fmt"
	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/mysql"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceCdsMySQLReadonly() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCdsMySQLReadonly,
		Read:   readResourceCdsMySQLReadonly,
		Update: updateResourceCdsMySQLReadonly,
		Delete: deleteResourceCdsMySQLReadonly,
		Schema: map[string]*schema.Schema{
			"instance_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"paas_goods_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"test_group_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
	}
}

func createResourceCdsMySQLReadonly(data *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_mysql_readonly.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	mysqlService := MySQLService{client: meta.(*CdsClient).apiConn}
	request := mysql.NewCreateReadOnlyDBInstanceRequest()
	instanceUuid := data.Get("instance_uuid")
	request.InstanceUuid = common.StringPtr(instanceUuid.(string))
	instanceName := data.Get("instance_name")
	request.InstanceName = common.StringPtr(instanceName.(string))

	diskType := data.Get("disk_type")
	diskValue := data.Get("disk_value")

	request.DiskType = common.StringPtr(diskType.(string))
	request.DiskValue = common.IntPtr(diskValue.(int))
	paasGoodsId := data.Get("paas_goods_id")

	request.PaasGoodsId = common.IntPtr(paasGoodsId.(int))
	request.Amount = common.IntPtr(0)
	nums, ok := data.GetOk("amount")
	if ok {
		request.Amount = common.IntPtr(nums.(int))
	}

	testGroupId, ok := data.GetOk("test_group_id")
	if ok {
		request.TestGroupId = common.IntPtr(testGroupId.(int))
	}
	response, err := mysqlService.CreateReadOnlyMySQL(ctx, request)
	if err != nil {
		return err
	}
	if *response.Code != "Success" {
		fmt.Errorf("create readonly instance request,api response:%v", response)
	}
	req := mysql.NewDescribeDBInstancesRequest()
	req.InstanceUuid = request.InstanceUuid
	resp, _ := mysqlService.DescribeDBInstances(ctx, req)
	if resp != nil && resp.Data != nil {
		for _, item := range resp.Data[0].RoGroups {
			if *item.Status == "CREATING" {
				data.SetId(*item.ServiceId)
			}
		}
	}
	if err := waitMysqlReadonlyRunning(ctx, mysqlService, *request.InstanceUuid, data.Id()); err != nil {
		return err
	}
	return readResourceCdsMySQLReadonly(data, meta)
}

func readResourceCdsMySQLReadonly(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateResourceCdsMySQLReadonly(data *schema.ResourceData, meta interface{}) error {
	return nil
}
func deleteResourceCdsMySQLReadonly(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func waitMysqlReadonlyRunning(ctx context.Context, service MySQLService, instanceUuid string, readonlyInstance string) error {
	request := mysql.NewDescribeDBInstancesRequest()
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
		if response.Data == nil || len(response.Data) == 0 {
			return errors.New("response data is null")
		}
		for _, item := range response.Data[0].RoGroups {
			if *item.ServiceId == readonlyInstance && *item.Status == "RUNNING" {
				return nil
			}
		}
	}
}
