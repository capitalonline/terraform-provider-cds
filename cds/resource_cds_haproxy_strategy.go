package cds

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/haproxy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCdsHaproxyStrategy() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCdsHaproxyStrategy,
		Read:   readResourceCdsHaproxyStrategy,
		Update: updateResourceCdsHaproxyStrategy,
		Delete: deleteResourceCdsHaproxyStrategy,
		Schema: map[string]*schema.Schema{
			"instance_uuid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Haproxy Instance uuid.",
			},
			"http_listeners": {
				Type:        schema.TypeList,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Optional:    true,
				Computed:    true,
				Description: "http listeners",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_white_list": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"backend_server": {
							Type:       schema.TypeList,
							Optional:   true,
							Computed:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_conn": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"weight": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"certificate_ids": {
							Type:       schema.TypeList,
							Optional:   true,
							Computed:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"certificate_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"certificate_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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
							Optional: true,
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
	}
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

func readResourceCdsHaproxyStrategy(data *schema.ResourceData, meta interface{}) error {
	log.Println("read haproxy strategy")
	defer logElapsed("resource.cds_haproxy_strategy.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	descriptionLoadBalancerStrategyRequest := haproxy.NewDescribeLoadBalancerStrategysRequest()

	response, err := haproxyService.DescribeLoadBalancerStrategys(ctx, descriptionLoadBalancerStrategyRequest)
	if err != nil {
		return err
	}

	// TODO 更新
	for _, entry := range response.Data.HttpListeners {
		log.Println(entry)
	}

	for _, entry := range response.Data.TcpListeners {
		log.Println(entry)
	}
	return nil
}

func createResourceCdsHaproxyStrategy(data *schema.ResourceData, meta interface{}) error {
	log.Println("create haproxy strategy")
	defer logElapsed("resource.cds_haproxy_strategy.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	request := haproxy.NewModifyLoadBalancerStrategysRequest()
	instanceUuid := data.Get("instance_uuid").(string)
	request.InstanceUuid = &instanceUuid

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
		httpListenerEntry := &haproxy.DescribeLoadBalancerStrategysHttpListeners{
			ServerTimeoutUnit:  &httpListener.ServerTimeoutUnit,
			ServerTimeout:      &httpListener.ServerTimeout,
			StickySession:      &httpListener.StickySession,
			AclWhiteList:       common.StringPtrs(strings.Split(strings.TrimSpace(httpListener.AclWhiteList), ",")),
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
		}

		request.HttpListeners = append(request.HttpListeners, httpListenerEntry)
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
		tcpListenerEntry := &haproxy.DescribeLoadBalancerStrategysTcpListeners{
			ServerTimeoutUnit:  &tcpListener.ServerTimeoutUnit,
			ServerTimeout:      &tcpListener.ServerTimeout,
			AclWhiteList:       common.StringPtrs(strings.Split(strings.TrimSpace(tcpListener.AclWhiteList), ",")),
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
		request.TcpListeners = append(request.TcpListeners, tcpListenerEntry)
	}

	response, err := haproxyService.ModifyHaproxyStrategy(ctx, request)
	if err != nil {
		return err
	}
	if *response.Code != "Success" {
		return fmt.Errorf("Haproxy modify haproxy strategy with error: %s", err)
	}
	data.SetId(time.Now().String())
	return nil
}

func updateResourceCdsHaproxyStrategy(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteResourceCdsHaproxyStrategy(data *schema.ResourceData, meta interface{}) error {
	return nil
}
