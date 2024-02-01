package cds

import (
	"context"
	"fmt"
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

const success = "Success"

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

func (me *InstanceService) DescribeInstance(ctx context.Context, request *instance.DescribeInstanceRequest) (response instance.DescribeInstanceResponse, errRet error) {

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
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		response = *result
		return
	}

	errRet = err
	return
}

func (me *InstanceService) ResetInstancesPassword(ctx context.Context, request *instance.ResetInstancesPasswordRequest) (*instance.ResetInstancesPasswordResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	response, err := me.client.UseCvmClient().ResetInstancesPassword(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *InstanceService) ResetImage(ctx context.Context, request *instance.ResetImageRequest) (*instance.ResetImageResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	response, err := me.client.UseCvmClient().ResetImage(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *InstanceService) ModifyInstanceChargeType(ctx context.Context, request *instance.ModifyInstanceChargeTypeRequest) (*instance.ModifyInstanceChargeTypeResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	response, err := me.client.UseCvmClient().ModifyInstanceChargeType(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *InstanceService) StopInstance(ctx context.Context, request *instance.StopInstanceRequest) (*instance.StopInstanceResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	response, err := me.client.UseCvmGetClient().StopInstance(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *InstanceService) StartInstance(ctx context.Context, request *instance.StartInstanceRequest) (*instance.StartInstanceResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	response, err := me.client.UseCvmGetClient().StartInstance(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *InstanceService) RebootInstance(ctx context.Context, request *instance.RebootInstanceRequest) (*instance.RebootInstanceResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	response, err := me.client.UseCvmGetClient().RebootInstance(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *InstanceService) AllocateDedicatedHosts(ctx context.Context, request *instance.AllocateDedicatedHostsRequest) (*instance.AllocateDedicatedHostsResponse, error) {
	var (
		err      error
		response *instance.AllocateDedicatedHostsResponse
	)
	logId := getLogId(ctx)
	defer func() {
		respStr := ""
		if response != nil {
			respStr = response.ToJsonString()
		}
		log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s] err [%v]", logId, request.GetAction(), request.ToJsonString(), respStr, err))
	}()
	ratelimit.Check(request.GetAction())
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	response, err = me.client.UseCvmGetClient().AllocateDedicatedHosts(request)
	if err != nil || response == nil {
		return nil, fmt.Errorf("request err:%v ,response is:%v", err, response)
	}
	if *response.Code != success {
		return nil, fmt.Errorf("request failed with code:%v ,message:%v", response.Code, response.Message)
	}
	return response, nil
}

func (me *InstanceService) DescribeDedicatedHosts(ctx context.Context, request *instance.DescribeDedicatedHostsRequest) (*instance.DescribeDedicatedHostsResponse, error) {
	var (
		err      error
		response *instance.DescribeDedicatedHostsResponse
	)
	logId := getLogId(ctx)
	defer func() {
		respStr := ""
		if response != nil {
			respStr = response.ToJsonString()
		}
		log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s] err [%v]", logId, request.GetAction(), request.ToJsonString(), respStr, err))
	}()
	ratelimit.Check(request.GetAction())
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	response, err = me.client.UseCvmGetClient().DescribeDedicatedHosts(request)
	if err != nil || response == nil {
		return nil, fmt.Errorf("request err:%v ,response is:%v", err, response)
	}
	if *response.Code != success {
		return nil, fmt.Errorf("request failed with code:%v ,message:%v", response.Code, response.Message)
	}
	return response, err
}

func (me *InstanceService) DescribeDedicatedHostTypes(ctx context.Context, request *instance.DescribeDedicatedHostTypesRequest) (*instance.DescribeDedicatedHostTypesResponse, error) {
	var (
		err      error
		response *instance.DescribeDedicatedHostTypesResponse
	)
	logId := getLogId(ctx)
	defer func() {
		respStr := ""
		if response != nil {
			respStr = response.ToJsonString()
		}
		log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s] err [%v]", logId, request.GetAction(), request.ToJsonString(), respStr, err))
	}()
	ratelimit.Check(request.GetAction())
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	response, err = me.client.UseCvmGetClient().DescribeDedicatedHostTypes(request)
	if err != nil || response == nil {
		return nil, fmt.Errorf("request err:%v ,response is:%v", err, response)
	}
	if *response.Code != success {
		return nil, fmt.Errorf("request failed with code:%v ,message:%v", response.Code, response.Message)
	}
	return response, err
}
