package vdc

import (
	"fmt"
	"terraform-provider-cds/cds-sdk-go/common"
	"terraform-provider-cds/cds-sdk-go/common/profile"
	"terraform-provider-cds/cds-sdk-go/common/regions"
	"terraform-provider-cds/cds-sdk-go/task"
	"testing"
)

func TestClient_CreateVdc(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	tp := NewAddVdcRequest()

	tp.RegionId = common.StringPtr(regions.Beijing)
	tp.VdcName = common.StringPtr("vdc name")
	tp.PublicNetwork = &PublicNetwork{
		IPNum: common.IntPtr(8),
		Qos:   common.IntPtr(10),
		Name:  common.StringPtr("name"),
		//FloatBandwidth: common.StringPtr("200"),
		BillingMethod: common.StringPtr("Bandwidth"),
		//AutoRenew:      common.StringPtr("1"),
		Type: common.StringPtr("Bandwidth_BGP"),
	}

	resp, err := client.CreateVdc(tp)

	if err != nil {
		fmt.Println("API request fail. >>>>>>> " + err.Error())
	} else {
		fmt.Printf(" Task: %v, code: %v", resp.TaskId, resp.Code)
	}

	// get instance id
	taskRequest := task.NewDescribeTaskRequest()
	taskRequest.TaskId = resp.TaskId
	cpftask := profile.NewClientProfile()
	cpftask.HttpProfile.ReqMethod = "GET"
	taskclient, _ := task.NewClient(credential, regions.Beijing, cpftask)
	taskResponse, err := taskclient.DescribeTask(taskRequest)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", taskResponse.ToJsonString(), err)

}

func TestClient_DescribeVdc(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := DescribeVdcRequest()
	request.VdcId = common.StringPtr("vdc id")
	//request.RegionId = common.StringPtr("")
	//request.PageNumber = common.IntPtr(1)
	//request.PageSize = common.IntPtr(30)
	//request.Keyword = common.StringPtr("")
	resp, err := client.DescribeVdc(request)
	if err != nil {
		fmt.Println("API request fail. >>>>>>> " + err.Error())
	} else {
		fmt.Printf(">>>>> Response Data: %s", resp.ToJsonString())
	}
}

func TestClient_DeletePublicNetwork(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDeletePublicNetworkRequest()
	request.PublicId = common.StringPtr("public id")
	resp, err := client.DeletePublicNetwork(request)
	if err != nil {
		fmt.Println("API request fail. >>>>>>> " + err.Error())
	} else {
		fmt.Printf(">>>>> Response Data: %s", resp.ToJsonString())
	}
}

func TestClient_DeleteVdc(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDeleteVdcRequest()
	request.VdcId = common.StringPtr("vdc id")
	resp, err := client.DeleteVdc(request)
	if err != nil {
		fmt.Println("API request fail. >>>>>>> " + err.Error())
	} else {
		fmt.Printf(">>>>> Response Data: %s", resp.ToJsonString())
	}
}

func TestClient_CreatePrivateNetwork(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewAddPrivateNetworkRequest()
	request.VdcId = common.StringPtr("vpc id")
	request.Name = common.StringPtr("name")
	request.Type = common.StringPtr("manual")
	request.Addres = common.StringPtr("192.168.0.0")
	request.Mask = common.IntPtr(16)
	resp, err := client.AddPrivateNetwork(request)
	if err != nil {
		fmt.Println("API request fail. >>>>>>> " + err.Error())
	} else {
		fmt.Printf(">>>>> Response Data: %s", resp.ToJsonString())
	}

}

func TestClient_DeletePrivateNetwork(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDeleteVdcRequest()
	request.VdcId = common.StringPtr("vdc id")
	resp, err := client.DeleteVdc(request)
	if err != nil {
		fmt.Println("API request fail. >>>>>>> " + err.Error())
	} else {
		fmt.Printf(">>>>> Response Data: %s", resp.ToJsonString())
	}
}
