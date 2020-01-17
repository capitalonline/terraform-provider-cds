package task

import (
	"encoding/json"
	"errors"
	"time"

	"terraform-provider-cds/cds-sdk-go/common"
	cdshttp "terraform-provider-cds/cds-sdk-go/common/http"
	"terraform-provider-cds/cds-sdk-go/common/profile"
)

const ApiVersion = "2019-08-08"

type Client struct {
	common.Client
}

type DescribeTaskRequest struct {
	*cdshttp.BaseRequest
	TaskId *string `json:"TaskId" name:"TaskId"`
}

func (instance *DescribeTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DescribeTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type DescribeTaskResponse struct {
	*cdshttp.BaseResponse
	Code    *string `json:"Code" name:"Code"`
	Message *string `json:"Message" name:"Message"`
	TaskId  *string `json:"TaskId" name:"TaskId"`
	Data    *struct {
		Status      *string   `json:"Status"`
		ResourceID  *string   `json:"ResourceId"`
		ResourceIds []*string `json:"ResourceIds"`
		TaskType    *string   `json:"TaskType"`
	} `json:"Data"`
}

func (instance *DescribeTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}
func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
	request = &DescribeTaskRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", ApiVersion, "DescribeTask")
	return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
	response = &DescribeTaskResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

// Create Instance
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
	if request == nil {
		request = NewDescribeTaskRequest()
	}
	response = NewDescribeTaskResponse()

	for i := 0; i < 100; i++ {
		err = c.Send(request, response)
		if err != nil {
			return
		}
		switch *response.Data.Status {
		case "FINISH":
			return
		case "ERROR":
			err = errors.New("get task status error")
			return
		}
		time.Sleep(10 * time.Second)
	}
	err = errors.New("get task timeout")
	return
}
