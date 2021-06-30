package cds

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/haproxy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCdsHaproxy() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCdsHaproxy,
		Read:   readResourceCdsHaproxy,
		Update: updateResourceCdsHaproxy,
		Delete: deleteRresourceCdsHaproxy,
		Schema: map[string]*schema.Schema{
			"instance_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance UUID.",
			},
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "regon id.",
			},
			"vdc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "vdc id.",
			},
			"base_pipe_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "base pipe id.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance name.",
			},
			"paas_goods_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "pass goods id.",
			},
			"ips": {
				Type:       schema.TypeList,
				ConfigMode: schema.SchemaConfigModeAttr,
				Required:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pipe_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"pipe_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"segment_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"strategies": {
				Type:       schema.TypeList,
				ConfigMode: schema.SchemaConfigModeAttr,
				Required:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_listeners": {
							Type:        schema.TypeList,
							ConfigMode:  schema.SchemaConfigModeAttr,
							MaxItems:    15,
							Optional:    true,
							Description: "http listeners",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl_white_list": {
										Type:     schema.TypeString,
										Required: true,
									},
									"backend_server": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"max_conn": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"port": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"weight": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"certificate_ids": {
										Type:       schema.TypeList,
										Required:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"certificate_id": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"certificate_name": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"client_timeout": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"client_timeout_unit": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"connect_timeout": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"connect_timeout_unit": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"listener_mode": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"listener_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"listener_port": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"max_conn": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"scheduler": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"server_timeout": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"server_timeout_unit": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"sticky_session": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"tcp_listeners": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "base pipe id.",
							ConfigMode:  schema.SchemaConfigModeAttr,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl_white_list": {
										Type:     schema.TypeString,
										Required: true,
									},
									"backend_server": {
										Type:       schema.TypeList,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"max_conn": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"port": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"weight": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"client_timeout": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"client_timeout_unit": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"connect_timeout": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"connect_timeout_unit": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"listener_mode": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"listener_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"listener_port": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"max_conn": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"scheduler": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"server_timeout": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"server_timeout_unit": {
										Type:     schema.TypeString,
										Optional: true,
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

func createResourceCdsHaproxy(data *schema.ResourceData, meta interface{}) error {
	log.Println("create haproxy")
	defer logElapsed("resource.cds_haproxy.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	request := haproxy.NewCreateLoadBalancerRequest()

	if inter, ok := data.GetOk("region_id"); ok {
		regionId, exist := inter.(string)
		if exist {
			request.RegionId = common.StringPtr(regionId)
		}
	}
	if inter, ok := data.GetOk("vdc_id"); ok {
		vdcId, exist := inter.(string)
		if exist {
			request.VdcId = common.StringPtr(vdcId)
		}
	}
	if inter, ok := data.GetOk("base_pipe_id"); ok {
		basePipeId, exist := inter.(string)
		if exist {
			request.BasePipeId = common.StringPtr(basePipeId)
		}
	}
	if inter, ok := data.GetOk("instance_name"); ok {
		instanceName, exist := inter.(string)
		if exist {
			request.InstanceName = common.StringPtr(instanceName)
		}
	}
	if inter, ok := data.GetOk("paas_goods_id"); ok {
		paasGoodsId, exist := inter.(int)
		if exist {
			request.PaasGoodsId = common.IntPtr(paasGoodsId)
		}
	}
	if inter, ok := data.GetOk("ips"); ok {
		ips, exist := inter.([]interface{})
		if exist {
			for _, ip := range ips {
				ipEntry, isExist := ip.(map[string]interface{})
				if isExist {
					entry := &haproxy.CreateLoadBalancerIps{}
					if ipEntry["pipe_type"] != nil {
						entry.PipeType = common.StringPtr(ipEntry["pipe_type"].(string))
					}
					if ipEntry["pipe_id"] != nil {
						entry.PipeId = common.StringPtr(ipEntry["pipe_id"].(string))
					}
					if ipEntry["pipe_type"].(string) == "public" {
						if ipEntry["segment_id"] != nil {
							entry.SegmentId = common.StringPtr(ipEntry["segment_id"].(string))
						}
					}
					request.Ips = append(request.Ips, entry)
				}
			}
		}
	}
	request.Amount = common.IntPtr(1)

	_, err := haproxyService.CreateHaproxy(ctx, request)
	if err != nil {
		return err
	}

	data.SetId("Need to Set Instance UUID" + strconv.FormatInt(time.Now().Unix(), 10))

	// TODO wait until task id has effect
	// taskService := TaskService{client: meta.(*CdsClient).apiConn}
	// detail, err := taskService.DescribeTask(ctx, response.TaskId)
	// if err != nil {
	// 	return err
	// }

	// if *detail.Code != "Success" {
	// 	return fmt.Errorf("[ERROR] create haproxy task info get failed, error: %s", *detail.Message)
	// }

	// data.SetId(strings.Join(common.StringValues(detail.Data.ResourceIds), ","))
	// time.Sleep(time.Second * 10)
	// return readResourceCdsHaproxy(data, meta)
	return nil
}

func readResourceCdsHaproxy(data *schema.ResourceData, meta interface{}) error {
	log.Println("read haproxy")
	defer logElapsed("resource.cds_haproxy.read")()
	// logId := getLogId(contextNil)
	// ctx := context.WithValue(context.TODO(), "logId", logId)

	// if !strings.Contains(data.Id(), "Need to Set Instance UUID") {
	// 	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	// 	request := haproxy.NewDescribeLoadBalancersRequest()

	// 	response, err := haproxyService.DescribeHaproxy(ctx, request)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	for _, d := range response.Data {
	// 		if request.InstanceUuid == d.InstanceUuid {
	// 			data.Set("display_name", d.DisplayName)
	// 			data.Set("ip", d.IP)
	// 			data.Set("instance_uuid", d.InstanceUuid)
	// 			data.Set("link_type", d.LinkType)
	// 			data.Set("master_info", d.MasterInfo)
	// 			data.Set("port", d.Port)
	// 			data.Set("resource_id", d.ResourceId)
	// 			data.Set("status", d.Status)
	// 			data.Set("sub_product_name", d.SubProductName)
	// 			data.Set("version", d.Version)
	// 			break
	// 		}
	// 	}
	// }

	return nil
}

func updateResourceCdsHaproxy(data *schema.ResourceData, meta interface{}) error {
	log.Println("create instance")
	defer logElapsed("resource.cds_instance.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	if data.HasChange("instance_uuid") {
		data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	}

	if !strings.Contains(data.Id(), "Need to Set Instance UUID") {
		if data.HasChanges("paas_goods_id") {

			request := haproxy.NewModifyLoadBalancerInstanceSpecRequest()
			if paasGoodsId, ok := data.GetOk("paas_goods_id"); ok {
				request.PaasGoodsId = common.IntPtr(paasGoodsId.(int))
			}

			if instanceUUID, ok := data.GetOk("instance_uuid"); ok {
				request.InstanceUuid = common.StringPtr(instanceUUID.(string))
			}

			_, err := haproxyService.ModifyHaproxy(ctx, request)
			if err != nil {
				return err
			}
		}

		modifyStrategyRequest := createModitifyStrategyRequest(data)
		_, err := haproxyService.ModifyHaproxyStrategy(ctx, modifyStrategyRequest)
		if err != nil {
			return err
		}
	} else {
		return errors.New("You must update instance_uuid in cds_haproxy block by your self, it will fix in the future")
	}

	time.Sleep(time.Second * 5)

	return readResourceCdsHaproxy(data, meta)
}

func deleteRresourceCdsHaproxy(data *schema.ResourceData, meta interface{}) error {
	log.Println("create instance")
	defer logElapsed("resource.cds_instance.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	if !strings.Contains(data.Id(), "Need to Set Instance UUID") {
		haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

		request := haproxy.NewDeleteLoadBalancerRequest()
		if instanceUUID, ok := data.GetOk("instance_uuid"); ok {
			request.InstanceUuid = common.StringPtr(instanceUUID.(string))
		}

		_, err := haproxyService.DeleteHaproxy(ctx, request)
		if err != nil {
			return err
		}

	} else {
		return errors.New("You must update instance_uuid in cds_haproxy block by your self, it will fix in the future")
	}
	return nil
}

func createModitifyStrategyRequest(data *schema.ResourceData) *haproxy.ModifyLoadBalancerStrategysRequest {
	request := haproxy.NewModifyLoadBalancerStrategysRequest()

	if inter, ok := data.GetOk("instance_uuid"); ok {
		instanceUuid, exist := inter.(string)
		if exist {
			request.InstanceUuid = common.StringPtr(instanceUuid)
		}
	}

	inter, ok := data.GetOk("strategies")
	if !ok {
		return request
	}
	strategies, ok := inter.([]interface{})
	if !ok {
		return request
	}

	for _, inter := range strategies {
		strategy, ok := inter.(map[string]interface{})
		if !ok {
			break
		}

		if inter, ok := strategy["http_listeners"]; ok {
			datas := inter.([]interface{})
			listeners := make([]*haproxy.DescribeLoadBalancerStrategysHttpListeners, 0)
			for _, data := range datas {
				if data == nil {
					continue
				}
				dataMap := data.(map[string]interface{})

				backendServers := make([]*haproxy.DescribeLoadBalancerStrategysBackendServer, 0)
				if dataMap["backend_server"] != nil {
					backendServersDatas := dataMap["backend_server"].([]interface{})
					for _, backendServersData := range backendServersDatas {
						backendServerMap := backendServersData.(map[string]interface{})
						backendServers = append(backendServers, &haproxy.DescribeLoadBalancerStrategysBackendServer{
							IP:      common.StringPtr(backendServerMap["ip"].(string)),
							MaxConn: common.IntPtr(backendServerMap["max_conn"].(int)),
							Port:    common.IntPtr(backendServerMap["port"].(int)),
							Weight:  common.StringPtr(backendServerMap["weight"].(string)),
						})
					}
				}

				certificateIds := make([]*haproxy.DescribeLoadBalancerStrategysCertificateIds, 0)
				if dataMap["certificate_ids"] != nil {
					certificateIdsDatas := dataMap["certificate_ids"].([]interface{})
					for _, certificateIdsData := range certificateIdsDatas {
						certificateIdsMap := certificateIdsData.(map[string]interface{})
						certificateIds = append(certificateIds, &haproxy.DescribeLoadBalancerStrategysCertificateIds{
							CertificateId:   common.StringPtr(certificateIdsMap["certificate_id"].(string)),
							CertificateName: common.StringPtr(certificateIdsMap["certificate_name"].(string)),
						})
					}
				}

				listener := &haproxy.DescribeLoadBalancerStrategysHttpListeners{
					BackendServer:  backendServers,
					CertificateIds: certificateIds,
				}

				if dataMap["acl_white_list"] != nil {
					listener.AclWhiteList = common.StringPtrs(strings.Split(dataMap["acl_white_list"].(string), ","))
				}
				if dataMap["client_timeout"] != nil {
					listener.ClientTimeout = common.StringPtr(dataMap["client_timeout"].(string))
				}
				if dataMap["client_timeout_unit"] != nil {
					listener.ClientTimeoutUnit = common.StringPtr(dataMap["client_timeout_unit"].(string))
				}
				if dataMap["connect_timeout"] != nil {
					listener.ConnectTimeout = common.StringPtr(dataMap["connect_timeout"].(string))
				}
				if dataMap["connect_timeout_unit"] != nil {
					listener.ConnectTimeoutUnit = common.StringPtr(dataMap["connect_timeout_unit"].(string))
				}
				if dataMap["listener_mode"] != nil {
					listener.ListenerMode = common.StringPtr(dataMap["listener_mode"].(string))
				}
				if dataMap["listener_name"] != nil {
					listener.ListenerName = common.StringPtr(dataMap["listener_name"].(string))
				}
				if dataMap["listener_port"] != nil {
					listener.ListenerPort = common.IntPtr(dataMap["listener_port"].(int))
				}
				if dataMap["max_conn"] != nil {
					listener.MaxConn = common.IntPtr(dataMap["max_conn"].(int))
				}
				if dataMap["scheduler"] != nil {
					listener.Scheduler = common.StringPtr(dataMap["scheduler"].(string))
				}
				if dataMap["server_timeout"] != nil {
					listener.ServerTimeout = common.StringPtr(dataMap["server_timeout"].(string))
				}
				if dataMap["server_timeout_unit"] != nil {
					listener.ServerTimeoutUnit = common.StringPtr(dataMap["server_timeout_unit"].(string))
				}
				if dataMap["sticky_session"] != nil {
					listener.StickySession = common.StringPtr(dataMap["sticky_session"].(string))
				}
				listeners = append(listeners, listener)
			}
			request.HttpListeners = listeners
		}

		if inter, ok := strategy["tcp_listeners"]; ok {
			datas := inter.([]interface{})
			listeners := make([]*haproxy.DescribeLoadBalancerStrategysTcpListeners, 0)
			for _, data := range datas {
				if data == nil {
					continue
				}
				dataMap := data.(map[string]interface{})

				backendServers := make([]*haproxy.DescribeLoadBalancerStrategysBackendServer, 0)
				if dataMap["backend_server"] != nil {
					backendServersDatas := dataMap["backend_server"].([]interface{})
					for _, backendServersData := range backendServersDatas {
						backendServerMap := backendServersData.(map[string]interface{})
						backendServers = append(backendServers, &haproxy.DescribeLoadBalancerStrategysBackendServer{
							IP:      common.StringPtr(backendServerMap["ip"].(string)),
							MaxConn: common.IntPtr(backendServerMap["max_conn"].(int)),
							Port:    common.IntPtr(backendServerMap["port"].(int)),
							Weight:  common.StringPtr(backendServerMap["weight"].(string)),
						})
					}
				}

				listener := &haproxy.DescribeLoadBalancerStrategysTcpListeners{
					BackendServer: backendServers,
				}

				if dataMap["acl_white_list"] != nil {
					listener.AclWhiteList = common.StringPtrs(strings.Split(dataMap["acl_white_list"].(string), ","))
				}
				if dataMap["client_timeout"] != nil {
					listener.ClientTimeout = common.StringPtr(dataMap["client_timeout"].(string))
				}
				if dataMap["client_timeout_unit"] != nil {
					listener.ClientTimeoutUnit = common.StringPtr(dataMap["client_timeout_unit"].(string))
				}
				if dataMap["connect_timeout"] != nil {
					listener.ConnectTimeout = common.StringPtr(dataMap["connect_timeout"].(string))
				}
				if dataMap["connect_timeout_unit"] != nil {
					listener.ConnectTimeoutUnit = common.StringPtr(dataMap["connect_timeout_unit"].(string))
				}
				if dataMap["listener_mode"] != nil {
					listener.ListenerMode = common.StringPtr(dataMap["listener_mode"].(string))
				}
				if dataMap["listener_name"] != nil {
					listener.ListenerName = common.StringPtr(dataMap["listener_name"].(string))
				}
				if dataMap["listener_port"] != nil {
					listener.ListenerPort = common.IntPtr(dataMap["listener_port"].(int))
				}
				if dataMap["max_conn"] != nil {
					listener.MaxConn = common.IntPtr(dataMap["max_conn"].(int))
				}
				if dataMap["scheduler"] != nil {
					listener.Scheduler = common.StringPtr(dataMap["scheduler"].(string))
				}
				if dataMap["server_timeout"] != nil {
					listener.ServerTimeout = common.StringPtr(dataMap["server_timeout"].(string))
				}
				if dataMap["server_timeout_unit"] != nil {
					listener.ServerTimeoutUnit = common.StringPtr(dataMap["server_timeout_unit"].(string))
				}
				listeners = append(listeners, listener)
			}
			request.TcpListeners = listeners
		}
	}

	return request
}
