package vdc

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

func NewAddVdcRequest() (request *AddVdcRequest) {
	request = &AddVdcRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("network", ApiVersion, "CreateVdc")
	return
}

func NewAddVdcResponse() (response *AddVdcResponse) {
	response = &AddVdcResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func DescribeVdcRequest() (request *DescVdcRequest) {
	request = &DescVdcRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("network", ApiVersion, "DescribeVdc")
	return
}

func DescribeVdcResponse() (response *DescVdcResponse) {
	response = &DescVdcResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewAddPublicNetworkRequest() (request *AddPublicNetworkRequest) {
	request = &AddPublicNetworkRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("network", ApiVersion, "CreatePublicNetwork")
	return
}

func NewAddPublicNetworkResponse() (response *AddPublicNetworkResponse) {
	response = &AddPublicNetworkResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewAddPrivateNetworkRequest() (request *AddPrivateNetworkRequest) {
	request = &AddPrivateNetworkRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("network", ApiVersion, "CreatePrivateNetwork")
	return
}

func NewAddPrivateNetworkResponse() (response *AddPrivateNetworkResponse) {
	response = &AddPrivateNetworkResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDeletePublicNetworkRequest() (request *DeletePublicNetworkRequest) {
	request = &DeletePublicNetworkRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("network", ApiVersion, "DeletePublicNetwork")
	return
}

func NewDeletePublicNetworkResponse() (response *DeletePublicNetworkResponse) {
	response = &DeletePublicNetworkResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDeletePrivateNetworkRequest() (request *DeletePrivateNetworkRequest) {
	request = &DeletePrivateNetworkRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("network", ApiVersion, "DeletePrivateNetwork")
	return
}

func NewDeletePrivateNetworkResponse() (response *DeletePrivateNetworkResponse) {
	response = &DeletePrivateNetworkResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDeleteVdcRequest() (request *DeleteVdcRequest) {
	request = &DeleteVdcRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("network", ApiVersion, "DeleteVdc")
	return
}

func NewDeleteVdcResponse() (response *DeleteVdcResponse) {
	response = &DeleteVdcResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewModifyPublicNetworkRequest() (request *ModifyPublicNetworkRequest) {
	request = &ModifyPublicNetworkRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("network", ApiVersion, "ModifyPublicNetwork")
	return
}

func NewModifyPublicNetworkResponse() (response *ModifyPublicNetworkResponse) {
	response = &ModifyPublicNetworkResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewAddPublicIpRequest() (request *AddPublicIpRequest) {
	request = &AddPublicIpRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("network", ApiVersion, "AddPublicIp")
	return
}

func NewAddPublicIpResponse() (response *AddPublicIpResponse) {
	response = &AddPublicIpResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewRenewPublicNetworkRequest() (request *RenewPublicNetworkRequest) {
	request = &RenewPublicNetworkRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("network", ApiVersion, "RenewPublicNetwork")
	return
}

func NewRenewPublicNetworkResponse() (response *RenewPublicNetworkResponse) {
	response = &RenewPublicNetworkResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func (c *Client) CreateVdc(request *AddVdcRequest) (response *AddVdcResponse, err error) {
	if request == nil {
		request = NewAddVdcRequest()
	}
	response = NewAddVdcResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) DescribeVdc(request *DescVdcRequest) (response *DescVdcResponse, err error) {
	if request == nil {
		request = DescribeVdcRequest()
	}
	response = DescribeVdcResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) AddPublicNetwork(request *AddPublicNetworkRequest) (response *AddPublicNetworkResponse, err error) {

	if request == nil {
		request = NewAddPublicNetworkRequest()
	}
	response = NewAddPublicNetworkResponse()
	err = c.Send(request, response)
	return

}
func (c *Client) AddPrivateNetwork(request *AddPrivateNetworkRequest) (response *AddPrivateNetworkResponse, err error) {
	if request == nil {
		request = NewAddPrivateNetworkRequest()
	}
	response = NewAddPrivateNetworkResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) DeletePublicNetwork(request *DeletePublicNetworkRequest) (response *DeletePublicNetworkResponse, err error) {
	if request == nil {
		request = NewDeletePublicNetworkRequest()
	}
	response = NewDeletePublicNetworkResponse()
	err = c.Send(request, response)
	return
}

func (c *Client) DeletePrivateNetwork(request *DeletePrivateNetworkRequest) (response *DeletePrivateNetworkResponse, err error) {
	if request == nil {
		request = NewDeletePrivateNetworkRequest()
	}
	response = NewDeletePrivateNetworkResponse()
	err = c.Send(request, response)
	return
}
func (c *Client) DeleteVdc(request *DeleteVdcRequest) (response *DeleteVdcResponse, err error) {
	if request == nil {
		request = NewDeleteVdcRequest()
	}
	response = NewDeleteVdcResponse()
	err = c.Send(request, response)
	return
}
func (c *Client) ModifyPublicNetwork(request *ModifyPublicNetworkRequest) (response *ModifyPublicNetworkResponse, err error) {
	if request == nil {
		request = NewModifyPublicNetworkRequest()
	}
	response = NewModifyPublicNetworkResponse()
	err = c.Send(request, response)
	return
}
func (c *Client) AddPublicIpNetwork(request *AddPublicIpRequest) (response *AddPublicIpResponse, err error) {
	if request == nil {
		request = NewAddPublicIpRequest()
	}
	response = NewAddPublicIpResponse()
	err = c.Send(request, response)
	return
}
func (c *Client) RenewPublicNetwork(request *RenewPublicNetworkRequest) (response *RenewPublicNetworkResponse, err error) {
	if request == nil {
		request = NewRenewPublicNetworkRequest()
	}
	response = NewRenewPublicNetworkResponse()
	err = c.Send(request, response)
	return
}
