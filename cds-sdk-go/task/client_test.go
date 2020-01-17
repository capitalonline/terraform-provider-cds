package task

import (
	"fmt"
	"testing"

	"terraform-provider-cds/cds-sdk-go/common"
	"terraform-provider-cds/cds-sdk-go/common/profile"
	"terraform-provider-cds/cds-sdk-go/common/regions"
)

func TestClient_DescribeTask(t *testing.T) {
	credential := common.NewCredential("", "")

	cpftask := profile.NewClientProfile()
	cpftask.HttpProfile.ReqMethod = "GET"
	taskclient, _ := NewClient(credential, regions.Beijing, cpftask)

	taskRequest := NewDescribeTaskRequest()
	taskRequest.TaskId = common.StringPtr("task id")
	taskResponse, err := taskclient.DescribeTask(taskRequest)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", taskResponse.ToJsonString(), err)
}
