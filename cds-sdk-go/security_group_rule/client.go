package security_group_rule

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

func NewAddSecurityGroupRuleRequest() (request *AddSecurityGroupRuleRequest) {
	request = &AddSecurityGroupRuleRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "AddSecurityGroupRule")
	return
}

func NewAddSecurityGroupRuleResponse() (response *AddSecurityGroupRuleResponse) {
	response = &AddSecurityGroupRuleResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDeleteSecurityGroupRuleRequest() (request *DeleteSecurityGroupRuleRequest) {
	request = &DeleteSecurityGroupRuleRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "RemoveSecurityGroupRule")
	return
}

func NewDeleteSecurityGroupRuleResponse() (response *DeleteSecurityGroupRuleResponse) {
	response = &DeleteSecurityGroupRuleResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateSecurityGroupRule(request *AddSecurityGroupRuleRequest) (response *AddSecurityGroupRuleResponse, err error) {
	if request == nil {
		request = NewAddSecurityGroupRuleRequest()
	}
	response = NewAddSecurityGroupRuleResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) DeleteSecurityGroupRule(request *DeleteSecurityGroupRuleRequest) (response *DeleteSecurityGroupRuleResponse, err error) {
	if request == nil {
		request = NewDeleteSecurityGroupRuleRequest()
	}
	response = NewDeleteSecurityGroupRuleResponse()
	err = c.Send(request, response)
	return
}
