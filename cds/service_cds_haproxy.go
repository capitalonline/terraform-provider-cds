package cds

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"terraform-provider-cds/cds/connectivity"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/haproxy"
	"github.com/terraform-providers/terraform-provider-tencentcloud/tencentcloud/ratelimit"
)

type HaproxyService struct {
	client *connectivity.CdsClient
}

// Create Haproxy
func (me *HaproxyService) CreateHaproxy(ctx context.Context, request *haproxy.CreateLoadBalancerRequest) (*haproxy.CreateLoadBalancerResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyClient().CreateLoadBalancer(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

// Delete Haproxy
func (me *HaproxyService) DeleteHaproxy(ctx context.Context, request *haproxy.DeleteLoadBalancerRequest) (*haproxy.DeleteLoadBalancerResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyClient().DeleteLoadBalancer(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

// Modify Haproxy
func (me *HaproxyService) ModifyHaproxy(ctx context.Context, request *haproxy.ModifyLoadBalancerInstanceSpecRequest) (*haproxy.ModifyLoadBalancerInstanceSpecResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyClient().ModifyLoadBalancerInstanceSpec(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *HaproxyService) DescribeHaproxy(ctx context.Context, request *haproxy.DescribeLoadBalancersRequest) (*haproxy.DescribeLoadBalancersResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyGetClient().DescribeLoadBalancers(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *HaproxyService) DescribeLoadBalancersSpec(ctx context.Context, request *haproxy.DescribeLoadBalancersSpecRequest) (*haproxy.DescribeLoadBalancersSpecResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyGetClient().DescribeLoadBalancersSpec(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *HaproxyService) DescribeZones(ctx context.Context, request *haproxy.DescribeZonesRequest) (*haproxy.DescribeZonesResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyGetClient().DescribeZones(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *HaproxyService) DescribeLoadBalancersModifySpec(ctx context.Context, request *haproxy.DescribeLoadBalancersModifySpecRequest) (*haproxy.DescribeLoadBalancersModifySpecResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyGetClient().DescribeLoadBalancersModifySpec(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

// read Haproxy strategy
func (me *HaproxyService) DescribeLoadBalancerStrategys(ctx context.Context, request *haproxy.DescribeLoadBalancerStrategysRequest) (*haproxy.DescribeLoadBalancerStrategysResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyGetClient().DescribeLoadBalancerStrategys(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

// Modify Haproxy Strategy
func (me *HaproxyService) ModifyHaproxyStrategy(ctx context.Context, request *haproxy.ModifyLoadBalancerStrategysRequest) (*haproxy.ModifyLoadBalancerStrategysResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyClient().ModifyLoadBalancerStrategys(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

// Describe Certs
func (me *HaproxyService) DescribeCACertificates(ctx context.Context, request *haproxy.DescribeCACertificatesRequest) (*haproxy.DescribeCACertificatesResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyGetClient().DescribeCACertificates(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

// Describe Cert
func (me *HaproxyService) DescribeCACertificate(ctx context.Context, request *haproxy.DescribeCACertificateRequest) (*haproxy.DescribeCACertificateResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyGetClient().DescribeCACertificate(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

// Delete Cert
func (me *HaproxyService) DeleteCACertificate(ctx context.Context, request *haproxy.DeleteCACertificateRequest) (*haproxy.DeleteCACertificateResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyClient().DeleteCACertificate(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

// UploadCACertificate
func (me *HaproxyService) UploadCACertificate(ctx context.Context, request *haproxy.UploadCACertificateRequest) (*haproxy.UploadCACertificateResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyClient().UploadCACertificate(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

// Modify Haproxy name
func (me *HaproxyService) ModifyLoadBalancerName(ctx context.Context, request *haproxy.ModifyLoadBalancerNameRequest) (*haproxy.ModifyLoadBalancerNameResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseHaproxyClient().ModifyLoadBalancerName(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func flattenHaproxyInstanceVipsMapping(vips []*haproxy.DescribeLoadBalancersVips) []map[string]interface{} {
	listVip := make([]map[string]interface{}, 0, len(vips))
	for _, vip := range vips {
		data := map[string]interface{}{
			"ip":   vip.IP,
			"type": vip.Type,
		}
		listVip = append(listVip, data)
	}
	return listVip
}

func flattenHaproxyStrategyBackendServerMapping(list []*haproxy.DescribeLoadBalancerStrategysBackendServer) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, v := range list {
		data := map[string]interface{}{
			"ip":       *v.IP,
			"max_conn": *v.MaxConn,
			"port":     *v.Port,
			"weight":   *v.Weight,
		}
		result = append(result, data)
	}
	return result
}

func flattenHaproxyStrategyCertificateIdsMapping(list []*haproxy.DescribeLoadBalancerStrategysCertificateIds) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, v := range list {
		data := map[string]interface{}{
			"certificate_id":   *v.CertificateId,
			"certificate_name": *v.CertificateName,
		}
		result = append(result, data)
	}
	return result
}

func flattenHaproxyStrategyHttpMapping(httpListeners []*haproxy.DescribeLoadBalancerStrategysHttpListeners) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(httpListeners))
	for _, v := range httpListeners {
		data := map[string]interface{}{
			"acl_white_list":       v.AclWhiteList,
			"backend_server":       flattenHaproxyStrategyBackendServerMapping(v.BackendServer),
			"certificate_ids":      flattenHaproxyStrategyCertificateIdsMapping(v.CertificateIds),
			"client_timeout":       *v.ClientTimeout,
			"client_timeout_unit":  *v.ClientTimeoutUnit,
			"connect_timeout":      *v.ClientTimeout,
			"connect_timeout_unit": *v.ConnectTimeoutUnit,
			"listener_mode":        *v.ListenerMode,
			"listener_name":        *v.ListenerName,
			"listener_port":        *v.ListenerPort,
			"max_conn":             *v.MaxConn,
			"scheduler":            *v.Scheduler,
			"server_timeout":       *v.ServerTimeout,
			"server_timeout_unit":  *v.ServerTimeoutUnit,
			"sticky_session":       *v.StickySession,
			"session_persistence":  v.SessionPersistence,
			"option":               v.Option,
		}
		result = append(result, data)
	}

	return result
}

func flattenHaproxyStrategyTcpMapping(tcpListeners []*haproxy.DescribeLoadBalancerStrategysTcpListeners) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(tcpListeners))
	for _, v := range tcpListeners {
		data := map[string]interface{}{
			"acl_white_list":       v.AclWhiteList,
			"backend_server":       flattenHaproxyStrategyBackendServerMapping(v.BackendServer),
			"client_timeout":       *v.ClientTimeout,
			"client_timeout_unit":  *v.ClientTimeoutUnit,
			"connect_timeout":      *v.ClientTimeout,
			"connect_timeout_unit": *v.ConnectTimeoutUnit,
			"listener_mode":        *v.ListenerMode,
			"listener_name":        *v.ListenerName,
			"listener_port":        *v.ListenerPort,
			"max_conn":             *v.MaxConn,
			"scheduler":            *v.Scheduler,
			"server_timeout":       *v.ServerTimeout,
			"server_timeout_unit":  *v.ServerTimeoutUnit,
		}
		result = append(result, data)
	}

	return result
}
