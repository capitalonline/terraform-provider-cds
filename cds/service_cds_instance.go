package cds

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/instance"
	"terraform-provider-cds/cds/connectivity"

	"github.com/terraform-providers/terraform-provider-tencentcloud/tencentcloud/ratelimit"
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
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
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
