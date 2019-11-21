package instance

import (
	"fmt"
	"terraform-provider-cds/cds-sdk-go/common"
	"terraform-provider-cds/cds-sdk-go/common/profile"
	"terraform-provider-cds/cds-sdk-go/common/regions"
	"terraform-provider-cds/cds-sdk-go/task"
	"testing"
)

func TestClient_CreateInstance(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewAddInstanceRequest()
	request.RegionId = common.StringPtr("CN_Beijing_A")
	request.VdcId = common.StringPtr("vdc id")
	request.Password = common.StringPtr("password")
	request.InstanceName = common.StringPtr("name")
	request.InstanceChargeType = common.StringPtr("PostPaid")
	request.AutoRenew = common.IntPtr(0)
	request.Cpu = common.IntPtr(1)
	request.Ram = common.IntPtr(1)
	request.PrepaidMonth = common.IntPtr(1)
	request.Amount = common.IntPtr(1)
	//request.AssignCCSId = common.StringPtr()
	request.ImageId = common.StringPtr("Ubuntu_16.04_64")
	request.PublicIp = common.StringPtrs([]string{"auto"})
	request.InstanceType = common.StringPtr("high_ccs") //7960400

	dd1 := DataDisk{
		Size: common.IntPtr(100),
		Type: common.StringPtr("high_disk"),
	}
	ip := PrivateIp{
		PrivateID: common.StringPtr(""),
		IP:        common.StringPtrs([]string{"auto"}),
	}
	request.DataDisks = []*DataDisk{&dd1}
	request.PrivateIp = []*PrivateIp{&ip}

	response, err := client.CreateInstance(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

	// get instance id
	taskRequest := task.NewDescribeTaskRequest()
	taskRequest.TaskId = response.TaskId
	cpftask := profile.NewClientProfile()
	cpftask.HttpProfile.ReqMethod = "GET"
	taskclient, _ := task.NewClient(credential, regions.Beijing, cpftask)
	taskResponse, err := taskclient.DescribeTask(taskRequest)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", taskResponse.ToJsonString(), err)
}

func TestClient_DescribeInstance(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDescribeInstanceRequest()
	//request.VdcId = common.StringPtr("")
	//request.PageNumber = common.IntPtr(1)
	//request.PageSize = common.IntPtr(1000)
	request.InstanceId = common.StringPtr("vdc id")
	response, err := client.DescribeInstance(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}

// wait for test
func TestClient_DeleteInstance(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDeleteInstanceRequest()
	request.InstanceIds = common.StringPtrs([]string{"instance id"})
	response, err := client.DeleteInstance(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}
