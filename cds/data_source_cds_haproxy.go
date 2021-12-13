package cds

import (
	"context"
	"fmt"
	"log"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/haproxy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHaproxy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ha id",
			},
			"instance_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "instance uuid",
			},
			"ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "haproxy ip.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "instance name.",
			},
			"start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "start time.",
			},
			"end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "end time.",
			},
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Region ID.",
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Used to save results.",
			},
			"ha_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "haproxy list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "cpu count",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "create time",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "display name",
						},
						"ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ip",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "instance name",
						},
						"instance_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "instance uuid",
						},
						"link_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "link type",
						},
						"link_type_str": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "link type str",
						},
						"master_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "master info",
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "port",
						},
						"ram": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ram",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "region id",
						},
						"resource_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "resource id",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "status",
						},
						"status_str": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "status str",
						},
						"sub_product_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "sub product name",
						},
						"vdc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "vdc id",
						},
						"vdc_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "vdc name",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "version",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "project name",
						},
						"vips": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "vips list",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "vips ip",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "vips type",
									},
								},
							},
						},

						"http_listeners": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "http listeners list",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl_white_list": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "acl white list",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"backend_server": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "backend server",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "ip",
												},
												"max_conn": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "max conn",
												},
												"port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "port",
												},
												"weight": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "weight",
												},
											},
										},
									},
									"certificate_ids": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "certificate ids",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"certificate_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "certificate id",
												},
												"certificate_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "certificate name",
												},
											},
										},
									},
									"client_timeout": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "client timeout",
									},
									"client_timeout_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "client timeout unit",
									},

									"connect_timeout": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "connect timeout",
									},
									"connect_timeout_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "connect timeout unit",
									},
									"listener_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "listener mode",
									},
									"listener_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "listener name",
									},
									"listener_port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "listener port",
									},
									"max_conn": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "max conn",
									},
									"scheduler": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "scheduler",
									},
									"server_timeout": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "server timeout",
									},
									"server_timeout_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "server timeout unit",
									},
									"sticky_session": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "sticky session",
									},
								},
							},
						},

						"tcp_listeners": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "tcp listeners list",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl_white_list": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "acl white list",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"backend_server": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "backend server",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "ip",
												},
												"max_conn": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "max conn",
												},
												"port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "port",
												},
												"weight": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "weight",
												},
											},
										},
									},
									"client_timeout": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "client timeout",
									},
									"client_timeout_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "client timeout unit",
									},
									"connect_timeout": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "connect timeout",
									},
									"connect_timeout_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "connect timeout unit",
									},
									"listener_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "listener mode",
									},
									"listener_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "listener name",
									},
									"listener_port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "listener port",
									},
									"max_conn": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "max conn",
									},
									"scheduler": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "scheduler",
									},
									"server_timeout": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "server timeout",
									},
									"server_timeout_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "server timeout unit",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceHaproxyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.haproxy.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	request := haproxy.NewDescribeLoadBalancersRequest()
	var instance_uuid, ip, instanceName, startTime, endTime string
	var exist bool
	if inter, ok := d.GetOk("instance_uuid"); ok {
		instance_uuid, exist = inter.(string)
		if exist {
			request.InstanceUuid = common.StringPtr(instance_uuid)
		}
	}
	if inter, ok := d.GetOk("ip"); ok {
		ip, exist = inter.(string)
		if exist {
			request.IP = common.StringPtr(ip)
		}
	}
	if inter, ok := d.GetOk("instance_name"); ok {
		instanceName, exist = inter.(string)
		if exist {
			request.InstanceName = common.StringPtr(instanceName)
		}
	}
	if inter, ok := d.GetOk("start_time"); ok {
		startTime, exist = inter.(string)
		if exist {
			request.StartTime = common.StringPtr(startTime)
		}
	}
	if inter, ok := d.GetOk("end_time"); ok {
		endTime, exist = inter.(string)
		if exist {
			request.EndTime = common.StringPtr(endTime)
		}
	}

	response, err := haproxyService.DescribeHaproxy(ctx, request)
	if err != nil {
		return err
	}
	haproxyList := make([]map[string]interface{}, 0)
	datas := []map[string]interface{}{}

	for _, ha := range response.Data {
		strategyRequest := haproxy.NewDescribeLoadBalancerStrategysRequest()
		strategyRequest.InstanceUuid = ha.InstanceUuid
		strategtResponse, err := haproxyService.DescribeLoadBalancerStrategys(ctx, strategyRequest)
		if err != nil {
			return nil
		}

		datas = append(datas, map[string]interface{}{
			"instance": ha,
			"strategy": strategtResponse.Data,
		})

		mapping := map[string]interface{}{
			"cpu":              *ha.Cpu,
			"create_time":      *ha.CreatedTime,
			"display_name":     *ha.DisplayName,
			"ip":               *ha.IP,
			"instance_name":    *ha.InstanceName,
			"instance_uuid":    *ha.InstanceUuid,
			"link_type":        *ha.LinkType,
			"link_type_str":    *ha.LinkTypeStr,
			"master_info":      *ha.MasterInfo,
			"port":             *ha.Port,
			"ram":              *ha.Ram,
			"region_id":        *ha.RegionId,
			"resource_id":      *ha.ResourceId,
			"status":           *ha.Status,
			"status_str":       *ha.StatusStr,
			"sub_product_name": *ha.SubProductName,
			"vdc_id":           *ha.VdcId,
			"vdc_name":         *ha.VdcName,
			"version":          *ha.Version,
			"project_name":     *ha.ProjectName,
			"vips":             flattenHaproxyInstanceVipsMapping(ha.Vips),
			"http_listeners":   flattenHaproxyStrategyHttpMapping(strategtResponse.Data.HttpListeners),
			"tcp_listeners":    flattenHaproxyStrategyTcpMapping(strategtResponse.Data.TcpListeners),
		}
		haproxyList = append(haproxyList, mapping)
	}

	id := fmt.Sprintf("%s%s%s%s%s", instance_uuid, ip, instanceName, startTime, endTime)

	if id == "" {
		id = "all_ha"
	}

	d.SetId(id)
	log.Printf("ha_list:%v", haproxyList)
	err = d.Set("ha_list", haproxyList)
	if err != nil {
		log.Printf("[CRITAL]%s provider set configuration list fail, reason:%s\n ", logId, err.Error())
	}

	goodsRequest := haproxy.NewDescribeLoadBalancersSpecRequest()

	if inter, ok := d.GetOk("region_id"); ok {
		regionId, exist := inter.(string)
		if exist {
			goodsRequest.RegionId = common.StringPtr(regionId)
		}
	}

	goodsResponse, err := haproxyService.DescribeLoadBalancersSpec(ctx, goodsRequest)
	if err != nil {
		return err
	}

	zonesResponse, err := haproxyService.DescribeZones(ctx, haproxy.NewDescribeZonesRequest())
	if err != nil {
		return err
	}

	result := map[string]interface{}{
		"haproxy": datas,
		"goods":   goodsResponse.Data,
		"zones":   zonesResponse.Data,
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), result); err != nil {
			return err
		}
	}

	return nil
}
