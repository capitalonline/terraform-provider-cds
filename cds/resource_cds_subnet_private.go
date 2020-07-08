package cds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/vdc"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCdsPrivateSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceCdsPrivateSubnetCreate,
		Read:   resourceCdsPrivateSubnetRead,
		Update: resourceCdsPrivateSubnetUpdate,
		Delete: resourceCdsPrivateSubnetDelete,
		Schema: map[string]*schema.Schema{
			"vdc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "private network name.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "private network type.",
			},
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "private network address.",
			},
			"mask": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "private network mask.",
			},
		},
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
