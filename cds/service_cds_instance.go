package cds

import (
	"context"
	"github.com/terraform-providers/terraform-provider-tencentcloud/tencentcloud/ratelimit"
	"log"
	"terraform-provider-cds/cds-sdk-go/instance"
	"terraform-provider-cds/cds/connectivity"
)

type InstanceService struct {
	client *connectivity.CdsClient
}

// Create Instance
func (me *InstanceService) CreateInstance(ctx context.Context, request *instance.AddInstanceRequest) (taskId string, errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().CreateInstance(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		taskId = *response.TaskId
		return
	}

	errRet = err
	return
}

func (me *InstanceService) DescribeInstance(ctx context.Context, request *instance.DescribeInstanceRequest) (response instance.DescribeInstanceReponse, errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	result, err := me.client.UseCvmClient().DescribeInstance(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		response = *result
		return
	}

	errRet = err
	return
}
