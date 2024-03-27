package cds

import (
	"context"
	"errors"
	"fmt"
	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/vdc"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	u "terraform-provider-cds/cds/utils"
)

func resourceCdsPublicNetwork() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCdsPublicNetwork,
		Read:   readResourceCdsPublicNetwork,
		Update: updateResourceCdsPublicNetwork,
		Delete: deleteResourceCdsPublicNetwork,
		Schema: map[string]*schema.Schema{
			"public_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Public network id.",
			},
			"vdc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vdc id.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Public network name.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Public network type. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%99%9A%E6%8B%9F%E6%95%B0%E6%8D%AE%E4%B8%AD%E5%BF%83%E6%A6%82%E8%A7%88.md#4createpublicnetwork)",
			},
			"billing_method": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Public network address. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%99%9A%E6%8B%9F%E6%95%B0%E6%8D%AE%E4%B8%AD%E5%BF%83%E6%A6%82%E8%A7%88.md#4createpublicnetwork)",
			},
			"qos": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Public network qos. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%99%9A%E6%8B%9F%E6%95%B0%E6%8D%AE%E4%B8%AD%E5%BF%83%E6%A6%82%E8%A7%88.md#4createpublicnetwork)",
			},
			"ip_num": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The number of IPs purchased. The valid values are:4, 8, 16, 32, 64. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%99%9A%E6%8B%9F%E6%95%B0%E6%8D%AE%E4%B8%AD%E5%BF%83%E6%A6%82%E8%A7%88.md#4createpublicnetwork)",
			},
			"auto_renew": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Whether to automatically renew. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%99%9A%E6%8B%9F%E6%95%B0%E6%8D%AE%E4%B8%AD%E5%BF%83%E6%A6%82%E8%A7%88.md#4createpublicnetwork)",
			},
			"float_bandwidth": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Float bandwidth. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%99%9A%E6%8B%9F%E6%95%B0%E6%8D%AE%E4%B8%AD%E5%BF%83%E6%A6%82%E8%A7%88.md#4createpublicnetwork)",
			},
			"subject_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Subject id. ",
			},
		},
		Description: "Public network.\n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
resource "cds_public_network" "pb1" {
  ip_num          = 4
  qos             = 10
  # To identify multiple different public networks, the 'name' field is required .
  name            = "terraform-copy"
  float_bandwidth = 200
  billing_method  = "BandwIdth"
  auto_renew      = 1
  type            = "Bandwidth_Multi_ISP_BGP"
  vdc_id          = "xxxxxxxx-xxxx"
}
` +
			"\n```",
	}
}

func createResourceCdsPublicNetwork(data *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_public_network.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)
	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	taskService := TaskService{client: meta.(*CdsClient).apiConn}

	request := vdc.NewCreatePublicNetworkRequest()
	vdcId := data.Get("vdc_id")
	request.VdcId = common.StringPtr(vdcId.(string))
	_ = waitVdcUpdateFinished(ctx, vdcService, vdcId.(string))
	name := data.Get("name")
	nameStr := name.(string)
	if len(nameStr) == 0 {
		return errors.New("public network name is not valid")
	}
	request.Name = common.StringPtr(nameStr)
	publicNetworkType := data.Get("type")
	request.Type = common.StringPtr(publicNetworkType.(string))

	billingMethod := data.Get("billing_method")
	request.BillingMethod = common.StringPtr(billingMethod.(string))
	qos := data.Get("qos")
	request.Qos = common.IntPtr(qos.(int))

	ipNum := data.Get("ip_num")
	request.IPNum = common.IntPtr(ipNum.(int))

	autoRenew := data.Get("auto_renew")
	request.AutoRenew = common.IntPtr(autoRenew.(int))
	floatBandwidth := data.Get("float_bandwidth")
	request.FloatBandwidth = common.IntPtr(floatBandwidth.(int))
	if subject, ok := data.GetOk("subject_id"); ok {
		subjectId, ok := subject.(int)
		if !ok {
			return errors.New("subject_id must be int")
		}
		request.SubjectId = common.IntPtr(subjectId)
	}

	response, err := vdcService.CreatePublicNetwork(ctx, request)
	if err != nil {
		return err
	}
	if *response.Code != "Success" {
		return errors.New("request failed with code " + *response.Code)
	}
	resp, err := taskService.DescribeTask(ctx, *response.TaskId)
	if err != nil {
		return err
	}
	if *response.Code != "Success" {
		return errors.New("query task failed with message " + *resp.Message)
	}
	req := vdc.DescribeVdcRequest()
	req.VdcId = common.StringPtr(vdcId.(string))
	res, err := vdcService.DescribeVdc(ctx, req)
	if err != nil {
		return err
	}

	if *res.Code != "Success" {
		return errors.New(*res.Message)
	}
	if len(res.Data) < 1 {
		return errors.New(fmt.Sprintf("vdc %s not exist", vdcId))
	}
	for _, publicNetwork := range res.Data[0].PublicNetwork {
		if *publicNetwork.Name == data.Get("name").(string) {
			data.SetId(*publicNetwork.PublicId)
			data.Set("public_id", *publicNetwork.PublicId)
		}
	}
	return nil
}

func readResourceCdsPublicNetwork(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_public_network.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	request := vdc.DescribeVdcRequest()
	vdcId := d.Get("vdc_id")
	request.VdcId = common.StringPtr(vdcId.(string))

	response, err := vdcService.DescribeVdc(ctx, request)
	if err != nil {
		return err
	}

	if *response.Code != "Success" {
		return errors.New(*response.Message)
	}
	if len(response.Data) < 1 {
		return errors.New(fmt.Sprintf("vdc %s not exist", vdcId))
	}
	for _, publicNetwork := range response.Data[0].PublicNetwork {
		if *publicNetwork.Name == d.Get("name").(string) {
			d.SetId(*publicNetwork.PublicId)
			d.Set("public_id", *publicNetwork.PublicId)
		}
	}
	return nil
}

func updateResourceCdsPublicNetwork(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_public_network.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)
	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	taskService := TaskService{client: meta.(*CdsClient).apiConn}
	vdcId := d.Get("vdc_id")
	_ = waitVdcUpdateFinished(ctx, vdcService, vdcId.(string))
	if d.HasChange("name") {
		oldName, _ := d.GetChange("name")
		name := oldName.(string)
		d.Set("name", name)
		return errors.New("field name can not be modified")
	}
	if d.HasChange("vdc_id") {
		oldId, _ := d.GetChange("vdc_id")
		vdcId := oldId.(string)
		d.Set("vdc_id", vdcId)
		return errors.New("field vdc_id can not be modified")
	}
	if d.HasChange("type") {
		oldType, _ := d.GetChange("type")
		old := oldType.(string)
		d.Set("type", old)
		return errors.New("field type can not be modified")
	}
	if d.HasChange("billing_method") {
		oldBill, _ := d.GetChange("billing_method")
		old := oldBill.(string)
		d.Set("billing_method", old)
		return errors.New("field billing_method can not be modified")
	}
	if d.HasChange("float_bandwidth") {
		oldBandwidth, _ := d.GetChange("float_bandwidth")
		old := oldBandwidth.(string)
		d.Set("float_bandwidth", old)
		return errors.New("field float_bandwidth can not be modified")
	}
	publicId := d.Id()
	if ok := d.HasChange("ip_num"); ok {
		o, n := d.GetChange("ip_num")
		oldNum := o.(int)
		newNum := n.(int)
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
		} else {
			d.Set("ip_num", oldNum)
			if !u.ContainsInt(validNums, newNum) {
				return errors.New("ipnum is invalid!")
			}
			return errors.New("Public network IP can not be deleted with Terraform currently.")
		}
	}
	if ok := d.HasChange("qos"); ok {
		oldQos, newQos := d.GetChange("qos")
		request := vdc.NewModifyPublicNetworkRequest()
		request.PublicId = common.StringPtr(publicId)
		qosNum := newQos.(int)
		if qosNum <= 0 {
			d.Set("qos", oldQos)
			return errors.New(fmt.Sprintf("invalid value %v of qos", newQos))
		}
		request.Qos = common.StringPtr(strconv.Itoa(qosNum))
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
	}

	if ok := d.HasChange("auto_renew"); ok {
		o, n := d.GetChange("auto_renew")
		oldVal := o.(int)
		newVal := n.(int)
		if newVal != 0 && newVal != 1 {
			d.Set("auto_renew", oldVal)
			return errors.New("field must be one of [0,1]")
		}
		request := vdc.NewRenewPublicNetworkRequest()
		request.PublicId = common.StringPtr(publicId)
		request.AutoRenew = common.IntPtr(newVal)
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
	}
	return nil
}

func deleteResourceCdsPublicNetwork(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_public_network.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)
	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	//taskService := TaskService{client: meta.(*CdsClient).apiConn}
	request := vdc.NewDeletePublicNetworkRequest()

	vdcId := d.Get("vdc_id")
	id, _ := vdcId.(string)
	_ = waitVdcUpdateFinished(ctx, vdcService, id)
	publicId := d.Get("public_id")
	request.PublicId = common.StringPtr(publicId.(string))
	_, errRet := vdcService.DeletePublicNetwork(ctx, request)
	log.Println("删除公网,task_id")
	if errRet != nil {
		return errRet
	}
	_ = waitVdcUpdateFinished(ctx, vdcService, id)
	return nil
}
