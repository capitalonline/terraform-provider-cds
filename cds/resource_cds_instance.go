package cds

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	u "terraform-provider-cds/cds/utils"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/instance"
	"github.com/capitalonline/cds-gic-sdk-go/security_group"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCdsCcsInstance() *schema.Resource {
	return &schema.Resource{

		Create: resourceCdsCcsInstanceCreate,
		Read:   resourceCdsCcsInstanceRead,
		Update: resourceCdsCcsInstanceUpdate,
		Delete: resourceCdsCcsInstanceDelete,
		Schema: map[string]*schema.Schema{
			"region_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CN_Beijing_A",
			},
			"instance_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: u.ValidateStringLengthInRange(1, 36),
			},
			"vdc_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "Ubuntu_16.04_64",
				Optional: true,
			},
			"cpu": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ram": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_charge_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "PostPaid",
			},
			"auto_renew": &schema.Schema{
				Type:     schema.TypeInt,
				Default:  1,
				Optional: true,
			},
			"prepaid_month": &schema.Schema{
				Type:     schema.TypeInt,
				Default:  1,
				Optional: true,
			},
			"amount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"public_ip": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "公网ip",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old != "" && new == "auto" {
						return true
					}
					return false
				},
			},
			"private_ip": &schema.Schema{
				Type:        schema.TypeList,
				ConfigMode:  schema.SchemaConfigModeAttr,
				MaxItems:    15,
				Optional:    true,
				Description: "私网IP",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"private_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if old != "" && new == "auto" {
									return true
								}
								return false
							},
						},
						"interface_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"system_disk": {
				Type:       schema.TypeMap,
				ConfigMode: schema.SchemaConfigModeAuto,
				// ConfigMode:  schema.SchemaConfigModeAttr,
				Optional:    true,
				Description: "System disk info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"size": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"iops": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"data_disks": {
				Type:       schema.TypeList,
				ConfigMode: schema.SchemaConfigModeAttr,
				Optional:   true,
				MinItems:   1,
				MaxItems:   15,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"size": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "high_disk",
						},
						"iops": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"update_data_disks": {
				Type:       schema.TypeList,
				ConfigMode: schema.SchemaConfigModeAttr,
				Optional:   true,
				MinItems:   1,
				MaxItems:   15,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"size": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "high_disk",
						},
						"iops": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"delete_data_disks": {
				Type:       schema.TypeList,
				ConfigMode: schema.SchemaConfigModeAttr,
				Optional:   true,
				MinItems:   1,
				MaxItems:   15,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"security_group_binding": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 1,
				MaxItems: 15,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"security_group_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"instance_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			//service
			"operate_instance_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"utc": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"image_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCdsCcsInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	log.Println("create instance")
	defer logElapsed("resource.cds_instance.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	instanceService := InstanceService{client: meta.(*CdsClient).apiConn}
	taskService := TaskService{client: meta.(*CdsClient).apiConn}
	securityGroupService := SecurityGroupService{client: meta.(*CdsClient).apiConn}

	createInstanceRequest := instance.NewAddInstanceRequest()

	if regionId, ok := d.GetOk("region_id"); ok {
		regionId := regionId.(string)
		if len(regionId) > 0 {
			createInstanceRequest.RegionId = common.StringPtr(regionId)
		}
	}
	if instanceName, ok := d.GetOk("instance_name"); ok {
		insName := instanceName.(string)
		if len(insName) > 0 {
			createInstanceRequest.InstanceName = common.StringPtr(insName)
		}
	}
	if vdcId, ok := d.GetOk("vdc_id"); ok {
		vdcId := vdcId.(string)
		if len(vdcId) > 0 {
			createInstanceRequest.VdcId = common.StringPtr(vdcId)
		}
	}
	if cpu, ok := d.GetOk("cpu"); ok {
		cpu := cpu.(int)
		if cpu > 0 {
			createInstanceRequest.Cpu = common.IntPtr(cpu)
		}
	}
	if ram, ok := d.GetOk("ram"); ok {
		ram := ram.(int)
		if ram > 0 {
			createInstanceRequest.Ram = common.IntPtr(ram)
		}
	}
	if instanceType, ok := d.GetOk("instance_type"); ok {
		insType := instanceType.(string)
		if len(insType) > 0 {
			createInstanceRequest.InstanceType = common.StringPtr(insType)
		}
	}
	if imageId, ok := d.GetOk("image_id"); ok {
		imageId := imageId.(string)
		if len(imageId) > 0 {
			createInstanceRequest.ImageId = common.StringPtr(imageId)
		}
	}
	if instanceChargeType, ok := d.GetOk("instance_charge_type"); ok {
		insChargeType := instanceChargeType.(string)
		if len(insChargeType) > 0 {
			createInstanceRequest.InstanceChargeType = common.StringPtr(insChargeType)
		}
	}
	if password, ok := d.GetOk("password"); ok {
		passwd := password.(string)
		if len(passwd) > 0 {
			createInstanceRequest.Password = common.StringPtr(passwd)
		}
	}
	if publicKey, ok := d.GetOk("public_key"); ok {
		publicKey := publicKey.(string)
		if len(publicKey) > 0 {
			createInstanceRequest.PublicKey = common.StringPtr(publicKey)
		}
	}
	if autoRenew, ok := d.GetOk("auto_renew"); ok {
		autoRenew := autoRenew.(int)
		if autoRenew > 0 {
			createInstanceRequest.AutoRenew = common.IntPtr(autoRenew)
		}
	}
	if prepaidMonth, ok := d.GetOk("prepaid_month"); ok {
		prepaidMonth := prepaidMonth.(int)
		if prepaidMonth > 0 {
			createInstanceRequest.PrepaidMonth = common.IntPtr(prepaidMonth)
		}
	}
	if amount, ok := d.GetOk("amount"); ok {
		amount := amount.(int)
		if amount > 0 {
			createInstanceRequest.Amount = common.IntPtr(amount)
		}
	}
	if utc, ok := d.GetOk("utc"); ok {
		utc := utc.(bool)
		createInstanceRequest.UTC = common.BoolPtr(utc)
	}
	if publicIp, ok := d.GetOk("public_ip"); ok {
		publicIp := publicIp.(string)
		if len(publicIp) > 0 {
			publicIpList := strings.Split(strings.TrimSpace(publicIp), ";")
			if len(publicIpList) > 0 {
				createInstanceRequest.PublicIp = u.MergeSlice(createInstanceRequest.PublicIp, common.StringPtrs(publicIpList))
			}
		}
	}

	if subnets, ok := d.GetOk("private_ip"); ok {
		nets := subnets.([]interface{})
		for i := range nets {
			net := nets[i].(map[string]interface{})
			ips := strings.Split(net["address"].(string), ",")
			createInstanceRequest.PrivateIp = append(createInstanceRequest.PrivateIp, &instance.PrivateIp{
				PrivateID: common.StringPtr(net["private_id"].(string)),
				IP:        common.StringPtrs(ips),
			})
		}
	}

	if v, ok := d.GetOk("system_disk"); ok {
		var sysdisk = instance.SystemDisk{}
		err := u.Mapstructure(v.(map[string]interface{}), &sysdisk)
		if err != nil {
			return err
		}
		createInstanceRequest.SystemDisk = &sysdisk
	} else {
		var sysdisk = instance.SystemDisk{}
		sysdisk.IOPS = common.IntPtr(0)
		sysdisk.Type = common.StringPtr("system_disk")
		sysdisk.Size = common.IntPtr(20)
		createInstanceRequest.SystemDisk = &sysdisk
	}

	//添加数据盘
	if addDataDisks, ok := d.GetOk("data_disks"); ok {
		disks := addDataDisks.([]interface{})
		for i := range disks {
			disk := disks[i].(map[string]interface{})
			createInstanceRequest.DataDisks = append(createInstanceRequest.DataDisks, &instance.DataDisk{
				Type: common.StringPtr(disk["type"].(string)),
				Size: common.IntPtr(disk["size"].(int)),
				IOPS: common.IntPtr(disk["iops"].(int)),
			})
		}
	}

	// 镜像密码
	if imagePassword, ok := d.GetOk("image_password"); ok {
		passwd := imagePassword.(string)
		if len(passwd) > 0 {
			createInstanceRequest.ImagePassword = common.StringPtr(passwd)
		}
	}

	taskId, errRet := instanceService.CreateInstance(ctx, createInstanceRequest)
	if errRet != nil {
		return errRet
	}
	//get create result
	time.Sleep(30 * time.Second)
	detail, errRet := taskService.DescribeTask(ctx, taskId)
	if errRet != nil {
		return errRet
	}
	nowId := strings.Join(common.StringValues(detail.Data.ResourceIds), ",")
	d.SetId(nowId)
	//if security_group_binding exist, bind it

	if securityGroupBinds, ok := d.GetOk("security_group_binding"); ok {
		binds := securityGroupBinds.(*schema.Set).List()
		for _, id := range detail.Data.ResourceIds {
			for _, v := range binds {
				bind := v.(map[string]interface{})
				joinRequest := security_group.NewJoinSecurityGroupRequest()
				joinRequest.SecurityGroupId = common.StringPtr(bind["security_group_id"].(string))

				bindData := security_group.BindData{
					InstanceId: id,
				}
				if bind["type"].(string) == "public" {
					bindData.PublicId = common.StringPtr(bind["subnet_id"].(string))
					joinRequest.BindData = append(joinRequest.BindData, &bindData)

				} else if bind["type"].(string) == "private" {
					bindData.PrivateId = common.StringPtr(bind["subnet_id"].(string))
					joinRequest.BindData = append(joinRequest.BindData, &bindData)
				}
				taskId, _ := securityGroupService.JoinSecurityGroup(ctx, joinRequest)
				log.Println("task: ", taskId)
			}
		}
	}
	waitInstanceRunning(context.Background(), instanceService, nowId)
	return resourceCdsCcsInstanceRead(d, meta)
}

func resourceCdsCcsInstanceRead(d *schema.ResourceData, meta interface{}) error {
	log.Println("read instance")
	defer logElapsed("resource.cds_instance.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)
	ids := d.Id()
	id := strings.Split(ids, ",")[0]
	instanceService := InstanceService{client: meta.(*CdsClient).apiConn}

	request := instance.NewDescribeInstanceRequest()
	request.InstanceId = common.StringPtr(id)
	response, errRet := instanceService.DescribeInstance(ctx, request)
	if errRet != nil {
		return errRet
	}

	instanceInfo := instance.DescribeReturnInfo{}

	for _, value := range response.Data.Instances {
		if *value.InstanceId == id {
			instanceInfo = *value
		}
	}
	if *(instanceInfo.InstanceId) == "" || instanceInfo.InstanceId == nil {
		return fmt.Errorf("【ERROR】%s", "Read instance info faild")
	}

	jsondata, _ := json.Marshal(instanceInfo)
	log.Printf("DEBUG_INSTANCEINFO: %s", string(jsondata))

	// for instance name
	d.Set("instance_name", *instanceInfo.InstanceName)

	// for instance spec
	d.Set("cpu", *instanceInfo.Cpu)
	d.Set("ram", *instanceInfo.Ram)

	// for instance status
	log.Printf("DEBUG_INSTANCEINFO: status: %#v", *instanceInfo.InstanceStatus)
	d.Set("instance_status", *instanceInfo.InstanceStatus)

	sysDisk := instanceInfo.Disks.SystemDisk
	if sysDisk != nil {
		sys_disk_type := ""
		if sysDisk.DiskType != nil {
			sys_disk_type = *sysDisk.DiskType
		}

		sys_disk_size := *sysDisk.Size
		str_sys_disk_size := strconv.Itoa(sys_disk_size)

		sys_disk_iops := "0"
		if sysDisk.Iops != nil {
			sys_disk_iops = strconv.Itoa(*sysDisk.Iops)
		}

		sysDiskMapping := map[string]interface{}{
			"type": common.StringPtr(sys_disk_type),
			"size": common.StringPtr(str_sys_disk_size),
			"iops": common.StringPtr(sys_disk_iops),
		}

		if err := d.Set("system_disk", sysDiskMapping); err != nil {
			return err
		}

	}

	// for PublicNetworkInterface
	// TBD: keep the data structure of user resource temporarily,
	// will change the public res config structure to slice in next big version
	publicId0 := instanceInfo.PublicNetworkInterface
	if len(publicId0) != 0 {
		publicId1 := instanceInfo.PublicNetworkInterface[0]
		d.Set("public_ip", *publicId1.IP)
	}

	// for PrivateNetworkInterface
	var privateNets []map[string]interface{}
	for _, p := range instanceInfo.PrivateNetworkInterface {
		if *p.IP != "" {
			nets := map[string]interface{}{
				"private_id":   p.PrivateId,
				"address":      p.IP,
				"interface_id": p.InterfaceId,
			}
			privateNets = append(privateNets, nets)
		}
	}
	if len(privateNets) > 0 && privateNets != nil {
		if err := d.Set("private_ip", privateNets); err != nil {
			return err
		}
	}
	time.Sleep(3 * time.Second)

	return nil
}

func resourceCdsCcsInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Println("update instance")
	defer logElapsed("resource.cds_instance.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	instanceService := InstanceService{client: meta.(*CdsClient).apiConn}
	securityGroupService := SecurityGroupService{client: meta.(*CdsClient).apiConn}
	taskService := TaskService{client: meta.(*CdsClient).apiConn}
	ids := d.Id()
	idArray := strings.Split(ids, ",")
	if len(idArray) > 1 {
		return errors.New("Batch creation does not allow modification")
	}
	id := idArray[0]
	d.Partial(true)
	if d.HasChange("region_id") {
		info := "openapi does not support modification region_id"
		log.Println(info)
		return errors.New(info)
	}
	if d.HasChange("vdc_id") {
		info := "openapi does not support modification vdc_id"
		log.Println(info)
		return errors.New(info)
	}
	if d.HasChange("public_key") {
		info := "openapi does not support modification public_key"
		log.Println(info)
		return errors.New(info)
	}

	if d.HasChange("region_id") || d.HasChange("vdc_id") || d.HasChange("public_key") {
		return errors.New("region_id/vdc_id/public_key does not support modify")
	}
	if d.HasChange("image_id") || d.HasChange("password") || d.HasChange("instance_type") ||
		d.HasChange("instance_charge_type") || d.HasChange("auto_renew") || d.HasChange("prepaid_month") {
		return errors.New(" image_id/password/instance_type/instance_charge_type/auto_renew/prepaid_month/amount does not support modify in this version")
	}
	// modify instance name
	if d.HasChange("instance_name") {
		d.SetPartial("instance_name")
		_, newName := d.GetChange("instance_name")

		request := instance.NewModifyInstanceNameRequest()
		request.InstanceId = common.StringPtr(id)
		request.InstanceName = common.StringPtr(newName.(string))
		_, err := instanceService.client.UseCvmClient().ModifyInstanceName(request)
		if err != nil {
			return err
		}
	}

	// modify private nets
	if d.HasChange("private_ip") {
		d.SetPartial("private_ip")
		err := resourceCdsInstanceUpdatePrivateIp(d, meta, id, ctx)
		if err != nil {
			return err
		}
		waitInstanceUpdated(context.Background(), instanceService, id)
	}

	// modify ModifyInstanceSpec: cpu, ram
	if d.HasChange("cpu") || d.HasChange("ram") || d.HasChange("instance_type") {
		d.SetPartial("cpu")
		d.SetPartial("ram")
		_, newCpu := d.GetChange("cpu")
		_, newRam := d.GetChange("ram")

		request := instance.NewModifyInstanceSpecRequest()
		request.InstanceId = common.StringPtr(id)
		request.Cpu = common.IntPtr(newCpu.(int))
		request.Ram = common.IntPtr(newRam.(int))
		if d.HasChange("instance_type") {
			instanceType := d.Get("instance_type")
			if instanceType != nil {
				request.InstanceType = common.StringPtr(instanceType.(string))
			}
		}
		requestdata, _ := json.Marshal(request)
		log.Printf("DEBUG_REQUEST: %s", string(requestdata))

		_, err := instanceService.client.UseCvmClient().ModifyInstanceSpec(request)
		if err != nil {
			return err
		}
		waitInstanceUpdated(context.Background(), instanceService, id)
	}

	if d.HasChange("security_group_binding") {
		d.SetPartial("security_group_binding")
		o, n := d.GetChange("security_group_binding")
		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}
		ois := o.(*schema.Set)
		nis := n.(*schema.Set)
		removeIngress := ois.Difference(nis).List()
		newIngress := nis.Difference(ois).List()

		for _, ing := range removeIngress {

			oldbind := ing.(map[string]interface{})

			request := security_group.NewLeaveSecurityGroupRequest()
			request.SecurityGroupId = common.StringPtr(oldbind["security_group_id"].(string))
			//response, err := client.LeaveSecurityGroup(request)

			if oldbind["type"].(string) == "public" {
				request.BindData = append(request.BindData, &security_group.BindData{
					InstanceId: common.StringPtr(id),
					PublicId:   common.StringPtr(oldbind["subnet_id"].(string)),
				})
				_, errRet := securityGroupService.LeaveSecurityGroup(ctx, request)
				if errRet != nil {
					return errRet
				}
			} else if oldbind["type"].(string) == "private" {
				request.BindData = append(request.BindData, &security_group.BindData{
					InstanceId: common.StringPtr(id),
					PrivateId:  common.StringPtr(oldbind["subnet_id"].(string)),
				})
				_, errRet := securityGroupService.LeaveSecurityGroup(ctx, request)
				if errRet != nil {
					return errRet
				}
			}
		}
		waitInstanceUpdated(context.Background(), instanceService, id)
		for _, ing := range newIngress {
			newbind := ing.(map[string]interface{})
			request := security_group.NewJoinSecurityGroupRequest()
			request.SecurityGroupId = common.StringPtr(newbind["security_group_id"].(string))
			if newbind["type"].(string) == "public" {
				request.BindData = append(request.BindData, &security_group.BindData{
					InstanceId: common.StringPtr(id),
					PublicId:   common.StringPtr(newbind["subnet_id"].(string)),
				})
				securityGroupService.JoinSecurityGroup(ctx, request)
			} else if newbind["type"].(string) == "private" {
				request.BindData = append(request.BindData, &security_group.BindData{
					InstanceId: common.StringPtr(id),
					PrivateId:  common.StringPtr(newbind["subnet_id"].(string)),
				})
				securityGroupService.JoinSecurityGroup(ctx, request)
			}
		}
	}

	if d.HasChange("system_disk") {
		osd, nsd := d.GetChange("system_disk")
		var sysdisk = instance.SystemDisk{}
		err := u.Mapstructure(nsd.(map[string]interface{}), &sysdisk)
		if err != nil {
			return err
		}
		var oldSysdisk = instance.SystemDisk{}
		err = u.Mapstructure(osd.(map[string]interface{}), &oldSysdisk)
		if err != nil {
			return err
		}
		var changed bool = false
		// 变更后的值为nil或者
		if sysdisk.Type != nil && *sysdisk.Type != *oldSysdisk.Type {
			changed = true
		}
		if sysdisk.Size != nil && *sysdisk.Size != *oldSysdisk.Size {
			changed = true
		}
		if sysdisk.IOPS != nil && *sysdisk.IOPS != *oldSysdisk.IOPS {
			changed = true
		}

		if changed {
			extendSdRequest := instance.NewExtendSystemDiskRequest()
			extendSdRequest.InstanceId = common.StringPtr(id)

			extendSdRequest.Size = sysdisk.Size
			extendSdRequest.IOPS = sysdisk.IOPS

			extendSdResponse, err := instanceService.client.UseCvmClient().ExtendSystemDisk(extendSdRequest)
			if err != nil {
				return err
			}

			log.Printf("system disk update action:[%v],request[%v],response[%v],err[%v]",
				extendSdRequest.GetAction(), extendSdRequest.ToJsonString(), extendSdResponse.ToJsonString(), err)

			taskId := extendSdResponse.TaskId
			_, err = taskService.DescribeTask(ctx, *taskId)
			if err != nil {
				return err
			}
		}
	}

	if d.HasChange("data_disks") {
		_, nd := d.GetChange("data_disks")
		n := make([]map[string]interface{}, 0)
		for _, v := range nd.([]interface{}) {
			n = append(n, v.(map[string]interface{}))
		}
		addList := make([]*instance.DataDisk, 0)
		for _, v := range n {
			//添加磁盘
			i := v["size"].(int)
			temp := instance.DataDisk{
				Size: &i,
				Type: common.StringPtr(v["type"].(string)),
				IOPS: common.IntPtr(v["iops"].(int)),
			}
			addList = append(addList, &temp)

		}

		if len(addList) > 0 {
			diskRequest := instance.NewCreateDiskRequest()
			diskRequest.InstanceId = common.StringPtr(id)
			diskRequest.DataDisks = addList
			respose, err := instanceService.client.UseCvmClient().CreateDisk(diskRequest)
			if err != nil {
				return err
			}
			taskId := respose.TaskId
			_, err = taskService.DescribeTask(ctx, *taskId)
			if err != nil {
				return err
			}
		}

	}

	if d.HasChange("update_data_disks") {
		_, nd := d.GetChange("update_data_disks")

		editList := make([]map[string]interface{}, 0)
		for _, v := range nd.([]interface{}) {
			editList = append(editList, v.(map[string]interface{}))
		}

		if len(editList) > 0 {
			for _, v := range editList {
				request := instance.NewResizeDiskRequest()
				request.InstanceId = common.StringPtr(id)
				request.DataSize = common.IntPtr(v["size"].(int))
				request.DiskId = common.StringPtr(v["disk_id"].(string))
				request.IOPS = common.IntPtr(v["iops"].(int))
				respose, err := instanceService.client.UseCvmClient().ResizeDisk(request)
				if err != nil {
					return err
				}
				taskId := respose.TaskId
				_, err = taskService.DescribeTask(ctx, *taskId)
				if err != nil {
					return err
				}
			}
		}
	}

	if d.HasChange("delete_data_disks") {
		_, nd := d.GetChange("delete_data_disks")
		delList := make([]string, 0)
		for _, v := range nd.([]interface{}) {
			value := v.(map[string]interface{})
			delList = append(delList, value["disk_id"].(string))
		}

		if len(delList) > 0 {
			diskRequest := instance.NewDeleteDiskRequest()
			diskRequest.InstanceId = common.StringPtr(id)
			diskRequest.DiskIds = common.StringPtrs(delList)
			respose, err := instanceService.client.UseCvmClient().DeleteDisk(diskRequest)
			if err != nil {
				return err
			}
			taskId := respose.TaskId
			_, err = taskService.DescribeTask(ctx, *taskId)
			if err != nil {
				return err
			}
		}
	}
	// reset instances password
	if d.HasChange("password") {
		_, nd := d.GetChange("password")
		request := instance.NewResetInstancesPasswordRequest()
		request.InstanceIds = common.StringPtr(id)
		request.Password = common.StringPtr(nd.(string))
		response, err := instanceService.ResetInstancesPassword(ctx, request)
		if err != nil {
			return err
		}
		if *response.Code != "Success" {
			log.Println("Reset instances password failed")
		}
		waitInstanceUpdated(ctx, instanceService, id)
	}
	// reset image
	if d.HasChange("image_id") {
		imageId := d.Get("image_id")
		request := instance.NewResetImageRequest()
		request.InstanceId = common.StringPtr(id)
		request.ImageId = common.StringPtr(imageId.(string))
		if password, ok := d.GetOk("password"); ok {
			request.Password = common.StringPtr(password.(string))
		}
		if imagePassword, ok := d.GetOk("image_password"); ok {
			request.ImagePassword = common.StringPtr(imagePassword.(string))
		}
		if publicKey, ok := d.GetOk("public_key"); ok {
			request.PublicKey = common.StringPtr(publicKey.(string))
		}
		response, err := instanceService.ResetImage(ctx, request)
		if err != nil {
			return err
		}
		if *response.Code != "Success" {
			log.Println("Reset instances password failed")
		}
		waitTaskFinished(ctx, taskService, *response.TaskId)
	}

	if d.HasChange("instance_charge_type") || d.HasChange("auto_renew") || d.HasChange("prepaid_month") {
		instanceChargeType := d.Get("instance_charge_type")
		autoRenew := d.Get("auto_renew")
		prepaidMonth := d.Get("prepaid_month")
		request := instance.NewModifyInstanceChargeTypeRequest()
		request.InstanceId = common.StringPtr(id)
		request.InstanceChargeType = common.StringPtr(instanceChargeType.(string))
		if autoRenew != nil {
			request.AutoRenew = common.IntPtr(autoRenew.(int))
		}
		if prepaidMonth != nil {
			request.PrepaidMonth = common.IntPtr(prepaidMonth.(int))
		}
		response, err := instanceService.ModifyInstanceChargeType(ctx, request)
		if err != nil {
			return err
		}
		if *response.Code != "Success" {
			log.Println("Reset instances password failed")
		}
		waitTaskFinished(ctx, taskService, *response.TaskId)
	}
	d.Partial(false)
	waitInstanceUpdated(context.Background(), instanceService, id)
	return resourceCdsCcsInstanceRead(d, meta)
}

func resourceCdsCcsInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Println("delete instance")
	defer logElapsed("resource.cds_instance.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)
	ids := d.Id()
	idArray := strings.Split(ids, ",")
	securityGroupService := SecurityGroupService{client: meta.(*CdsClient).apiConn}
	instanceService := InstanceService{client: meta.(*CdsClient).apiConn}

	if securityGroupBinds, ok := d.GetOk("security_group_binding"); ok {
		binds := securityGroupBinds.(*schema.Set).List()
		for _, value := range idArray {
			for _, v := range binds {
				bind := v.(map[string]interface{})
				request := security_group.NewLeaveSecurityGroupRequest()
				request.SecurityGroupId = common.StringPtr(bind["security_group_id"].(string))
				if bind["type"].(string) == "public" {
					request.BindData = append(request.BindData, &security_group.BindData{
						InstanceId: common.StringPtr(value),
						PublicId:   common.StringPtr(bind["subnet_id"].(string)),
					})
					_, errRet := securityGroupService.LeaveSecurityGroup(ctx, request)
					if errRet != nil {
						return errRet
					}
				} else if bind["type"].(string) == "private" {
					request.BindData = append(request.BindData, &security_group.BindData{
						InstanceId: common.StringPtr(value),
						PrivateId:  common.StringPtr(bind["subnet_id"].(string)),
					})
					_, errRet := securityGroupService.LeaveSecurityGroup(ctx, request)
					if errRet != nil {
						return errRet
					}

				}
			}
		}

	}
	time.Sleep(30 * time.Second)
	request := instance.NewDeleteInstanceRequest()
	for _, value := range idArray {
		request.InstanceIds = append(request.InstanceIds, common.StringPtr(value))
	}
	_, err := instanceService.client.UseCvmClient().DeleteInstance(request)
	if err != nil {
		return err
	}
	time.Sleep(30 * time.Second)
	return nil
}

func In_slice(val map[string]interface{}, slice []map[string]interface{}, key string) bool {
	for _, v := range slice {
		if v[key] == val[key] {
			return true
		}
	}
	return false
}

func In_slice_value(val map[string]interface{}, slice []map[string]interface{}, key string) (bool, map[string]interface{}) {
	for _, v := range slice {
		if v[key] == val[key] {
			return true, v
		}
	}
	return false, nil
}

func resourceCdsInstanceUpdatePrivateIp(
	d *schema.ResourceData, meta interface{}, id string, ctx context.Context) error {

	od, nd := d.GetChange("private_ip")
	o := make([]map[string]interface{}, 0)
	n := make([]map[string]interface{}, 0)
	for _, v := range od.([]interface{}) {
		o = append(o, v.(map[string]interface{}))
	}
	for _, v := range nd.([]interface{}) {
		n = append(n, v.(map[string]interface{}))
	}
	editList := make([]map[string]interface{}, 0)

	for _, v := range n {
		if v["address"] != "auto" {
			editList = append(editList, v)
		}
	}
	instanceService := InstanceService{client: meta.(*CdsClient).apiConn}
	for _, v := range editList {
		request := instance.NewModifyIpRequest()
		request.InstanceId = common.StringPtr(id)
		request.InterfaceId = common.StringPtr(v["interface_id"].(string))
		request.Address = common.StringPtr(v["address"].(string))
		if value, ok := d.GetOk("password"); ok {
			request.Password = common.StringPtr(value.(string))
		}
		_, err := instanceService.client.UseCvmClient().ModifyIpAddress(request)
		if err != nil {
			return err
		}
	}

	return nil
}

func waitInstanceRunning(ctx context.Context, service InstanceService, instanceUuid string) error {
	request := instance.NewDescribeInstanceRequest()
	request.InstanceId = &instanceUuid

	for {
		time.Sleep(time.Second * 15)
		response, err := service.DescribeInstance(ctx, request)
		if err != nil {
			return err
		}

		if *response.Code != "Success" {
			return errors.New(*response.Message)
		}
		for _, entry := range response.Data.Instances {
			if *entry.InstanceStatus == "running" && *entry.InstanceId == instanceUuid {
				return nil
			}
		}
	}
}

func waitInstanceUpdated(ctx context.Context, service InstanceService, instanceUuid string) error {
	request := instance.NewDescribeInstanceRequest()
	request.InstanceId = &instanceUuid

	for {
		time.Sleep(time.Second * 15)
		response, err := service.DescribeInstance(ctx, request)
		if err != nil {
			return err
		}

		if *response.Code != "Success" {
			return errors.New(*response.Message)
		}
		for _, entry := range response.Data.Instances {
			if *entry.InstanceId == instanceUuid {
				if *entry.InstanceStatus == "error" {
					return errors.New("updating instance failed")
				}
				if *entry.InstanceStatus != "updating" {
					return nil
				}
			}
		}
	}
}

func waitTaskFinished(ctx context.Context, service TaskService, taskId string) error {
	time.Sleep(time.Second * 15)
	response, err := service.DescribeTask(ctx, taskId)
	if err != nil {
		return err
	}

	if *response.Code != "Success" {
		return errors.New(*response.Message)
	}
	return nil
}
