package security_group

import (
	"terraform-provider-cds/cds-sdk-go/common"
	cdshttp "terraform-provider-cds/cds-sdk-go/common/http"
	"terraform-provider-cds/cds-sdk-go/common/profile"
)

const ApiVersion = "2019-08-08"

type Client struct {
	common.Client
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewAddSecurityGroupRequest() (request *AddSecurityGroupRequest) {
	request = &AddSecurityGroupRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "CreateSecurityGroup")
	return
}

func NewAddSecurityGroupResponse() (response *AddSecurityGroupResponse) {
	response = &AddSecurityGroupResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDescribeSecurityGroupRuleRequest() (request *DescribeSecurityGroupRuleRequest) {
	request = &DescribeSecurityGroupRuleRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "DescribeSecurityGroupAttribute")
	return
}

func NewDescirbeSecurityGroupRuleResponse() (response *DescribeSecurityGroupRuleResponse) {
	response = &DescribeSecurityGroupRuleResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDescribeSecurityGroupRequest() (request *DescribeSecurityGroupRequest) {
	request = &DescribeSecurityGroupRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "DescribeSecurityGroups")
	return
}

func NewDescribeSecurityGroupResponse() (response *DescribeSecurityGroupResponse) {
	response = &DescribeSecurityGroupResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
	request = &DeleteSecurityGroupRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "DeleteSecurityGroup")
	return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupRespone) {
	response = &DeleteSecurityGroupRespone{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewJoinSecurityGroupRequest() (request *JoinSecurityGroupRequest) {
	request = &JoinSecurityGroupRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "JoinSecurityGroup")
	return
}

func NewJoinSecurityGroupResponse() (response *JoinSecurityGroupResponse) {
	response = &JoinSecurityGroupResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewLeaveSecurityGroupRequest() (request *LeaveSecurityGroupRequest) {
	request = &LeaveSecurityGroupRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "LeaveSecurityGroup")
	return
}

func NewLeaveSecurityGroupResponse() (response *LeaveSecurityGroupResponse) {
	response = &LeaveSecurityGroupResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewModifySecurityGroupRequest() (request *ModifySecurityGroupRequest) {
	request = &ModifySecurityGroupRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ModifySecurityGroupAttribute")
	return
}

func NewModifySecurityGroupResponse() (response *ModifySecurityGroupResponse) {
	response = &ModifySecurityGroupResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSecurityGroup(request *AddSecurityGroupRequest) (response *AddSecurityGroupResponse, err error) {
	if request == nil {
		request = NewAddSecurityGroupRequest()
	}
	response = NewAddSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) DescribeSecurityGroupAttribute(request *DescribeSecurityGroupRuleRequest) (response *DescribeSecurityGroupRuleResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupRuleRequest()
	}
	response = NewDescirbeSecurityGroupRuleResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (response *DeleteSecurityGroupRespone, err error) {
	if request == nil {
		request = NewDeleteSecurityGroupRequest()
	}
	response = NewDeleteSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) DescribeSecurityGroup(request *DescribeSecurityGroupRequest) (response *DescribeSecurityGroupResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupRequest()
	}
	response = NewDescribeSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) JoinSecurityGroup(request *JoinSecurityGroupRequest) (response *JoinSecurityGroupResponse, err error) {
	if request == nil {
		request = NewJoinSecurityGroupRequest()
	}
	response = NewJoinSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) LeaveSecurityGroup(request *LeaveSecurityGroupRequest) (response *LeaveSecurityGroupResponse, err error) {
	if request == nil {
		request = NewLeaveSecurityGroupRequest()
	}
	response = NewLeaveSecurityGroupResponse()
	err = c.Send(request, response)
	return
}
func (c *Client) ModifySecurityGroup(request *ModifySecurityGroupRequest) (response *ModifySecurityGroupResponse, err error) {
	if request == nil {
		request = NewModifySecurityGroupRequest()
	}
	response = NewModifySecurityGroupResponse()
	err = c.Send(request, response)
	return
}
