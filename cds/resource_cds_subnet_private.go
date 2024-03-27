package cds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/vdc"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCdsPrivateSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceCdsPrivateSubnetCreate,
		Read:   resourceCdsPrivateSubnetRead,
		Update: resourceCdsPrivateSubnetUpdate,
		Delete: resourceCdsPrivateSubnetDelete,
		Schema: map[string]*schema.Schema{
			"vdc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vdc id.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Private network name.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Private network type. (auto/manual), default is auto.",
			},
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Private network address.",
			},
			"mask": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Private network mask.",
			},
		},
		Description: "Private network. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%99%9A%E6%8B%9F%E6%95%B0%E6%8D%AE%E4%B8%AD%E5%BF%83%E6%A6%82%E8%A7%88.md#5createprivatenetwork)\n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
resource "cds_private_subnet" "my_private_subnet_1" {
  vdc_id  = "xxx"
  name    = "private_1"
  type    = "private"
  address = ""
  mask    = "26"
}
` +
			"\n```",
	}
}

func resourceCdsPrivateSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("create private network")
	defer logElapsed("resource.cds_subnet_private.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	taskService := TaskService{client: meta.(*CdsClient).apiConn}

	vdcSign := d.Get("vdc_id").(string)
	createPrivateNetworkRequest := vdc.NewAddPrivateNetworkRequest()
	createPrivateNetworkRequest.VdcId = common.StringPtr(vdcSign)
	_ = waitVdcUpdateFinished(ctx, vdcService, vdcSign)
	if name, ok := d.GetOk("name"); ok {
		name := name.(string)
		if len(name) > 0 {
			createPrivateNetworkRequest.Name = common.StringPtr(name)
		}
	}
	if subnetType, ok := d.GetOk("type"); ok {
		subnetType := subnetType.(string)
		if len(subnetType) > 0 {
			createPrivateNetworkRequest.Type = common.StringPtr(subnetType)
		}
	}
	if address, ok := d.GetOk("address"); ok {
		address := address.(string)
		if len(address) > 0 {
			createPrivateNetworkRequest.Addres = common.StringPtr(address)
		}
	}
	if mask, ok := d.GetOk("mask"); ok {
		mask := mask.(int)
		if mask > 0 {
			createPrivateNetworkRequest.Mask = common.IntPtr(mask)
		}
	}

	taskId, err := vdcService.CreatePrivateNetwork(ctx, createPrivateNetworkRequest)
	if err != nil {
		return err
	}
	//
	_, err = taskService.DescribeTask(ctx, taskId)
	if err != nil {
		return err
	}

	descVdcRequest := vdc.DescribeVdcRequest()
	descVdcRequest.VdcId = common.StringPtr(vdcSign)
	result, err := vdcService.DescribeVdc(ctx, descVdcRequest)
	if err != nil {
		return err
	}
	for _, value := range result.Data {
		if *value.VdcId == vdcSign {
			for _, value := range value.PrivateNetwork {
				if *value.Name == d.Get("name").(string) {
					d.SetId(*value.PrivateId)
				}
			}
		}
	}
	return resourceCdsPrivateSubnetRead(d, meta)
}

func resourceCdsPrivateSubnetRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceCdsPrivateSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	if d.HasChange("vdc_id") {
		old, _ := d.GetChange("vdc_id")
		d.Set("vdc_id", old)
	}
	return errors.New("Private network does not support modification")
}
func resourceCdsPrivateSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()
	fmt.Println("delete private network")
	defer logElapsed("resource.cds_subnet_private.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	vdcService := VdcService{client: meta.(*CdsClient).apiConn}
	request := vdc.NewDeletePrivateNetworkRequest()
	request.PrivateId = common.StringPtr(id)
	_, errRet := vdcService.DeletePrivateNetwork(ctx, request)
	if errRet != nil {
		return errRet
	}
	time.Sleep(20 * time.Second)
	return nil
}
