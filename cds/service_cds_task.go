package cds

import (
	"context"
	"github.com/terraform-providers/terraform-provider-tencentcloud/tencentcloud/ratelimit"
	"log"
	"terraform-provider-cds/cds-sdk-go/common"
	"terraform-provider-cds/cds-sdk-go/task"
	"terraform-provider-cds/cds/connectivity"
)

type TaskService struct {
	client *connectivity.CdsClient
}

// get task result
func (me *TaskService) DescribeTask(ctx context.Context, taskId string) (detail task.DescribeTaskResponse, errRet error) {

	logId := getLogId(ctx)
	request := task.NewDescribeTaskRequest()
	request.TaskId = common.StringPtr(taskId)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTaskGetClient().DescribeTask(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		detail = *response
		return
	}

	errRet = err
	return
}
