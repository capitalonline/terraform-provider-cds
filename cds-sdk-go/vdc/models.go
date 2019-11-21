package vdc

import (
	"encoding/json"
	cdshttp "terraform-provider-cds/cds-sdk-go/common/http"
)

// Create VDC Request
type AddVdcRequest struct {
	*cdshttp.BaseRequest

	RegionId      *string        `json:"RegionId" name:"RegionId"`
	VdcName       *string        `json:"VdcName" name:"VdcName"`
	PublicNetwork *PublicNetwork `json:"PublicNetwork,omitempty" name:"PublicNetwork"`
}

func (r *AddVdcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVdcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublicNetwork struct {
	Name          *string `json:"Name,omitempty" name:"Name" tf:"name"`
	Type          *string `json:"Type,omitempty" name:"Type" tf:"type"`
	BillingMethod *string `json:"BillingMethod,omitempty" name:"BillingMethod" tf:"billingmethod"`
	//BandwIdth			*string `json:"BandwIdth,omitempty" name:"BandwIdth"`
	Qos            *int    `json:"Qos,omitempty" name:"VQos" tf:"qos"`
	IPNum          *int    `json:"IPNum,omitempty" name:"IPNum" tf:"ipnum"`
	AutoRenew      *string `json:"AutoRenew,omitempty" name:"AutoRenew" tf:"autorenew"`
	FloatBandwidth *string `json:"FloatBandwidth,omitempty" name:"FloatBandwidth" tf:"floatbandwidth"`
}

// Create VDC Reponse
type AddVdcResponse struct {
	*cdshttp.BaseResponse
	Message *string `json:"Message,omitempty" name:"Message"`
	Code    *string `json:"Code,omitempty" name:"Code"`
	TaskId  *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *AddVdcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVdcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

//Describe VDC Request
type DescVdcRequest struct {
	*cdshttp.BaseRequest
	VdcId      *string `json:"VdcId" name:"VdcId"`
	RegionId   *string `json:"RegionId" name:"RegionId"`
	Keyword    *string `json:"Keyword" name:"Keyword"`
	PageNumber *int    `json:"PageNumber" name:"PageNumber"`
	PageSize   *int    `json:"PageSize" name:"PageSize"`
}

func (r *DescVdcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescVdcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

//Describe VDC Reponse
type DescVdcResponse struct {
	*cdshttp.BaseResponse
	Message *string        `json:"Message"`
	Code    *string        `json:"Code"`
	Data    []*DescVdcData `json:"Data"`
}

func (r *DescVdcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescVdcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescVdcData struct {
	VdcId          *string              `json:"VdcId" name:"VdcId"`
	VdcName        *string              `json:"VdcName" name:"VdcName"`
	RegionId       *string              `json:"RegionId" name:"RegionId"`
	PrivateNetwork []*PrivateNetwork    `json:"PrivateNetwork" name:"PrivateNetwork"`
	PublicNetwork  []*PublicNetworkInfo `json:"PublicNetwork" name:"PublicNetwork"`
}

func (r *DescVdcData) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescVdcData) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrivateNetwork struct {
	PrivateId  *string   `json:"PrivateId" name:"PrivateId"`
	Status     *string   `json:"status" name:"Status"`
	Name       *string   `json:"name" name:"Name"`
	UnuseIpNum *int      `json:"unuse_ip_num" name:"UnuseIpNum"`
	Segments   *[]string `json:"segments" name:"Segments"`
}

func (r *PrivateNetwork) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PrivateNetwork) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublicNetworkInfo struct {
	PublicId   *string          `json:"PublicId" name:"PublicId"`
	Status     *string          `json:"status" name:"Status"`
	Qos        *int             `json:"Qos" name:"Qos"`
	Name       *string          `json:"name" name:"Name"`
	UnuseIpNum *int             `json:"unuse_ip_num" name:"UnuseIpNum"`
	Segments   *[]PublicSegment `json:"segments" name:"Segments"`
}

func (r *PublicNetworkInfo) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublicNetworkInfo) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublicSegment struct {
	Mask      *int    `json:"Mask" name:"Mask"`
	Gateway   *string `json:"Gateway" name:"Gateway"`
	SegmentId *string `json:"SegmentId" name:"SegmentId"`
	Address   *string `json:"Address" name:"Address"`
}

func (r *PublicSegment) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublicSegment) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Create Public Network Request
type AddPublicNetworkRequest struct {
	*cdshttp.BaseRequest
	VdcId          *string `json:"VdcId,omitempty"`
	Name           *string `json:"Name,omitempty" tf:"name"`
	Type           *string `json:"Type,omitempty" tf:"type"`
	BillingMethod  *string `json:"BillingMethod,omitempty" tf:billingmethod"`
	Qos            *int    `json:"Qos,omitempty" tf:"qos"`
	IPNum          *int    `json:"IPNum,omitempty" tf:"ipnum"`
	AutoRenew      *int    `json:"AutoRenew,omitempty" tf:"autorenew"`
	FloatBandwidth *int    `json:"FloatBandwidth,omitempty" tf:"floatbandwidth"`
}

// Create Public Network Reponse
type AddPublicNetworkResponse struct {
	*cdshttp.BaseResponse
	Code   *string `json:"Code"`
	TaskId *string `json:"TaskId"`
}

// Create Private Network Request
type AddPrivateNetworkRequest struct {
	*cdshttp.BaseRequest
	VdcId  *string `json:"VdcId,omitempty"`
	Name   *string `json:"Name,omitempty"`
	Type   *string `json:"Type,omitempty"`
	Addres *string `json:"Address,omitempty"`
	Mask   *int    `json:"Mask,omitempty"`
}

func (r *AddPrivateNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPrivateNetworkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Create Private Network Reponse
type AddPrivateNetworkResponse struct {
	*cdshttp.BaseResponse
	Code   *string `json:"Code"`
	TaskId *string `json:"TaskId"`
}

func (r *AddPrivateNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPrivateNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Delete Public Network Request
type DeletePublicNetworkRequest struct {
	*cdshttp.BaseRequest
	PublicId *string `json:"PublicId" name:"PublicId"`
}

func (r *DeletePublicNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePublicNetworkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Delete Public Network Response
type DeletePublicNetworkResponse struct {
	*cdshttp.BaseResponse
	Code   *string `json:"Code"`
	TaskId *string `json:"TaskId"`
}

func (r *DeletePublicNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePublicNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Delete Private Network Request
type DeletePrivateNetworkRequest struct {
	*cdshttp.BaseRequest
	PrivateId *string `json:"PrivateId" name:"PrivateId"`
}

func (r *DeletePrivateNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivateNetworkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Delete Private Network Response
type DeletePrivateNetworkResponse struct {
	*cdshttp.BaseResponse
	Code   *string `json:"Code"`
	TaskId *string `json:"TaskId"`
}

func (r *DeletePrivateNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePrivateNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Delete VDC Request
type DeleteVdcRequest struct {
	*cdshttp.BaseRequest
	VdcId *string `json:"VdcId" name:"VdcId"`
}

func (r *DeleteVdcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVdcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Delete VDC Reponse
type DeleteVdcResponse struct {
	*cdshttp.BaseResponse
	Code   *string `json:"Code"`
	TaskId *string `json:"TaskId"`
}

func (r *DeleteVdcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVdcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Modify Public Network Request
type ModifyPublicNetworkRequest struct {
	*cdshttp.BaseRequest
	PublicId *string `json:"PublicId" name:"PublicId"`
	Qos      *string `json:"Qos" name:"Qos"`
}

func (r *ModifyPublicNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPublicNetworkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPublicNetworkResponse struct {
	*cdshttp.BaseResponse
	Code   *string `json:"Code"`
	TaskId *string `json:"TaskId"`
}

func (r *ModifyPublicNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPublicNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Add Public Ip Network Request
type AddPublicIpRequest struct {
	*cdshttp.BaseRequest
	PublicId *string `json:"PublicId" name:"PublicId"`
	Number   *string `json:"Qos" name:"Qos"`
}

func (r *AddPublicIpRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPublicIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddPublicIpResponse struct {
	*cdshttp.BaseResponse
	Code   *string `json:"Code"`
	TaskId *string `json:"TaskId"`
}

func (r *AddPublicIpResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPublicIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Renew public Network Request
type RenewPublicNetworkRequest struct {
	*cdshttp.BaseRequest
	PublicId  *string `json:"PublicId" name:"PublicId"`
	AutoRenew *int    `json:"AutoRenew" name:"AutoRenew"`
}

func (r *RenewPublicNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewPublicNetworkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewPublicNetworkResponse struct {
	*cdshttp.BaseResponse
	Code   *string `json:"Code"`
	TaskId *string `json:"TaskId"`
}

func (r *RenewPublicNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewPublicNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
