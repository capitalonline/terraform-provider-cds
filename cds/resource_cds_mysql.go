package cds

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/mysql"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCdsMySQL() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCdsMySQL,
		Read:   readResourceCdsMySQL,
		Update: updateResourceCdsMySQL,
		Delete: deleteResourceCdsMySQL,
		Schema: map[string]*schema.Schema{
			"instance_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vdc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"base_pipe_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ram": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"mysql_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"architecture_type": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"compute_type": {
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
			"ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Description: "mysql instance parameters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backup": {
				Type:        schema.TypeMap,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Description: "create backup",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_type": {
							Type:     schema.TypeString,
							Computed: true,
							Required: true,
						},
						"desc": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"db_list": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"data_backups": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time_slot": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"date_list": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"sign": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func readResourceCdsMySQL(data *schema.ResourceData, meta interface{}) error {
	log.Println("read mysql")
	defer logElapsed("resource.cds_mysql.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	mySQLService := MySQLService{client: meta.(*CdsClient).apiConn}

	request := mysql.NewDescribeDBInstancesRequest()
	request.InstanceUuid = common.StringPtr(data.Id())
	request.InstanceName = common.StringPtr(data.Get("instance_name").(string))
	response, err := mySQLService.DescribeDBInstances(ctx, request)

	if err != nil {
		return err
	}

	if *response.Code != "Success" {
		return errors.New(*response.Message)
	}

	if len(response.Data) == 0 {
		return errors.New("not found")
	}
	log.Printf("read mysql request:%v, response:%v", request.ToJsonString(), response.ToJsonString())
	data.Set("instance_name", *response.Data[0].InstanceName)
	data.Set("region_id", *response.Data[0].RegionId)
	data.Set("ip", *response.Data[0].IP)
	data.Set("disk_value", *response.Data[0].Disks)
	return nil
}

func createResourceCdsMySQL(data *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_mysql.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	mysqlService := MySQLService{client: meta.(*CdsClient).apiConn}

	paasGoodsId, err := matchMysqlPassGoodsId(ctx, mysqlService, data.Get("cpu").(int), data.Get("ram").(int), data.Get("architecture_type").(int), data.Get("compute_type").(int), data.Get("mysql_version").(string), data.Get("region_id").(string))
	if err != nil {
		return err
	}

	request := mysql.NewCreateDBInstanceRequest()
	request.PaasGoodsId = &paasGoodsId
	request.RegionId = common.StringPtr(data.Get("region_id").(string))
	request.VdcId = common.StringPtr(data.Get("vdc_id").(string))
	request.BasePipeId = common.StringPtr(data.Get("base_pipe_id").(string))
	request.InstanceName = common.StringPtr(data.Get("instance_name").(string))
	request.DiskType = common.StringPtr(data.Get("disk_type").(string))
	request.DiskValue = common.IntPtr(data.Get("disk_value").(int))
	amount := 1
	request.Amount = common.IntPtr(amount)
	if timeZone, ok := data.GetOk("time_zone"); ok {
		request.TimeZone = common.StringPtr(timeZone.(string))
	}
	response, err := mysqlService.CreateMySQL(ctx, request)
	if err != nil {
		return err
	}
	if *response.Code != "Success" {
		return fmt.Errorf("create db instance failed, error: %s", err.Error())
	}

	if len(response.Data.InstancesUuid) == 0 {
		return fmt.Errorf("create db failed")
	}

	instanceUuid := response.Data.InstancesUuid[0]

	data.SetId(instanceUuid)
	data.Set("instance_uuid", instanceUuid)
	if err := waitMysqlRunning(ctx, mysqlService, instanceUuid); err != nil {
		return err
	}
	if dbParameters, ok := data.GetOk("parameters"); ok {
		parameterList := dbParameters.([]interface{})
		request := mysql.NewModifyDBParameterRequest()
		request.InstanceUuid = common.StringPtr(instanceUuid)
		var parameters = make([]*mysql.ModifyDBParameterParameters, 0, len(instanceUuid))
		for _, item := range parameterList {
			parameterMap := item.(map[string]interface{})
			parameters = append(parameters, &mysql.ModifyDBParameterParameters{
				Name:  common.StringPtr(parameterMap["name"].(string)),
				Value: common.StringPtr(parameterMap["value"].(string)),
			})
		}
		request.Parameters = parameters
		if request.Parameters != nil && len(request.Parameters) > 0 {
			response, err := mysqlService.ModifyDBParameter(ctx, request)
			if err != nil {
				return err
			}
			if *response.Code != "Success" {
				return fmt.Errorf("modify db parameters failed, error: %s", err.Error())
			}
		}
		if err := waitMysqlRunning(ctx, mysqlService, instanceUuid); err != nil {
			return err
		}
	}
	if backup, ok := data.GetOk("backup"); ok {
		request := mysql.NewCreateBackupRequest()
		backupMap := backup.(map[string]interface{})
		request.InstanceUuid = common.StringPtr(instanceUuid)
		if backupMap["backup_type"] != nil {
			request.BackupType = common.StringPtr(backupMap["backup_type"].(string))
		}
		if backupMap["desc"] != nil {
			request.Desc = common.StringPtr(backupMap["desc"].(string))
		}
		if backupMap["db_list"] != nil {
			dbStr := backupMap["db_list"].(string)
			dbStr = strings.Trim(strings.Trim(dbStr, "["), "]")
			dbList := strings.Split(dbStr, ",")
			request.DBList = dbList
		}
		response, err := mysqlService.CreateBackup(ctx, request)
		if err != nil {
			return err
		}
		if *response.Code != "Success" {
			fmt.Errorf("create backup failed, response: %v", response)
		}
		if err := waitMysqlRunning(ctx, mysqlService, instanceUuid); err != nil {
			return err
		}
	}
	if dataBackups, ok := data.GetOk("data_backups"); ok {
		request := mysql.NewModifyDbBackupPolicyRequest()
		backupMap := dataBackups.(map[string]interface{})
		request.InstanceUuid = common.StringPtr(instanceUuid)
		var dataBackups = new(mysql.ModifyDbBackupPolicyDataBackups)
		if backupMap["time_slot"] != nil {
			dataBackups.TimeSlot = common.StringPtr(backupMap["time_slot"].(string))
		}
		if backupMap["sign"] != nil {
			signStr := backupMap["sign"].(string)
			sign, _ := strconv.Atoi(signStr)
			dataBackups.Sign = common.IntPtr(sign)
		}
		if backupMap["date_list"] != nil {
			dbStr := backupMap["date_list"].(string)
			dbStr = strings.Trim(strings.Trim(dbStr, "["), "]")
			dataList := strings.Split(dbStr, ",")
			dataBackups.DateList = dataList
		}
		request.DataBackups = dataBackups
		response, err := mysqlService.ModifyDbBackupPolicy(ctx, request)
		if err != nil {
			return err
		}
		if *response.Code != "Success" {
			fmt.Errorf("create backup failed, response: %v", response)
		}
		if err := waitMysqlRunning(ctx, mysqlService, instanceUuid); err != nil {
			return err
		}
	}
	return readResourceCdsMySQL(data, meta)
}

func updateResourceCdsMySQL(data *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_mysql.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	if data.HasChange("region_id") {
		o_region_id, _ := data.GetChange("region_id")
		data.Set("region_id", o_region_id)
		return fmt.Errorf("region_id %s not support modify with openapi", data.Get("region_id").(string))
	}

	if data.HasChange("vdc_id") {
		o_vdc_id, _ := data.GetChange("vdc_id")
		data.Set("vdc_id", o_vdc_id)
		return fmt.Errorf("vdc_id %s not support modify with openapi", data.Get("vdc_id").(string))
	}

	if data.HasChange("base_pipe_id") {
		o_base_pipe_id, _ := data.GetChange("base_pipe_id")
		data.Set("base_pipe_id", o_base_pipe_id)
		return fmt.Errorf("base_pipe_id %s not support modify with openapi", data.Get("base_pipe_id").(string))
	}

	if data.HasChange("instance_name") {
		o_instance_name, _ := data.GetChange("instance_name")
		data.Set("instance_name", o_instance_name)
		return fmt.Errorf("instance_name %s not support modify with openapi", data.Get("instance_name").(string))
	}

	if data.HasChange("disk_type") {
		o_disk_type, _ := data.GetChange("disk_type")
		data.Set("disk_type", o_disk_type)
		return fmt.Errorf("disk type %s can not change with openapi", data.Get("disk_type").(string))
	}

	mysqlService := MySQLService{client: meta.(*CdsClient).apiConn}

	paasGoodsId, err := matchMysqlPassGoodsId(ctx, mysqlService, data.Get("cpu").(int), data.Get("ram").(int), data.Get("architecture_type").(int), data.Get("compute_type").(int), data.Get("mysql_version").(string), data.Get("region_id").(string))
	if err != nil {
		return err
	}

	request := mysql.NewModifyDBInstanceSpecRequest()

	request.InstanceUuid = common.StringPtr(data.Id())

	var hasChange bool

	if data.HasChange("cpu") || data.HasChange("ram") {
		request.PaasGoodsId = common.IntPtr(paasGoodsId)
		hasChange = true
	}

	if data.HasChange("disk_value") {
		hasChange = true
		o_disk_value, n_disk_value := data.GetChange("disk_value")
		o_val, ok := o_disk_value.(int)
		if !ok {
			return fmt.Errorf("old disk value %v is not int", o_disk_value)
		}
		n_val, ok := n_disk_value.(int)

		if !ok {
			return fmt.Errorf("new disk value %v is not int", n_disk_value)
		}
		add_disk := n_val - o_val
		request.DiskType = common.StringPtr(data.Get("disk_type").(string))
		request.DiskValue = common.IntPtr(add_disk)
	}

	if hasChange {
		response, err := mysqlService.ModifyDBInstanceSpec(ctx, request)
		if err != nil {
			return err
		}

		if *response.Code != "Success" {
			return errors.New(*response.Message)
		}

		if err := waitMysqlRunning(ctx, mysqlService, data.Id()); err != nil {
			return err
		}
	}
	if data.HasChange("parameters") {
		_, dbParameters := data.GetChange("parameters")
		request := mysql.NewModifyDBParameterRequest()
		request.InstanceUuid = common.StringPtr(data.Id())
		parametersList := dbParameters.([]interface{})
		parameters := make([]*mysql.ModifyDBParameterParameters, 0, len(parametersList))
		for _, item := range parametersList {
			parameter := item.(map[string]interface{})
			parameters = append(parameters, &mysql.ModifyDBParameterParameters{
				Name:  common.StringPtr(parameter["name"].(string)),
				Value: common.StringPtr(parameter["value"].(string)),
			})
		}
		if len(parameters) > 0 {
			request.Parameters = parameters
			response, err := mysqlService.ModifyDBParameter(ctx, request)
			if err != nil {
				return err
			}
			if *response.Code != "Success" {
				return fmt.Errorf("modify db parameters failed, error: %s", err.Error())
			}
		}
		if err := waitMysqlRunning(ctx, mysqlService, data.Id()); err != nil {
			return err
		}
	}
	if data.HasChange("backup") {
		_, newBackup := data.GetChange("backup")
		request := mysql.NewCreateBackupRequest()
		backupMap := newBackup.(map[string]interface{})
		request.InstanceUuid = common.StringPtr(data.Id())
		if backupMap["backup_type"] != nil {
			request.BackupType = common.StringPtr(backupMap["backup_type"].(string))
		}
		if backupMap["desc"] != nil {
			request.Desc = common.StringPtr(backupMap["desc"].(string))
		}
		if backupMap["db_list"] != nil {
			dbStr := backupMap["db_list"].(string)
			dbStr = strings.Trim(strings.Trim(dbStr, "["), "]")
			dbList := strings.Split(dbStr, ",")
			request.DBList = dbList
		}
		response, err := mysqlService.CreateBackup(ctx, request)
		if err != nil {
			return err
		}
		if *response.Code != "Success" {
			fmt.Errorf("create backup failed, response: %v", response)
		}
		if err := waitMysqlRunning(ctx, mysqlService, data.Id()); err != nil {
			return err
		}
	}
	// auto backup policy
	if data.HasChange("data_backups") {
		_, dataBackupsMap := data.GetChange("data_backups")
		request := mysql.NewModifyDbBackupPolicyRequest()
		backupMap := dataBackupsMap.(map[string]interface{})
		request.InstanceUuid = common.StringPtr(data.Id())
		var dataBackups = new(mysql.ModifyDbBackupPolicyDataBackups)
		if backupMap["time_slot"] != nil {
			dataBackups.TimeSlot = common.StringPtr(backupMap["time_slot"].(string))
		}
		if backupMap["sign"] != nil {
			sign, ok := backupMap["sign"].(int)
			if ok {
				dataBackups.Sign = common.IntPtr(sign)
			} else {
				signStr, _ := backupMap["sign"].(string)
				sign, _ = strconv.Atoi(signStr)
				dataBackups.Sign = common.IntPtr(sign)
			}
		}
		if backupMap["date_list"] != nil {
			dbStr := backupMap["date_list"].(string)
			dbStr = strings.Trim(strings.Trim(dbStr, "["), "]")
			dataList := strings.Split(dbStr, ",")
			dataBackups.DateList = dataList
			request.DataBackups = dataBackups
		}
		response, err := mysqlService.ModifyDbBackupPolicy(ctx, request)
		if err != nil {
			return err
		}
		if *response.Code != "Success" {
			fmt.Errorf("create backup failed, response: %v", response)
		}
		if err := waitMysqlRunning(ctx, mysqlService, data.Id()); err != nil {
			return err
		}
	}
	return nil
}

func deleteResourceCdsMySQL(data *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_mysql.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	request := mysql.NewDeleteDBInstanceRequest()
	request.InstanceUuid = common.StringPtr(data.Id())

	mysqlService := MySQLService{client: meta.(*CdsClient).apiConn}

	response, err := mysqlService.DeleteMySQL(ctx, request)
	if err != nil {
		return err
	}

	if *response.Code != "Success" {
		return errors.New(*response.Message)
	}

	if err := waitMysqlDeleted(ctx, mysqlService, data.Id()); err != nil {
		return err
	}
	return nil
}

func matchMysqlPassGoodsId(ctx context.Context, service MySQLService, cpu, ram int, architectureType, computeType int, mysqlVersion string, regionId string) (int, error) {
	goodsRequest := mysql.NewDescribeAvailableDBConfigRequest()

	goodsRequest.RegionId = common.StringPtr(regionId)

	goodsResponse, err := service.DescribeAvailableDBConfig(ctx, goodsRequest)
	if err != nil {
		return -1, err
	}

	for _, product := range goodsResponse.Data.Products {
		if *product.Version == mysqlVersion {
			for _, arch := range product.Architectures {
				if *arch.ArchitectureType == architectureType {
					for _, role := range arch.ComputeRoles {
						if *role.ComputeType == computeType {
							for _, cpuRam := range role.Standards.CpuRam {
								if *cpuRam.CPU == cpu && *cpuRam.RAM == ram {
									return *cpuRam.PaasGoodsId, nil
								}
							}
						}
					}
				}
			}
		}
	}

	return -1, fmt.Errorf("RegionId %v,architectureType %d , computeType %d ,cpu %d, ram %d not found paas_goods",
		regionId, architectureType, computeType, cpu, ram)
}

func waitMysqlRunning(ctx context.Context, service MySQLService, instanceUuid string) error {
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

		for _, entry := range response.Data {
			if *entry.Status == "RUNNING" {
				return nil
			}
		}
	}
}

func waitMysqlDeleted(ctx context.Context, service MySQLService, instanceUuid string) error {
	request := mysql.NewDescribeDBInstancesRequest()
	request.InstanceUuid = &instanceUuid

	for {
		time.Sleep(time.Second * 15)
		_, err := service.GetMySQLList(ctx, request)
		if err != nil {
			return nil
		}
	}
}
