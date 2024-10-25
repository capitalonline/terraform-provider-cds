package cds

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"terraform-provider-cds/cds/connectivity"
	"terraform-provider-cds/cds/utils"
	u "terraform-provider-cds/cds/utils"

	"github.com/capitalonline/cds-gic-sdk-go/vdc"

	"github.com/terraform-providers/terraform-provider-tencentcloud/tencentcloud/ratelimit"
)

type VdcService struct {
	client *connectivity.CdsClient
}

// ////////api
func (me *VdcService) CreateVdc(ctx context.Context, name string, region string,
	publicNetwork map[string]interface{}, subjectId int) (taskId string, errRet error) {

	logId := getLogId(ctx)
	request := vdc.NewAddVdcRequest()
	request.RegionId = &region
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.VdcName = &name
	var pn = vdc.PublicNetwork{}
	err := u.Mapstructure(publicNetwork, &pn)
	if err != nil {
		return
	}
	request.PublicNetwork = &pn
	if subjectId != 0 {
		request.SubjectId = &subjectId
	}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVdcClient().CreateVdc(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		taskId = *response.TaskId
		return
	}

	errRet = err
	return
}

func (me *VdcService) CreatePrivateNetwork(ctx context.Context, request *vdc.AddPrivateNetworkRequest) (taskId string, errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	timeout := time.Now().Add(20 * time.Minute)
	for {
		if time.Now().After(timeout) {
			errRet = errors.New("create private network timeout")
		}
		response, err := me.client.UseVdcClient().AddPrivateNetwork(request)
		if err != nil {
			return "", err
		}
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		taskId = *response.TaskId
		if *response.Code == "TaskDuplicate" {
			time.Sleep(5)
			continue
		}
		return
	}
}

func (me *VdcService) DescribeVdc(ctx context.Context, request *vdc.DescVdcRequest) (result vdc.DescVdcResponse, errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	response, err := me.client.UseVdcGetClient().DescribeVdc(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		result = *response
		return
	}

	errRet = err
	return
}

func (me *VdcService) DeletePrivateNetwork(
	ctx context.Context,
	request *vdc.DeletePrivateNetworkRequest) (
	taskId string,
	errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVdcGetClient().DeletePrivateNetwork(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		taskId = *response.TaskId
		return
	}

	errRet = err
	return
}
func (me *VdcService) DeletePublicNetwork(
	ctx context.Context,
	request *vdc.DeletePublicNetworkRequest) (
	taskId string,
	errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVdcGetClient().DeletePublicNetwork(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		//taskId = *response.TaskId
		return
	}

	errRet = err
	return
}

func (me *VdcService) CreatePublicNetwork(ctx context.Context, request *vdc.CreatePublicNetworkRequest) (*vdc.CreatePublicNetworkResponse, error) {
	logId := getLogId(ctx)
	var err error
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	return me.client.UseVdcClient().CreatePublicNetwork(request)
}

func (me *VdcService) DeleteVdc(
	ctx context.Context,
	request *vdc.DeleteVdcRequest) (
	taskId string,
	errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVdcGetClient().DeleteVdc(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		taskId = *response.TaskId
		return
	}

	errRet = err
	return
}

func (me *VdcService) ModifyPublicNetwork(
	ctx context.Context,
	request *vdc.ModifyPublicNetworkRequest) (
	taskId string,
	errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVdcGetClient().ModifyPublicNetwork(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		if response.TaskId != nil {
			taskId = *response.TaskId
		}
		return
	}

	errRet = err
	return
}

func (me *VdcService) AddPublicNetworkIp(
	ctx context.Context,
	request *vdc.AddPublicIpRequest) (
	taskId string,
	errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVdcGetClient().AddPublicIpNetwork(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		taskId = *response.TaskId
		return
	}

	errRet = err
	return
}
func (me *VdcService) RenewPublicNetwork(
	ctx context.Context,
	request *vdc.RenewPublicNetworkRequest) (
	taskId string,
	errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVdcGetClient().RenewPublicNetwork(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		taskId = *response.TaskId
		return
	}

	errRet = err
	return
}

func (me *VdcService) AddPublicIp(ctx context.Context, request *vdc.AddPublicIpRequest) (*vdc.AddPublicIpResponse, error) {
	logId := getLogId(ctx)
	var err error
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	return me.client.UseVdcGetClient().AddPublicIpNetwork(request)
}

func (me *VdcService) DeletePublicIp(ctx context.Context, request *vdc.DeletePublicIpRequest) (*vdc.DeletePublicIpResponse, error) {
	logId := getLogId(ctx)
	var err error
	defer func() {
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	return me.client.UseVdcGetClient().DeletePublicIpNetwork(request)
}

func flattenPrivateNetworkMappings(list []*vdc.PrivateNetwork) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))

	for _, v := range list {
		privateNetwork := map[string]interface{}{
			"private_id":   utils.PString(v.PrivateId),
			"status":       utils.PString(v.Status),
			"name":         utils.PString(v.Name),
			"unuse_ip_num": utils.PInt(v.UnuseIpNum),
			"segments":     v.Segments,
		}
		result = append(result, privateNetwork)
	}
	return result
}

func flattenPublicNetworkMappings(list []*vdc.PublicNetworkInfo) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, v := range list {
		publicNetwork := map[string]interface{}{
			"public_id":    utils.PString(v.PublicId),
			"status":       utils.PString(v.Status),
			"qos":          utils.PInt(v.Qos),
			"name":         utils.PString(v.Name),
			"unuse_ip_num": utils.PInt(v.UnuseIpNum),
			"segments":     flattenPublicNetworkSegmentsMappings(*v.Segments),
		}
		result = append(result, publicNetwork)
	}

	return result
}

func flattenPublicNetworkSegmentsMappings(list []vdc.PublicSegment) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, v := range list {
		segment := map[string]interface{}{
			"mask":       utils.PInt(v.Mask),
			"gateway":    utils.PString(v.Gateway),
			"segment_id": utils.PString(v.SegmentId),
			"address":    utils.PString(v.Address),
		}
		result = append(result, segment)
	}

	return result
}
