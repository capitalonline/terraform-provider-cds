package cds

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Regon id.",
			},
			"vdc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vdc id.",
			},
			"base_pipe_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Base pipe id.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance name.",
			},
			"cpu": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Instance cpu num.",
			},
			"ram": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Instance ram size.",
			},
			"ips": {
				Type:       schema.TypeList,
				ConfigMode: schema.SchemaConfigModeAttr,
				Required:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pipe_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Pipe type. Possible values: pipe_type、public.",
						},
						"pipe_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Pipe id. Internal network ID, or public network ID.",
						},
						"segment_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Segment id. The number of the occupied public network segment, which will occupy 3 public networks.",
						},
					},
				},
				Description: "The network used by HA. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#3createloadbalancer)",
			},
			"http_listeners": {
				Type:        schema.TypeList,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Optional:    true,
				Description: "Http listeners. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#3createloadbalancer)",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_white_list": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Acl white list. Set a whitelist, for example 192.168.12.1,192.168.1.1/20.",
						},
						"backend_server": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Ip. IP addresses of the backend servers (domain names are supported under the public network).",
									},
									"max_conn": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Max connection. Maximum number of connections for the backend servers.",
									},
									"port": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Port. Backend server port.",
									},
									"weight": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Weight. Backend server weight, with a range of 1-256.",
									},
								},
							},
							Description: "Backend server. Backend server configuration. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#%E4%BF%AE%E6%94%B9%E7%AD%96%E7%95%A5BackendServerObj)",
						},
						"certificate_ids": {
							Type:       schema.TypeList,
							Optional:   true,
							Computed:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"certificate_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Certificate id.",
									},
									"certificate_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Certificate name.",
									},
								},
							},
							Description: "[View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#%E4%BF%AE%E6%94%B9%E7%AD%96%E7%95%A5certificateIdsobj)",
						},
						"client_timeout": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Client timeout. Set the timeout period for client connections.",
						},
						"client_timeout_unit": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Client timeout unit. Set the timeout period for request connections in units [\"ms\", \"s\"].",
						},
						"connect_timeout": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Connect timeout. Set the timeout period for request connections.",
						},
						"connect_timeout_unit": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Connect timeout unit. Set the timeout unit for request connections [\"ms\", \"s\"].",
						},
						"listener_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Listener mode.",
						},
						"listener_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the listening strategy.The character length is 1~50, starting with a letter and ending with a letter or number. It consists of lowercase letters, numbers, or underscores. The name must be unique.",
						},
						"listener_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Listener port. Listening port for the strategy.The value range is 1-65535, with ports 22, 1080, 9100, and 9101 disabled.",
						},
						"max_conn": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Max connection. Maximum number of connections for the proxy end.",
						},
						"scheduler": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Scheduler. Scheduling algorithm [\"roundrobin\", \"leastconn\", \"static-rr\", \"source\"].",
						},
						"server_timeout": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Server timeout.",
						},
						"server_timeout_unit": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Server timeout unit. Set the server timeout unit [\"ms\", \"s\"].",
						},
						"sticky_session": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Sticky session. Enable keep-alive for long connections [\"on\", \"off\"].",
						},
						"session_persistence": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							ConfigMode:  schema.SchemaConfigModeAttr,
							Description: "Session persistence. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#%E4%BF%AE%E6%94%B9%E7%AD%96%E7%95%A5SessionPersistenceObj)",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Key. Set the Cookie key value in the HTTP response. The character length is 2-15, starting with a letter and ending with a letter or number. It consists of lowercase letters, numbers, or underscores, and the key must be unique.",
									},
									"mode": {
										Type:     schema.TypeInt,
										Optional: true,
										Description: `
Mode. The default value is 1.
0. Indicates that the content of the Cookie is retained in the cache (e.g., CDN).
1. Indicates that the content of the Cookie is not retained in the cache (e.g., CDN) and the cookie set by the load balancer is visible to the backend service.
2. Indicates that the content of the Cookie is not retained in the cache (e.g., CDN) and the cookie set by the load balancer is not visible to the backend service, i.e., in transparent mode, only the cookies set by the client are visible.
`,
									},
									"timer": {
										Type:        schema.TypeMap,
										Optional:    true,
										ConfigMode:  schema.SchemaConfigModeAttr,
										Description: "Timer.Set the session persistence time parameter, unit: seconds. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#%E4%BF%AE%E6%94%B9%E7%AD%96%E7%95%A5TimerObj)",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"max_idle": {
													Type:        schema.TypeInt,
													Required:    false,
													Description: "Max idle. The default value is 0, indicating that no setting is applied, meaning session persistence even during idle periods.Unit: seconds. Supported range: 0-7200.Set the maximum idle time for session persistence. When there are no new requests within this duration, the session persistence will end. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#%E4%BF%AE%E6%94%B9%E7%AD%96%E7%95%A5TimerObj)",
												},
												"max_life": {
													Type:        schema.TypeInt,
													Required:    false,
													Description: "Max life. The default value is 0, indicating no setting, allowing the session to persist indefinitely. Unit: seconds. Supported range: 0-7200.Set the maximum duration for session persistence. If the duration exceeds this limit, the session persistence will end. It is required that maxlife is greater than maxidle.[View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#%E4%BF%AE%E6%94%B9%E7%AD%96%E7%95%A5TimerObj)",
												},
											},
										},
									},
								},
							},
						},
						"option": {
							Type:        schema.TypeList,
							ConfigMode:  schema.SchemaConfigModeAttr,
							Optional:    true,
							Description: "Option. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#%E4%BF%AE%E6%94%B9%E7%AD%96%E7%95%A5OptionObj)",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"httpchk": {
										Type:       schema.TypeMap,
										Optional:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"method": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Method. Default value is \"GET\". Parameter range: [\"GET\", \"HEAD\", \"OPTIONS\"].",
												},
												"uri": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Uri. ",
												},
											},
										},
										Description: "Http check. Enable health checks for HTTP.",
									},
								},
							},
						},
					},
				},
			},
			"tcp_listeners": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Tcp listeners.[View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#%E4%BF%AE%E6%94%B9%E7%AD%96%E7%95%A5tcplistenersobj)",
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_white_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Acl white list. Set whitelist, for example: [\"129.12.12.1\", \"192.168.1.1/20\"].",
						},
						"backend_server": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Ip. Backend server IP address (support domain names in public network).",
									},
									"max_conn": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Max connection. Maximum number of connections for the backend server.",
									},
									"port": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Backend server port.",
									},
									"weight": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Backend server weight, with a range of 1-256.",
									},
								},
							},
							Description: "Backend server configuration. [View Document](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#%E4%BF%AE%E6%94%B9%E7%AD%96%E7%95%A5BackendServerObj)",
						},
						"client_timeout": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Client timeout. Set the client connection timeout duration.",
						},
						"client_timeout_unit": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Client timeout unit. Set the unit of client connection timeout duration to [\"ms\", \"s\"].",
						},
						"connect_timeout": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Connect timeout. Set the request connection timeout duration.",
						},
						"connect_timeout_unit": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Connect timeout unit. Set the unit of request connection timeout duration to [\"ms\", \"s\"].",
						},
						"listener_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Listener mode. ",
						},
						"listener_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Listener name. The name of the listening strategy.  Character length: 1-50.  Starts with a letter, ends with a letter or number. Consists of lowercase letters, numbers, or underscores. The name must be unique.",
						},
						"listener_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Listener port. Listening port for the strategy. Value range: 1-65535. Ports 22, 1080, 9100, and 9101 are disabled and not within the valid range.",
						},
						"max_conn": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Max connection. Maximum number of connections for the proxy end.",
						},
						"scheduler": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Scheduler. Scheduling algorithm options include: [\"roundrobin\", \"leastconn\", \"static-rr\",  \"source\"].",
						},
						"server_timeout": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Server timeout. Set the server-side timeout duration.",
						},
						"server_timeout_unit": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Server timeout unit. Set the unit of server-side timeout duration to [\"ms\", \"s\"].",
						},
					},
				},
			},
			"subject_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Subject ID.",
			},
		},
		Description: "Haproxy instance.\n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
resource cds_haproxy my_haproxy {
	region_id 		= "XXXXXXXX"

	vdc_id 		= "XXXXXXXX"
	base_pipe_id  = "XXXXXXXX"
	instance_name   = "ha_instance"
    # paas_goods_id 从data.json PaasGoodsId 获取
	cpu = 1
	ram = 2
	ips = [
		{
			# pipe_id is PrivateNetwork PrivateId from vdc info 
			pipe_id    = "xxx"
			pipe_type  = "private"
			segment_id = "xxx"
		},
		#This parameter is required if you want to create a public network(如创建公网，则需要)
		{
			# pipe_id is PublicNetwork PublicId from vdc info 
			pipe_id    = "xxx"
			pipe_type  = "public"
			segment_id = "xxx"
		}
	]
	http_listeners = [{
		server_timeout_unit = "s"
		server_timeout      = 1300
		sticky_session      = "on"
		acl_white_list      = "192.168.9.1"
		listener_mode       = "http"
		max_conn            = "2022"
		connect_timeout_unit = "s"
		scheduler           = "roundrobin"
		connect_timeout     = "1300"
		client_timeout      = "1022"
		listener_name       = "terraform"
		client_timeout_unit = "ms"
		listener_port       = 24354
		backend_server = [{
			ip       = "192.168.12.1"
			max_conn = 2022
			port     = 12314
			weight   = 255
		}]
		certificate_ids = []

#		The parameters option is a list,only one element at most
#		option = [{
#			httpchk = {
#				method = "GET"
#				uri = "/health"
#			}
#		}]
#		The parameters session_persistence is a list,only one element at most
#		session_persistence  = [
#			{
#				key = "test"
#				mode = 1
#				timer = {
#					max_idle=33
#					max_life=44
#				}
#			}
#		]
	}]
}
` +
			"\n```",
	}
}

func createResourceCdsHaproxy(data *schema.ResourceData, meta interface{}) error {
	log.Println("create haproxy")
	defer logElapsed("resource.cds_haproxy.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	request := haproxy.NewCreateLoadBalancerRequest()

	var cpu, ram int
	var region string

	if inter, ok := data.GetOk("region_id"); ok {
		regionId, exist := inter.(string)
		if exist {
			request.RegionId = common.StringPtr(regionId)
		}
		region = regionId
	}

	if inter, ok := data.GetOk("cpu"); ok {
		cpu, _ = inter.(int)
	}

	if inter, ok := data.GetOk("ram"); ok {
		ram, _ = inter.(int)
	}

	paasGoodsId, err := matchPassGoodsId(ctx, haproxyService, cpu, ram, region)
	if err != nil {
		return err
	}
	request.PaasGoodsId = common.IntPtr(paasGoodsId)

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

	if inter, ok := data.GetOk("subject_id"); ok {
		subjectId, exist := inter.(int)
		if exist {
			request.SubjectId = common.IntPtr(subjectId)
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

	resp, err := haproxyService.CreateHaproxy(ctx, request)
	if err != nil {
		return err
	}

	if len(resp.Data.InstancesUuid) == 0 {
		return errors.New("create haproxy failed, please check")
	}

	instanceUuid := resp.Data.InstancesUuid[0]

	data.SetId(instanceUuid)

	if err := waitHaproxyRunning(ctx, haproxyService, instanceUuid); err != nil {
		return err
	}

	// 创建策略
	strategyRequest := haproxy.NewModifyLoadBalancerStrategysRequest()
	strategyRequest.InstanceUuid = &instanceUuid

	httpListeners := make([]*HaproxyStrategyHttpListenerProviderInput, 0)
	bytesData, _ := json.Marshal(data.Get("http_listeners"))
	if err := json.Unmarshal(bytesData, &httpListeners); err != nil {
		return err
	}
	for _, httpListener := range httpListeners {
		backendServer := []*haproxy.DescribeLoadBalancerStrategysBackendServer{}
		for _, backendServerEntry := range httpListener.BackendServer {
			backendServer = append(backendServer, &haproxy.DescribeLoadBalancerStrategysBackendServer{
				IP:      &backendServerEntry.IP,
				MaxConn: &backendServerEntry.MaxConn,
				Port:    &backendServerEntry.Port,
				Weight:  &backendServerEntry.Weight,
			})
		}

		certificateIds := make([]*haproxy.DescribeLoadBalancerStrategysCertificateIds, 0)
		for _, certificate := range httpListener.CertificateIds {
			certificateIds = append(certificateIds, &haproxy.DescribeLoadBalancerStrategysCertificateIds{
				CertificateId:   &certificate.CertificateId,
				CertificateName: &certificate.CertificateName,
			})
		}

		acl := make([]*string, 0)
		if strings.TrimSpace(httpListener.AclWhiteList) != "" {
			acl = common.StringPtrs(strings.Split(strings.TrimSpace(httpListener.AclWhiteList), ","))
		}
		var sessionPersistence *haproxy.DescribeLoadBalancerStrategysSessionPersistence
		if httpListener.SessionPersistence != nil && len(httpListener.SessionPersistence) > 0 {
			inputParams := httpListener.SessionPersistence[0]
			sessionPersistence = new(haproxy.DescribeLoadBalancerStrategysSessionPersistence)
			sessionPersistence.Key = common.StringPtr(inputParams.Key)
			sessionPersistence.Mode = common.IntPtr(inputParams.Mode)
			timer := &haproxy.DescribeLoadBalancerStrategysTimer{
				MaxIdle: common.IntPtr(inputParams.Timer.MaxIdle),
				MaxLife: common.IntPtr(inputParams.Timer.MaxLife),
			}
			sessionPersistence.Timer = timer
		}
		var option *haproxy.DescribeLoadBalancerStrategysOption
		if httpListener.Option != nil && len(httpListener.Option) > 0 {
			inputParams := httpListener.Option[0]
			option = new(haproxy.DescribeLoadBalancerStrategysOption)
			method := inputParams.Httpchk.Method
			uri := inputParams.Httpchk.Uri
			httpchk := &haproxy.DescribeLoadBalancerStrategysOptionHttpchk{
				Method: common.StringPtr(method),
				Uri:    common.StringPtr(uri),
			}
			option.Httpchk = httpchk
		}
		httpListenerEntry := &haproxy.DescribeLoadBalancerStrategysHttpListeners{
			ServerTimeoutUnit:  &httpListener.ServerTimeoutUnit,
			ServerTimeout:      &httpListener.ServerTimeout,
			StickySession:      &httpListener.StickySession,
			AclWhiteList:       acl,
			ListenerMode:       &httpListener.ListenerMode,
			MaxConn:            &httpListener.MaxConn,
			ConnectTimeout:     &httpListener.ConnectTimeout,
			ConnectTimeoutUnit: &httpListener.ConnectTimeoutUnit,
			Scheduler:          &httpListener.Scheduler,
			ClientTimeout:      &httpListener.ClientTimeout,
			ClientTimeoutUnit:  &httpListener.ClientTimeoutUnit,
			ListenerName:       &httpListener.ListenerName,
			ListenerPort:       &httpListener.ListenerPort,
			BackendServer:      backendServer,
			CertificateIds:     certificateIds,
			SessionPersistence: sessionPersistence,
			Option:             option,
		}

		strategyRequest.HttpListeners = append(strategyRequest.HttpListeners, httpListenerEntry)
	}

	tcpListeners := make([]*HaproxyStrategyTcpListenerProviderInput, 0)
	bytesData, _ = json.Marshal(data.Get("tcp_listeners"))
	if err := json.Unmarshal(bytesData, &tcpListeners); err != nil {
		return err
	}

	for _, tcpListener := range tcpListeners {
		backendServer := []*haproxy.DescribeLoadBalancerStrategysBackendServer{}
		for _, backendServerEntry := range tcpListener.BackendServer {
			backendServer = append(backendServer, &haproxy.DescribeLoadBalancerStrategysBackendServer{
				IP:      &backendServerEntry.IP,
				MaxConn: &backendServerEntry.MaxConn,
				Port:    &backendServerEntry.Port,
				Weight:  &backendServerEntry.Weight,
			})
		}

		acl := make([]*string, 0)
		if strings.TrimSpace(tcpListener.AclWhiteList) != "" {
			acl = common.StringPtrs(strings.Split(strings.TrimSpace(tcpListener.AclWhiteList), ","))
		}

		tcpListenerEntry := &haproxy.DescribeLoadBalancerStrategysTcpListeners{
			ServerTimeoutUnit:  &tcpListener.ServerTimeoutUnit,
			ServerTimeout:      &tcpListener.ServerTimeout,
			AclWhiteList:       acl,
			ListenerMode:       &tcpListener.ListenerMode,
			MaxConn:            &tcpListener.MaxConn,
			ConnectTimeout:     &tcpListener.ConnectTimeout,
			ConnectTimeoutUnit: &tcpListener.ConnectTimeoutUnit,
			Scheduler:          &tcpListener.Scheduler,
			ClientTimeout:      &tcpListener.ClientTimeout,
			ClientTimeoutUnit:  &tcpListener.ClientTimeoutUnit,
			ListenerName:       &tcpListener.ListenerName,
			ListenerPort:       &tcpListener.ListenerPort,
			BackendServer:      backendServer,
		}
		strategyRequest.TcpListeners = append(strategyRequest.TcpListeners, tcpListenerEntry)
	}

	response, err := haproxyService.ModifyHaproxyStrategy(ctx, strategyRequest)
	if err != nil {
		return err
	}
	if *response.Code != "Success" {
		return fmt.Errorf("Haproxy modify haproxy strategy with error:" + err.Error())
	}

	if err := waitHaproxyRunning(ctx, haproxyService, instanceUuid); err != nil {
		return err
	}

	return readResourceCdsHaproxy(data, meta)
}

func readResourceCdsHaproxy(data *schema.ResourceData, meta interface{}) error {
	log.Println("read haproxy")
	defer logElapsed("resource.cds_haproxy.read")()

	return nil
}

func updateResourceCdsHaproxy(data *schema.ResourceData, meta interface{}) error {
	log.Println("create instance")
	defer logElapsed("resource.cds_instance.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	var cpu, ram int
	var hasChange bool

	if data.HasChange("cpu") {
		hasChange = true
		inter, _ := data.GetOk("cpu")
		cpu = inter.(int)
	}

	if data.HasChange("ram") {
		hasChange = true
		inter, _ := data.GetOk("ram")
		ram = inter.(int)
	}
	if data.HasChange("instance_name") {
		str, _ := data.GetOk("instance_name")
		name := str.(string)
		instanceUuid := data.Id()
		request := haproxy.NewModifyLoadBalancerNameRequest()
		request.InstanceUuid = common.StringPtr(instanceUuid)
		request.InstanceName = common.StringPtr(name)
		_, err := haproxyService.ModifyLoadBalancerName(context.Background(), request)
		if err != nil {
			return err
		}
	}

	inter, _ := data.GetOk("region_id")
	regionId := inter.(string)

	if hasChange {
		request := haproxy.NewModifyLoadBalancerInstanceSpecRequest()
		paasGoodsId, err := matchPassGoodsId(ctx, haproxyService, cpu, ram, regionId)
		if err != nil {
			return err
		}

		request.InstanceUuid = common.StringPtr(data.Id())
		request.PaasGoodsId = &paasGoodsId

		_, err = haproxyService.ModifyHaproxy(ctx, request)
		if err != nil {
			return err
		}

		if err := waitHaproxyRunning(ctx, haproxyService, data.Id()); err != nil {
			return err
		}
	}

	modifyResponse, err := haproxyService.ModifyHaproxyStrategy(ctx, createModitifyStrategyRequest(data))
	if err != nil {
		return err
	}

	if *modifyResponse.Code != "Success" {
		return errors.New(*modifyResponse.Message)
	}

	return readResourceCdsHaproxy(data, meta)
}

func deleteRresourceCdsHaproxy(data *schema.ResourceData, meta interface{}) error {
	log.Println("create instance")
	defer logElapsed("resource.cds_instance.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	request := haproxy.NewDeleteLoadBalancerRequest()
	request.InstanceUuid = common.StringPtr(data.Id())

	_, err := haproxyService.DeleteHaproxy(ctx, request)
	if err != nil {
		return err
	}

	if err := waitHaproxyDelete(ctx, haproxyService, data.Id()); err != nil {
		return err
	}
	return nil
}

func waitHaproxyRunning(ctx context.Context, service HaproxyService, instanceUuid string) error {
	descRequest := haproxy.NewDescribeLoadBalancersRequest()
	descRequest.InstanceUuid = &instanceUuid

	// 等待直到创建成功返回
	for {
		time.Sleep(time.Second * 15)
		descResp, err := service.DescribeHaproxy(ctx, descRequest)
		if err != nil {
			return err
		}

		if *descResp.Code != "Success" {
			return errors.New(*descResp.Message)
		}

		for _, data := range descResp.Data {
			if *data.Status == "RUNNING" {
				return nil
			}
		}
	}
}

func waitHaproxyDelete(ctx context.Context, service HaproxyService, instanceUuid string) error {
	descRequest := haproxy.NewDescribeLoadBalancersRequest()
	descRequest.InstanceUuid = &instanceUuid

	for {
		time.Sleep(time.Second * 15)
		response, err := service.DescribeHaproxy(ctx, descRequest)
		if err != nil {
			return nil
		}
		if len(response.Data) == 0 {
			return nil
		}
	}
}

func matchPassGoodsId(ctx context.Context, service HaproxyService, cpu, ram int, regionId string) (int, error) {
	goodsRequest := haproxy.NewDescribeLoadBalancersSpecRequest()

	goodsRequest.RegionId = common.StringPtr(regionId)

	goodsResponse, err := service.DescribeLoadBalancersSpec(ctx, goodsRequest)
	if err != nil {
		return -1, err
	}

	for _, product := range goodsResponse.Data.Products {
		for _, arch := range product.Architectures {
			for _, role := range arch.ComputeRoles {
				for _, cpuRam := range role.Standards.CpuRam {
					if *cpuRam.CPU == cpu && *cpuRam.RAM == ram {
						return *cpuRam.PaasGoodsId, nil
					}
				}
			}
		}
	}

	return -1, fmt.Errorf("cpu %d, ram %d not found paas_goods_id", cpu, ram)
}

func createModitifyStrategyRequest(data *schema.ResourceData) *haproxy.ModifyLoadBalancerStrategysRequest {
	request := haproxy.NewModifyLoadBalancerStrategysRequest()

	request.InstanceUuid = common.StringPtr(data.Id())

	if inter, ok := data.GetOk("http_listeners"); ok {
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
			} else {
				certificateIds = []*haproxy.DescribeLoadBalancerStrategysCertificateIds{}
			}

			listener := &haproxy.DescribeLoadBalancerStrategysHttpListeners{
				BackendServer:  backendServers,
				CertificateIds: certificateIds,
			}

			if dataMap["acl_white_list"] != nil && dataMap["acl_white_list"].(string) != "" {
				listener.AclWhiteList = common.StringPtrs(strings.Split(dataMap["acl_white_list"].(string), ","))
			} else {
				acl := make([]*string, 0)
				listener.AclWhiteList = acl
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
			if dataMap["session_persistence"] != nil {
				sessionPersistenceList := dataMap["session_persistence"].([]interface{})
				if sessionPersistenceList != nil && len(sessionPersistenceList) > 0 {
					paramsMap := sessionPersistenceList[0].(map[string]interface{})
					sessionPersistence := new(haproxy.DescribeLoadBalancerStrategysSessionPersistence)
					if paramsMap["key"] != nil {
						key := paramsMap["key"].(string)
						sessionPersistence.Key = common.StringPtr(key)
					}
					if paramsMap["mode"] != nil {
						mode := paramsMap["mode"].(int)
						sessionPersistence.Mode = common.IntPtr(mode)
					}
					if paramsMap["timer"] != nil {
						timerMap := paramsMap["timer"].(map[string]interface{})
						if timerMap["max_idle"] != nil && timerMap["max_life"] != nil {
							maxIdle, _ := strconv.Atoi(timerMap["max_idle"].(string))
							maxLife, _ := strconv.Atoi(timerMap["max_life"].(string))
							sessionPersistence.Timer = &haproxy.DescribeLoadBalancerStrategysTimer{
								MaxIdle: common.IntPtr(maxIdle),
								MaxLife: common.IntPtr(maxLife),
							}
						}
					}
					listener.SessionPersistence = sessionPersistence
				}
			}
			if dataMap["option"] != nil {
				optionList := dataMap["option"].([]interface{})
				if optionList != nil && len(optionList) > 0 {
					optionMap := optionList[0].(map[string]interface{})
					if optionMap["httpchk"] != nil {
						httpchkMap := optionMap["httpchk"].(map[string]interface{})
						if httpchkMap["method"] != nil && httpchkMap["uri"] != nil {
							method := httpchkMap["method"].(string)
							uri := httpchkMap["uri"].(string)
							listener.Option = &haproxy.DescribeLoadBalancerStrategysOption{
								Httpchk: &haproxy.DescribeLoadBalancerStrategysOptionHttpchk{
									Method: common.StringPtr(method),
									Uri:    common.StringPtr(uri),
								},
							}
						}
					}
				}
			}
			listeners = append(listeners, listener)
		}
		request.HttpListeners = listeners
	}

	if inter, ok := data.GetOk("tcp_listeners"); ok {
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

			if dataMap["acl_white_list"] != nil && dataMap["acl_white_list"].(string) != "" {
				listener.AclWhiteList = common.StringPtrs(strings.Split(dataMap["acl_white_list"].(string), ","))
			} else {
				acl := make([]*string, 0)
				listener.AclWhiteList = acl
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

	return request
}

type HaproxyStrategyHttpListenerProviderInput struct {
	ServerTimeoutUnit  string `json:"server_timeout_unit"`
	ServerTimeout      string `json:"server_timeout"`
	StickySession      string `json:"sticky_session"`
	AclWhiteList       string `json:"acl_white_list"`
	ListenerMode       string `json:"listener_mode"`
	MaxConn            int    `json:"max_conn"`
	Scheduler          string `json:"scheduler"`
	ConnectTimeout     string `json:"connect_timeout"`
	ConnectTimeoutUnit string `json:"connect_timeout_unit"`
	ClientTimeout      string `json:"client_timeout"`
	ClientTimeoutUnit  string `json:"client_timeout_unit"`
	ListenerPort       int    `json:"listener_port"`
	ListenerName       string `json:"listener_name"`
	BackendServer      []*struct {
		IP      string `json:"ip"`
		MaxConn int    `json:"max_conn"`
		Port    int    `json:"port"`
		Weight  string `json:"weight"`
	} `json:"backend_server"`
	CertificateIds []*struct {
		CertificateId   string `json:"certificate_id"`
		CertificateName string `json:"certificate_name"`
	} `json:"certificate_ids"`
	SessionPersistence []struct {
		Key   string `json:"Key"`
		Mode  int    `json:"Mode"`
		Timer struct {
			MaxIdle int `json:"MaxIdle"`
			MaxLife int `json:"MaxLife"`
		} `json:"Timer"`
	} `json:"SessionPersistence"`
	Option []struct {
		Httpchk struct {
			Method string `json:"Method"`
			Uri    string `json:"Uri"`
		} `json:"Httpchk"`
	} `json:"Option"`
}

type HaproxyStrategyTcpListenerProviderInput struct {
	ServerTimeoutUnit  string `json:"server_timeout_unit"`
	ServerTimeout      string `json:"server_timeout"`
	AclWhiteList       string `json:"acl_white_list"`
	ListenerMode       string `json:"listener_mode"`
	MaxConn            int    `json:"max_conn"`
	Scheduler          string `json:"scheduler"`
	ConnectTimeout     string `json:"connect_timeout"`
	ConnectTimeoutUnit string `json:"connect_timeout_unit"`
	ClientTimeout      string `json:"client_timeout"`
	ClientTimeoutUnit  string `json:"client_timeout_unit"`
	ListenerPort       int    `json:"listener_port"`
	ListenerName       string `json:"listener_name"`
	BackendServer      []*struct {
		IP      string `json:"ip"`
		MaxConn int    `json:"max_conn"`
		Port    int    `json:"port"`
		Weight  string `json:"weight"`
	} `json:"backend_server"`
}
