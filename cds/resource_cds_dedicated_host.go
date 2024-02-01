package cds

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/instance"
	set "github.com/deckarep/golang-set/v2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func resourceCdsDedicatedHost() *schema.Resource {
	return &schema.Resource{
		Create: createResourceDedicatedHost,
		Read:   readResourceDedicatedHost,
		Update: updateResourceDedicatedHost,
		Delete: deleteResourceDedicatedHost,
		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Region id.",
			},
			"dedicated_host_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Dedicated host type.Host machine sells computing types.You can obtain the compute types from the cds_dedicated_host_type resource.",
			},
			"dedicated_host_good_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Dedicated host good id. Host machine sells product ID. You can obtain the compute types from the cds_dedicated_host_type resource.",
			},
			"dedicated_host_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Dedicated host name",
			},
			"dedicated_host_cpu": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Dedicated host cpu",
			},
			"dedicated_host_ram": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Dedicated host ram",
			},
			"dedicated_host_limit": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Dedicated host limit.Overcommitment ratio information.",
			},
			"amount": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Amount.",
			},
			"prepaid_month": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Prepaid month.Purchase time (unit/month)",
			},
			"auto_renew": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Auto renew. Whether to enable automatic renewal.",
			},
			"description_num": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Description num. Whether to enable appending suffix to names.",
			},
			"subject_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Subject id。Test project ID.",
			},
			"dedicated_host_id_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Dedicated host id list",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readResourceDedicatedHost(data *schema.ResourceData, meta interface{}) error {
	log.Println("read dedicated_host")
	defer logElapsed("resource.cds_dedicated_host.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)
	instanceService := InstanceService{client: meta.(*CdsClient).apiConn}

	id := data.Id()
	request := instance.NewDescribeDedicatedHostsRequest()
	idSet := set.NewSet[string]()
	// check if all dedicated host exists
	for _, str := range strings.Split(id, ",") {
		request.HostId = common.StringPtr(str)
		resp, err := instanceService.DescribeDedicatedHosts(ctx, request)
		if err != nil {
			return fmt.Errorf("describe dedicated hosts failed:%v ", err)
		}
		if len(resp.Data.HostList) < 1 {
			bytes, _ := json.Marshal(resp)
			return fmt.Errorf("describe dedicated hosts return a wrong response:%s", string(bytes))
		}
		idSet.Add(*resp.Data.HostList[0].HostId)
	}
	data.Set("dedicated_host_id_list", idSet.ToSlice())

	return nil
}

func createResourceDedicatedHost(data *schema.ResourceData, meta interface{}) error {
	log.Println("create dedicated_host")
	defer logElapsed("resource.cds_dedicated_host.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)
	instanceService := InstanceService{client: meta.(*CdsClient).apiConn}
	request := instance.NewAllocateDedicatedHostsRequest()

	region, ok := data.Get("region_id").(string)
	if !ok || len(strings.Trim(region, " ")) == 0 {
		return errors.New("region_id is invalid")
	}
	request.RegionId = common.StringPtr(region)

	dedicatedHostType, ok := data.Get("dedicated_host_type").(string)
	if !ok || len(strings.Trim(dedicatedHostType, " ")) == 0 {
		return errors.New("dedicated_host_type is invalid")
	}
	request.DedicatedHostType = common.StringPtr(dedicatedHostType)

	goodsId, ok := data.Get("dedicated_host_good_id").(int)
	if !ok || goodsId < 0 {
		return errors.New("dedicated_host_good_id is invalid")
	}
	request.DedicatedHostGoodId = common.IntPtr(goodsId)

	dedicatedHostName, ok := data.Get("dedicated_host_name").(string)
	if !ok || len(strings.Trim(dedicatedHostName, " ")) == 0 {
		return errors.New("dedicated_host_type is invalid")
	}
	request.DedicatedHostName = common.StringPtr(dedicatedHostName)

	cpu, ok := data.Get("dedicated_host_cpu").(int)
	if !ok || cpu < 0 {
		return errors.New("dedicated_host_cpu is invalid")
	}

	ram, ok := data.Get("dedicated_host_ram").(int)
	if !ok || ram < 0 {
		return errors.New("dedicated_host_ram is invalid")
	}

	dedicatedHostLimit, ok := data.Get("dedicated_host_limit").(int)
	if !ok || dedicatedHostLimit < 0 {
		return errors.New("dedicated_host_limit is invalid")
	}
	request.DedicatedHostLimit = common.IntPtr(dedicatedHostLimit)

	request.Amount = common.IntPtr(1)

	if prepaidMonth, ok := data.GetOk("prepaid_month"); ok {
		prepaidMonthValue, ok := prepaidMonth.(int)
		if !ok || prepaidMonthValue < 0 {
			return errors.New("prepaid_month is invalid")
		}
		request.PrepaidMonth = common.IntPtr(prepaidMonthValue)
	}

	if autoRenew, ok := data.GetOk("auto_renew"); ok {
		autoRenewValue, ok := autoRenew.(int)
		if !ok || autoRenewValue < 0 {
			return errors.New("auto_renew is invalid")
		}
		request.AutoRenew = common.IntPtr(autoRenewValue)
	}

	if subjectId, ok := data.GetOk("subject_id"); !ok {
		subjectIdValue, ok := subjectId.(int)
		if !ok || subjectIdValue < 0 {
			return errors.New("subject_id is invalid")
		}
		request.SubjectId = common.IntPtr(subjectIdValue)
	}

	response, err := instanceService.AllocateDedicatedHosts(ctx, request)
	if err != nil {
		return fmt.Errorf("allocate dedicated hosts failed:%v", err)
	}
	if len(response.Data) == 0 {
		bytes, _ := json.Marshal(response)
		return fmt.Errorf("allocate dedicated hosts has return wrong response: %s", string(bytes))
	}
	data.SetId(*(response.Data[0]))
	err = data.Set("dedicated_host_id_list", common.StringValues(response.Data))
	if err != nil {
		return err
	}
	return readResourceDedicatedHost(data, nil)
}

func updateResourceDedicatedHost(data *schema.ResourceData, meta interface{}) error {

	return errors.New("unsupported update dedicated host")
}

func deleteResourceDedicatedHost(data *schema.ResourceData, meta interface{}) error {
	return errors.New("unsupported delete dedicated host")
}
