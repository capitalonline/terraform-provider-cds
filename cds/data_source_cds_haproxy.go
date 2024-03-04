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
				Description: "Ha id.",
			},
			"instance_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance uuid.",
			},
			"ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Haproxy ip.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance name.",
			},
			"start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Start time.",
			},
			"end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "End time.",
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
				Description: "Haproxy list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Cpu count.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Display name.",
						},
						"ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Ip.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance name.",
						},
						"instance_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance uuid.",
						},
						"link_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link type.",
						},
						"link_type_str": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link type str.",
						},
						"master_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Master info.",
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Port.",
						},
						"ram": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Ram.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region id.",
						},
						"resource_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource id.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status.",
						},
						"status_str": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status str.",
						},
						"sub_product_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Sub product name.",
						},
						"vdc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vdc id.",
						},
						"vdc_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vdc name.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Version.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Project name.",
						},
						"vips": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Vips list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Vips ip.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Vips type.",
									},
								},
							},
						},

						"http_listeners": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Http listeners list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl_white_list": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Acl white list.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"backend_server": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Backend server.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Ip.",
												},
												"max_conn": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Max connection.",
												},
												"port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Port.",
												},
												"weight": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Weight.",
												},
											},
										},
									},
									"certificate_ids": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Certificate ids",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"certificate_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Certificate id.",
												},
												"certificate_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Certificate name.",
												},
											},
										},
									},
									"client_timeout": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Client timeout.",
									},
									"client_timeout_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Client timeout unit.",
									},

									"connect_timeout": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Connect timeout.",
									},
									"connect_timeout_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Connect timeout unit.",
									},
									"listener_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Listener mode.",
									},
									"listener_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Listener name.",
									},
									"listener_port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Listener port.",
									},
									"max_conn": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Max connection.",
									},
									"scheduler": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Scheduler.",
									},
									"server_timeout": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Server timeout.",
									},
									"server_timeout_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Server timeout unit.",
									},
									"sticky_session": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Sticky session.",
									},
									"session_persistence": {
										Type:       schema.TypeList,
										Computed:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										MaxItems:   1,
										MinItems:   0,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"mode": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"timer": {
													Type:       schema.TypeMap,
													Computed:   true,
													ConfigMode: schema.SchemaConfigModeAttr,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"max_idle": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"max_life": {
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
											},
										},
										Description: "Session persistence.",
									},
									"option": {
										Type:       schema.TypeMap,
										ConfigMode: schema.SchemaConfigModeAttr,
										Computed:   true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"httpchk": {
													Type:       schema.TypeMap,
													Computed:   true,
													ConfigMode: schema.SchemaConfigModeAttr,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"method": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"uri": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
										Description: "Option. ",
									},
								},
							},
						},

						"tcp_listeners": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Tcp listeners list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl_white_list": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Acl white list.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"backend_server": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Backend server.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Ip.",
												},
												"max_conn": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Max connection.",
												},
												"port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Port.",
												},
												"weight": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Weight.",
												},
											},
										},
									},
									"client_timeout": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Client timeout.",
									},
									"client_timeout_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Client timeout unit.",
									},
									"connect_timeout": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Connect timeout.",
									},
									"connect_timeout_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Connect timeout unit.",
									},
									"listener_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Listener mode.",
									},
									"listener_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Listener name.",
									},
									"listener_port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Listener port.",
									},
									"max_conn": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Max connection.",
									},
									"scheduler": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Scheduler.",
									},
									"server_timeout": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Server timeout.",
									},
									"server_timeout_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Server timeout unit.",
									},
								},
							},
						},
					},
				},
			},
		},
		Description: "Data source haproxy. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#4describeloadbalancers)",
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

		if ha.InstanceUuid != nil && *ha.InstanceUuid == instance_uuid && ha.IP != nil {
			d.Set("ip", ha.IP)
		}
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
