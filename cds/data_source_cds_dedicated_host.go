package cds

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/instance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceDedicatedHost() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedHostRead,
		Schema: map[string]*schema.Schema{
			"host_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Host id.",
			},
			"host_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host name.",
			},
			"host_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host type.",
			},
			"ram_rate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Ram rate.",
			},
			"cpu_rate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cpu rate.",
			},
			"bill_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Bill method.",
			},
			"duration": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Duration.",
			},
			"end_bill_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "End bill time.",
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Region.",
			},
			"start_bill_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Start bill time.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status.",
			},
			"vm_num": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Vm num.",
			},
		},
		Description: "Data source dedicated host.",
	}
}

func dataSourceDedicatedHostRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.dedicated_host.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	instanceService := InstanceService{client: meta.(*CdsClient).apiConn}
	request := instance.NewDescribeDedicatedHostsRequest()
	id := d.Id()
	request.HostId = common.StringPtr(id)
	response, err := instanceService.DescribeDedicatedHosts(ctx, request)
	if err != nil {
		return fmt.Errorf("describe dedicated hosts err:%v", err)
	}
	if len(response.Data.HostList) == 0 {
		bytes, _ := json.Marshal(response)
		return fmt.Errorf("describe dedicated hosts return a wrong response:%s", string(bytes))
	}
	data := response.Data.HostList[0]
	if err = d.Set("host_id", data.HostId); err != nil {
		return err
	}
	if err = d.Set("host_name", data.HostName); err != nil {
		return err
	}
	if err = d.Set("host_type", data.HostType); err != nil {
		return err
	}
	if err = d.Set("ram_rate", data.RamRate); err != nil {
		return err
	}
	if err = d.Set("cpu_rate", data.CpuRate); err != nil {
		return err
	}
	if err = d.Set("bill_method", data.BillMethod); err != nil {
		return err
	}
	if err = d.Set("duration", data.Duration); err != nil {
		return err
	}

	if err = d.Set("end_bill_time", data.EndBillTime); err != nil {
		return err
	}
	if err = d.Set("region", data.Region); err != nil {
		return err
	}
	if err = d.Set("start_bill_time", data.StartBillTime); err != nil {
		return err
	}
	if err = d.Set("status", data.Status); err != nil {
		return err
	}
	if err = d.Set("vm_num", data.VmNum); err != nil {
		return err
	}
	d.SetId(*data.HostId)
	return nil
}
