package cds

import (
	"context"
	"github.com/terraform-providers/terraform-provider-tencentcloud/tencentcloud/ratelimit"
	"log"
	"terraform-provider-cds/cds-sdk-go/vdc"
	"terraform-provider-cds/cds/connectivity"
	u "terraform-provider-cds/cds/utils"
)

type VdcService struct {
	client *connectivity.CdsClient
}

// ////////api
func (me *VdcService) CreateVdc(ctx context.Context, name string, region string,
	publicNetwork map[string]interface{}) (taskId string, errRet error) {

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
	response, err := me.client.UseVdcClient().AddPrivateNetwork(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		taskId = *response.TaskId
		return
	}

	errRet = err
	return
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
		taskId = *response.TaskId
		return
	}

	errRet = err
	return
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
		taskId = *response.TaskId
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
