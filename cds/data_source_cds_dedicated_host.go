package cds

import (
	"context"
	"errors"
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
				Optional:    true,
				Description: "Host id.",
			},
			"host_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Host name.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Output file path.",
			},
			"dedicated_hosts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Host id.",
						},
						"host_name": {
							Type:        schema.TypeString,
							Optional:    true,
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
						"result_output_file": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Output file path.",
						},
					},
				},
			},
		},
		Description: "Data source dedicated host.\n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
data "cds_data_source_dedicated_host" "test" {
	host_id = "xx"
}
` +
			"\n```",
	}
}

func dataSourceDedicatedHostRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.dedicated_host.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	instanceService := InstanceService{client: meta.(*CdsClient).apiConn}
	request := instance.NewDescribeDedicatedHostsRequest()
	if hostId, ok := d.GetOk("host_id"); ok {
		id, ok := hostId.(string)
		if !ok {
			return errors.New("host_id must be string")
		}
		request.HostId = common.StringPtr(id)
	}

	if hostName, ok := d.GetOk("host_name"); ok {
		name, ok := hostName.(string)
		if !ok {
			return errors.New("host_name must be string")
		}
		request.HostName = common.StringPtr(name)
	}

	request.PageNumber = common.IntPtr(1)
	request.PageSize = common.IntPtr(10)
	response, err := instanceService.DescribeDedicatedHosts(ctx, request)
	if err != nil {
		return fmt.Errorf("describe dedicated hosts err:%v", err)
	}
	hosts := make([]map[string]interface{}, 0, len(response.Data.HostList))
	for i := 0; i < len(response.Data.HostList); i++ {
		item := response.Data.HostList[i]
		host := map[string]interface{}{
			"host_id":         item.HostId,
			"host_name":       item.HostName,
			"host_type":       item.HostType,
			"ram_rate":        item.RamRate,
			"cpu_rate":        item.CpuRate,
			"bill_method":     item.BillMethod,
			"duration":        item.Duration,
			"end_bill_time":   item.EndBillTime,
			"region":          item.Region,
			"start_bill_time": item.StartBillTime,
			"status":          item.Status,
			"vm_num":          item.VmNum,
		}
		hosts = append(hosts, host)
	}
	if path, ok := d.GetOk("result_output_file"); ok {
		if err = writeToFile(path.(string), hosts); err != nil {
			return err
		}
	}
	if d.Id() == "" {
		id := fmt.Sprintf("cds_datasource_dedicated_hosts")
		d.SetId(id)
	}
	return nil
}
