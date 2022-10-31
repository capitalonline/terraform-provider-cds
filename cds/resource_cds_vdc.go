package cds

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	u "terraform-provider-cds/cds/utils"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/vdc"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCdsVdc() *schema.Resource {
	return &schema.Resource{
		Create: resourceCdsVdcCreate,
		Read:   resourceCdsVdcRead,
		Update: resourceCdsVdcUpdate,
		Delete: resourceCdsVdcDelete,
		Schema: map[string]*schema.Schema{
			"vdc_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: u.ValidateStringLengthInRange(1, 36),
			},
			"region_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"public_network": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Public Network info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipnum": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"qos": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"floatbandwidth": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"billingmethod": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"autorenew": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"public_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public Network id.",
			},
			"add_public_ip": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"delete_public_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCdsVdcCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_vdc.create")()
	fmt.Println("create vdc")
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	taskService := TaskService{client: meta.(*CdsClient).apiConn}
	name := d.Get("vdc_name").(string)
	region := d.Get("region_id").(string)
	var publicNetwork = make(map[string]interface{})
	if v, ok := d.GetOk("public_network"); ok {
		publicNetwork = v.(map[string]interface{})
	}

	taskId, err := vdcService.CreateVdc(ctx, name, region, publicNetwork)
	if err != nil {
		return err
	}

	detail, err := taskService.DescribeTask(ctx, taskId)
	if err != nil {
		return err
	}
	d.SetId(*detail.Data.ResourceID)
	return resourceCdsVdcRead(d, meta)
}

func resourceCdsVdcRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_vdc.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	request := vdc.DescribeVdcRequest()
	request.VdcId = common.StringPtr(d.Id())

	response, err := vdcService.DescribeVdc(ctx, request)
	if err != nil {
		return err
	}

	if *response.Code != "Success" {
		return errors.New(*response.Message)
	}

	if len(response.Data) == 0 {
		return errors.New("not found")
	}

	d.Set("vdc_name", *response.Data[0].VdcName)
	d.Set("region_id", *response.Data[0].RegionId)
	if len(response.Data[0].PublicNetwork) > 0 {
		if _, ok := d.GetOk("public_id"); !ok {
			d.Set("public_id", *response.Data[0].PublicNetwork[0].PublicId)
		}
	} else {
		return errors.New("not public id")
	}

	return nil
}

func resourceCdsVdcUpdate(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("update vdc")
	defer logElapsed("resource.cds_vdc.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	id := d.Id()
	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	taskService := TaskService{client: meta.(*CdsClient).apiConn}
	if d.HasChange("vdc_name") {

		d.SetPartial("vdc_name")
		_, newName := d.GetChange("vdc_name")

		request := vdc.NewModifyVdcNameRequest()
		request.VdcId = common.StringPtr(id)
		request.VdcName = common.StringPtr(newName.(string))
		_, err := vdcService.client.UseVdcClient().ModifyVdcName(request)
		if err != nil {
			return err
		}

	}

	if d.HasChange("region_id") {
		return errors.New("region_id 不支持修改")
	}

	if d.HasChange("add_public_ip") {
		request := vdc.NewAddPublicIpRequest()
		request.PublicId = common.StringPtr(d.Get("public_id").(string))
		request.Number = common.IntPtr(d.Get("add_public_ip").(int))

		response, err := vdcService.AddPublicIp(ctx, request)

		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if err != nil {
			return err
		}

		if *response.Code != "Success" {
			return errors.New(*response.Message)
		}

		taskId := response.TaskId

		_, err = taskService.DescribeTask(ctx, *taskId)
		if err != nil {
			err = fmt.Errorf("[taskId]:%v api[%s] request body [%s], 任务执行失败,请检查参数", *taskId, request.GetAction(), request.ToJsonString())
			return err
		}

	}

	if d.HasChange("delete_public_ip") {
		request := vdc.NewDeletePublicIpRequest()
		request.SegmentId = d.Get("delete_public_ip").(string)

		response, err := vdcService.DeletePublicIp(ctx, request)
		if err != nil {
			return err
		}

		if *response.Code != "Success" {
			return errors.New(*response.Message)
		}

		taskId := response.TaskId
		_, err = taskService.DescribeTask(ctx, *taskId)

		if err != nil {
			err = fmt.Errorf("[taskId]:%v api[%s] request body [%s], 任务执行失败,请检查参数", *taskId, request.GetAction(), request.ToJsonString())
			return err
		}
	}

	if d.HasChange("public_network") {

		oi, ni := d.GetChange("public_network")

		if oi == nil {
			oi = new(map[string]interface{})
		}
		if ni == nil {
			ni = new(map[string]interface{})
		}
		ois := oi.(map[string]interface{})
		nis := ni.(map[string]interface{})
		// Add public network
		if len(ois) == 0 && len(nis) > 0 {

			request := vdc.NewAddPublicNetworkRequest()
			request.VdcId = common.StringPtr(id)
			terformErr := u.Mapstructure(nis, request)
			if terformErr != nil {
				return terformErr
			}
			_, err := vdcService.client.UseVdcClient().AddPublicNetwork(request)

			if err != nil {
				return err
			}
			_ = waitVdcUpdateFinished(ctx, vdcService, id)
			return nil
		}
		// Delete public network
		if len(nis) == 0 && len(ois) > 0 {
			if publicId, ok := d.GetOk("public_id"); ok {
				publicId := publicId.(string)
				if len(publicId) > 0 {
					request := vdc.NewDeletePublicNetworkRequest()
					request.PublicId = common.StringPtr(publicId)
					_, errRet := vdcService.DeletePublicNetwork(ctx, request)
					if errRet != nil {
						return errRet
					}
					_ = waitVdcUpdateFinished(ctx, vdcService, id)
				}
			}
			return nil
		}
		publicId := d.Get("public_id").(string)

		// Update public network
		result := u.Merge(ois, nis)
		curBillingmethod := result["billingmethod"][0]
		for key, value := range result {
			if len(value) != 2 {
				continue
			}

			switch key {
			case "ipnum":
				// Add public network IP
				oldNum, _ := strconv.Atoi(value[0].(string))
				newNum, _ := strconv.Atoi(value[1].(string))
				validNums := []int{4, 8, 16, 32, 64}
				if newNum > oldNum && u.ContainsInt(validNums, newNum) {
					request := vdc.NewAddPublicIpRequest()
					request.PublicId = common.StringPtr(publicId)
					request.Number = common.IntPtr(newNum)
					taskId, errRet := vdcService.AddPublicNetworkIp(ctx, request)
					if errRet != nil {
						return errRet
					}
					resp, err := taskService.DescribeTask(ctx, taskId)
					if err != nil {
						return err
					}
					if *resp.Code != "Success" {
						return errors.New("query task failed with message " + *resp.Message)
					}
				} else if newNum < oldNum && u.ContainsInt(validNums, newNum) {
					oldValue, newValue := d.GetChange("public_network")
					oldMap := oldValue.(map[string]interface{})
					newMap := newValue.(map[string]interface{})
					newMap["ipnum"] = oldMap["ipnum"]
					d.Set("public_network", newMap)
					return errors.New("Public network IP can not be deleted with Terraform currently.")
				} else {
					oldValue, newValue := d.GetChange("public_network")
					oldMap := oldValue.(map[string]interface{})
					newMap := newValue.(map[string]interface{})
					newMap["ipnum"] = oldMap["ipnum"]
					d.Set("public_network", newMap)
					return errors.New("ipnum is invalid!")
				}
			case "name":
				continue
			case "qos":
				if curBillingmethod != "Traffic" {
					request := vdc.NewModifyPublicNetworkRequest()
					request.PublicId = common.StringPtr(publicId)
					request.Qos = common.StringPtr(value[1].(string))
					qos, err := strconv.Atoi(value[1].(string))
					if err != nil || qos <= 0 {
						oldValue, newValue := d.GetChange("public_network")
						oldMap := oldValue.(map[string]interface{})
						newMap := newValue.(map[string]interface{})
						newMap["qos"] = oldMap["qos"]
						d.Set("public_network", newMap)
						return errors.New(fmt.Sprintf("invalid value %v of qos", value[1]))
					}
					taskId, errRet := vdcService.ModifyPublicNetwork(ctx, request)
					if errRet != nil {
						return errRet
					}
					resp, err := taskService.DescribeTask(ctx, taskId)
					if err != nil {
						return err
					}
					if *resp.Code != "Success" {
						return errors.New("query task failed with message " + *resp.Message)
					}
				} else {
					return errors.New("Qos can not be modified if the billingmethod is Traffic.")
				}
			case "floatbandwidth":
				continue
			case "billingmethod":
				continue
			case "autorenew":
				i, _ := strconv.Atoi(value[0].(string))
				request := vdc.NewRenewPublicNetworkRequest()
				request.PublicId = common.StringPtr(publicId)
				request.AutoRenew = common.IntPtr(i)
				taskId, errRet := vdcService.RenewPublicNetwork(ctx, request)
				if errRet != nil {
					return errRet
				}
				resp, err := taskService.DescribeTask(ctx, taskId)
				if err != nil {
					return err
				}
				if *resp.Code != "Success" {
					return errors.New("query task failed with message " + *resp.Message)
				}
			case "type":
				continue
			}
		}

		d.SetPartial("public_network")
	}
	_ = waitVdcUpdateFinished(ctx, vdcService, id)
	return nil
}

func resourceCdsVdcDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_vdc.delete")()
	fmt.Println("delete vdc")
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)
	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	id := d.Id()

	if publicId, ok := d.GetOk("public_id"); ok {
		publicId := publicId.(string)
		if len(publicId) > 0 {
			request := vdc.NewDeletePublicNetworkRequest()
			request.PublicId = common.StringPtr(publicId)
			_, errRet := vdcService.DeletePublicNetwork(ctx, request)
			if errRet != nil {
				return errRet
			}
		}
	}
	_ = waitVdcNetWorkDeleted(ctx, vdcService, id)
	request := vdc.NewDeleteVdcRequest()
	request.VdcId = common.StringPtr(id)

	_, errRet := vdcService.DeleteVdc(ctx, request)

	if errRet != nil {
		return errRet
	}
	_ = waitVdcUpdateDeleted(ctx, vdcService, id)
	return nil
}

func waitVdcUpdateFinished(ctx context.Context, service VdcService, vdcId string, flag ...string) error {
	var timeoutLine = time.Now().Add(20 * time.Minute)
	request := vdc.DescribeVdcRequest()
	request.VdcId = common.StringPtr(vdcId)
	for {
		time.Sleep(time.Second * 5)
		if time.Now().After(timeoutLine) {
			return errors.New("wait vdc update timeout")
		}
		response, err := service.DescribeVdc(ctx, request)
		if err != nil {
			return err
		}
		if *response.Code != "Success" {
			return errors.New(fmt.Sprintf("query vdc failed with message: %v", *response.Message))
		}
		if len(response.Data) <= 0 {
			return errors.New(fmt.Sprintf("can not find vdc %v", vdcId))
		}
		if response.Data[0].VdcStatus == nil {
			continue
		}
		if *response.Data[0].VdcStatus != "ok" {
			continue
		}
		if len(response.Data[0].PublicNetwork) != 0 {
			for _, pub := range response.Data[0].PublicNetwork {
				if *pub.Status != "ok" {
					continue
				}
			}
		}
		if len(response.Data[0].PrivateNetwork) != 0 {
			for _, private := range response.Data[0].PrivateNetwork {
				if *private.Status != "ok" {
					continue
				}
			}
		}
		if len(flag) != 0 {
			data, _ := json.Marshal(response)
			log.Println(fmt.Sprintf("创建实例时查询vdc:%s", string(data)))
		}
		return nil
	}
}

func waitVdcUpdateDeleted(ctx context.Context, service VdcService, vdcId string) error {
	var timeoutLine = time.Now().Add(20 * time.Minute)
	request := vdc.DescribeVdcRequest()
	request.VdcId = common.StringPtr(vdcId)
	for {
		if time.Now().After(timeoutLine) {
			return errors.New("wait vdc update timeout")
		}
		response, err := service.DescribeVdc(ctx, request)
		if err != nil {
			return err
		}
		if *response.Code == "VDCNotFound" || (*response.Code == "Success" && len(response.Data) == 0) {
			return nil
		}
		if *response.Code != "Success" {
			return errors.New(fmt.Sprintf("query vdc failed with message: %v", *response.Message))
		}
		time.Sleep(time.Second * 5)
	}
}

func waitVdcNetWorkDeleted(ctx context.Context, service VdcService, vdcId string) error {
	var timeoutLine = time.Now().Add(20 * time.Minute)
	request := vdc.DescribeVdcRequest()
	request.VdcId = common.StringPtr(vdcId)
	for {
		time.Sleep(time.Second * 5)
		if time.Now().After(timeoutLine) {
			return errors.New("wait vdc update timeout")
		}
		response, err := service.DescribeVdc(ctx, request)
		if err != nil {
			return err
		}
		if len(response.Data[0].PublicNetwork) == 0 && len(response.Data[0].PrivateNetwork) == 0 {
			return nil
		}

	}
}
