package cds

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"terraform-provider-cds/cds-sdk-go/common"
	"terraform-provider-cds/cds-sdk-go/instance"
	"terraform-provider-cds/cds-sdk-go/security_group"
	u "terraform-provider-cds/cds/utils"

	"github.com/hashicorp/terraform/helper/schema"
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
				ForceNew: true,
				Default:  "CN_Beijing_A",
			},
			"instance_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdc_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "Ubuntu_16.04_64",
				Optional: true,
				ForceNew: true,
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_charge_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "PostPaid",
				ForceNew: true,
			},
			"auto_renew": &schema.Schema{
				Type:     schema.TypeInt,
				Default:  1,
				Optional: true,
				ForceNew: true,
			},
			"prepaid_month": &schema.Schema{
				Type:     schema.TypeInt,
				Default:  1,
				Optional: true,
				ForceNew: true,
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
			},
			"private_ip": &schema.Schema{
				Type:        schema.TypeMap,
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
						},
					},
				},
			},
			"data_disks": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 1,
				MaxItems: 15,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk_id": {
							Type:     schema.TypeString,
							Computed: true,
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
			//"private_network_interface": &schema.Schema{
			//	Type:     schema.TypeList,
			//	Computed: true,
			//	Elem: &schema.r{
			//		Schema: map[string]*schema.Schema{
			//			"interface_id": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Computed: true,
			//			},
			//			"name": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Computed: true,
			//			},
			//			"ip": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Computed: true,
			//			},
			//			"mac": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Computed: true,
			//			},
			//			"connected": &schema.Schema{
			//				Type:     schema.TypeInt,
			//				Computed: true,
			//			},
			//			"private_id": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Computed: true,
			//			},
			//		},
			//	},
			//},
			//"public_network_interface": &schema.Schema{
			//	Type:     schema.TypeList,
			//	Computed: true,
			//	Elem: &schema.r{
			//		Schema: map[string]*schema.Schema{
			//			"interface_id": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Computed: true,
			//			},
			//			"name": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Computed: true,
			//			},
			//			"ip": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Computed: true,
			//			},
			//			"mac": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Computed: true,
			//			},
			//			"connected": &schema.Schema{
			//				Type:     schema.TypeInt,
			//				Computed: true,
			//			},
			//			"public_id": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Computed: true,
			//			},
			//		},
			//	},
			//},

			//service
			"operate_instance_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCdsCcsInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("create instance")
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
	if publicIp, ok := d.GetOk("public_ip"); ok {
		publicIp := publicIp.(string)
		if len(publicIp) > 0 {
			publicIpList := strings.Split(publicIp, ";")
			if len(publicIpList) > 0 {
				createInstanceRequest.PublicIp = u.MergeSlice(createInstanceRequest.PublicIp, common.StringPtrs(publicIpList))
			}
		}
	}
	if privateSubnets, ok := d.GetOk("private_ip"); ok {
		privateSubnet := privateSubnets.(map[string]interface{})
		createInstanceRequest.PrivateIp = append(createInstanceRequest.PrivateIp, &instance.PrivateIp{
			PrivateID: common.StringPtr(privateSubnet["private_id"].(string)),
			IP:        common.StringPtrs([]string{privateSubnet["address"].(string)}),
		})
	}
	if dataDisks, ok := d.GetOk("data_disks"); ok {
		disks := dataDisks.([]interface{})
		for i := range disks {
			disk := disks[i].(map[string]interface{})
			createInstanceRequest.DataDisks = append(createInstanceRequest.DataDisks, &instance.DataDisk{
				Type: common.StringPtr(disk["type"].(string)),
				Size: common.IntPtr(disk["size"].(int)),
			})
		}
	}

	taskId, errRet := instanceService.CreateInstance(ctx, createInstanceRequest)
	if errRet != nil {
		return errRet
	}
	//get create result
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
				fmt.Println("task: ", taskId)
				//if errRet != nil {
				//	return errRet
				//}
				//_, errRet = taskService.DescribeTask(ctx, taskId)
				//if errRet != nil {
				//	return errRet
				//}
			}
		}
	}
	return resourceCdsCcsInstanceRead(d, meta)
}

func resourceCdsCcsInstanceRead(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("read instance")
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

	var listDataDisks []map[string]interface{}
	for _, p := range instanceInfo.Disks.DataDisks {
		diskMapping := map[string]interface{}{
			"disk_id": p.DiskId,
			"type":    p.DiskType,
			"size":    p.Size,
		}
		listDataDisks = append(listDataDisks, diskMapping)
	}
	if len(listDataDisks) > 0 && listDataDisks != nil {
		if err := d.Set("data_disks", listDataDisks); err != nil {
			return err
		}
	}
	time.Sleep(30 * time.Second)
	return nil
}

func resourceCdsCcsInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("update instance")
	defer logElapsed("resource.cds_instance.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	instanceService := InstanceService{client: meta.(*CdsClient).apiConn}
	securityGroupService := SecurityGroupService{client: meta.(*CdsClient).apiConn}

	ids := d.Id()
	idArray := strings.Split(ids, ",")
	if len(idArray) > 1 {
		return errors.New("Batch creation does not allow modification")
	}
	id := idArray[0]
	d.Partial(true)

	if d.HasChange("private_ip") {
		d.SetPartial("private_ip")
		_, newPrivate := d.GetChange("private_ip")
		//var interfaceId string
		//oldPrivate, newPrivate := d.GetChange("private_ip")
		address := newPrivate.(map[string]interface{})["address"]
		privateId := newPrivate.(map[string]interface{})["private_id"]

		request := instance.NewModifyIpRequest()
		request.InstanceId = common.StringPtr(id)
		request.InterfaceId = common.StringPtr(privateId.(string))
		request.Address = common.StringPtr(address.(string))
		_, err := instanceService.client.UseCvmClient().ModifyIpAddress(request)
		if err != nil {
			return err
		}
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
		time.Sleep(10 * time.Second)
		for _, ing := range newIngress {
			newbind := ing.(map[string]interface{})
			request := security_group.NewJoinSecurityGroupRequest()
			request.SecurityGroupId = common.StringPtr(newbind["security_group_id"].(string))
			if newbind["type"].(string) == "public" {
				request.BindData = append(request.BindData, &security_group.BindData{
					InstanceId: common.StringPtr(id),
					PublicId:   common.StringPtr(newbind["subnet_id"].(string)),
				})
				// TODO 需要解决偶发接口重复调用
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
	if d.HasChange("data_disks") {
		d.SetPartial("data_disks")
		od, nd := d.GetChange("data_disks")
		o := make([]map[string]interface{}, 0)
		n := make([]map[string]interface{}, 0)
		for _, v := range od.([]interface{}) {
			o = append(o, v.(map[string]interface{}))
		}
		for _, v := range nd.([]interface{}) {
			n = append(n, v.(map[string]interface{}))
		}
		addList := make([]*instance.DataDisk, 0)
		delList := make([]string, 0)
		editList := make([]map[string]interface{}, 0)
		if len(o) > len(n) {
			for _, v := range o {
				if !In_slice(v, n, "size") {
					delList = append(delList, v["disk_id"].(string))
				}
			}
			diskRequest := instance.NewDeleteDiskRequest()
			diskRequest.InstanceId = common.StringPtr(id)
			diskRequest.DiskIds = common.StringPtrs(delList)
			_, err := instanceService.client.UseCvmClient().DeleteDisk(diskRequest)
			if err != nil {
				return err
			}
		} else if len(o) < len(n) {
			for _, v := range n {
				if !In_slice(v, o, "size") {
					i, _ := strconv.Atoi(v["size"].(string))
					temp := instance.DataDisk{
						Size: &i,
						Type: common.StringPtr(v["type"].(string)),
					}
					addList = append(addList, &temp)
				}
			}
			diskRequest := instance.NewCreateDiskRequest()
			diskRequest.InstanceId = common.StringPtr(id)
			diskRequest.DataDisks = addList
			_, err := instanceService.client.UseCvmClient().CreateDisk(diskRequest)
			if err != nil {
				return err
			}
		} else {
			for _, v := range n {
				if !In_slice(v, o, "size") {
					editList = append(editList, v)
				}
			}
			for _, v := range editList {
				request := instance.NewResizeDiskRequest()
				request.InstanceId = common.StringPtr(id)
				request.DataSize = common.IntPtr(v["size"].(int))
				request.DiskId = common.StringPtr(v["disk_id"].(string))

				_, err := instanceService.client.UseCvmClient().ResizeDisk(request)
				if err != nil {
					return err
				}
			}
		}

	}
	d.Partial(false)
	time.Sleep(30 * time.Second)
	return resourceCdsCcsInstanceRead(d, meta)
}

func resourceCdsCcsInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("delete instance")
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
	//todo 等待解绑安全组
	time.Sleep(30 * time.Second)
	request := instance.NewDeleteInstanceRequest()
	for _, value := range idArray {
		request.InstanceIds = append(request.InstanceIds, common.StringPtr(value))
	}
	_, err := instanceService.client.UseCvmClient().DeleteInstance(request)
	if err != nil {
		return err
	}
	//todo 等待删除实例，删除动作当前不提供taskid
	time.Sleep(50 * time.Second)
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
