package instance

import (
	"encoding/json"
	cdshttp "terraform-provider-cds/cds-sdk-go/common/http"
)

// Create Instance Request
type AddInstanceRequest struct {
	*cdshttp.BaseRequest
	RegionId           *string      `json:"RegionId,omitempty" name:"RegionId"`
	VdcId              *string      `json:"VdcId,omitempty" name:"VdcId"`
	InstanceName       *string      `json:"InstanceName,omitempty" name:"InstanceName"`
	InstanceChargeType *string      `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	Password           *string      `json:"Password,omitempty" name:"Password"`
	Cpu                *int         `json:"Cpu,omitempty" name:"Cpu"`
	Ram                *int         `json:"Ram,omitempty" name:"Ram"`
	InstanceType       *string      `json:"InstanceType,omitempty" name:"InstanceType"`
	ImageId            *string      `json:"ImageId,omitempty" name:"ImageId"`
	AssignCCSId        *string      `json:"AssignCCSId,omitempty" name:"AssignCCSId"`
	DataDisks          []*DataDisk  `json:"DataDisks,omitempty" name:"DataDisks"`
	AutoRenew          *int         `json:"AutoRenew,omitempty" name:"AutoRenew"`
	PrepaidMonth       *int         `json:"PrepaidMonth,omitempty" name:"PrepaidMonth"`
	PublicIp           []*string    `json:"PublicIp,omitempty" name:"PublicIp"`
	PrivateIp          []*PrivateIp `json:"PrivateIp,omitempty" name:"PrivateIp"`
	Amount             *int         `json:"Amount,omitempty" name:"Amount"`
}

type DataDisk struct {
	Size *int    `json:"Size,omitempty" name:"Size"`
	Type *string `json:"Type,omitempty" name:"Type"`
}

type PrivateIp struct {
	PrivateID *string   `json:"PrivateId,omitempty" name:"PrivateId"`
	IP        []*string `json:"IP,omitempty" name:"IP"`
}

func (instance *AddInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *AddInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Create Instance Reponse
type AddInstanceResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *AddInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *AddInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

//Describe Instance Request
type DescribeInstanceRequest struct {
	*cdshttp.BaseRequest
	VdcId      *string `json:"VdcId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty"`
	PublicIp   *string `json:"PublicIp,omitempty"`
	PageNumber *int    `json:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty"`
}

func (instance *DescribeInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DescribeInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type DescribeReturnInfo struct {
	InstanceStatus          *string    `json:"InstanceStatus"`
	InstanceName            *string    `json:"InstanceName"`
	InstanceId              *string    `json:"InstanceId"`
	VdcId                   *string    `json:"VdcId"`
	VdcName                 *string    `json:"VdcName"`
	Disks                   *DisksInfo `json:"Disks"`
	PrivateNetworkInterface []*PrivateNetworkInterfaceInfo
	PublicNetworkInterface  *PublicNetworkInterfaceInfo `json:"PublicNetworkInterface"`
	InstanceChargeType      *string                     `json:"InstanceChargeType"`
}
type DisksInfo struct {
	DataDisks       []*DataDisksInfo `json:"DataDisks"`
	SystemDisk      *map[string]int  `json:"SystemDisk"`
	LeftDataDiskNum *int             `json:"LeftDataDiskNum"`
}

type DataDisksInfo struct {
	DiskId   *string `json:"DiskId"`
	DiskType *string `json:"DiskType"`
	Size     *int    `json:"Size"`
	IopsSize *int    `json:"IopsSize"`
}
type PrivateNetworkInterfaceInfo struct {
	InterfaceId *string `json:"InterfaceId"`
	Name        *string `json:"Name"`
	IP          *string `json:"IP"`
	MAC         *string `json:"MAC"`
	Connected   *int    `json:"Connected"`
	PrivateId   *string `json:"PrivateId"`
}

type PublicNetworkInterfaceInfo struct {
	InterfaceId *string `json:"InterfaceId"`
	Name        *string `json:"Name"`
	IP          *string `json:"IP"`
	MAC         *string `json:"MAC"`
	Connected   *int    `json:"Connected"`
	PublicId    *string `json:"PublicId"`
}
type DescribeInstanceReponse struct {
	*cdshttp.BaseResponse
	Code    *string `json:"Code"`
	Message *string `json:"Message"`
	Data    struct {
		Instances  []*DescribeReturnInfo
		PageNumber *int `json:"PageNumber"`
		PageCount  *int `json:"PageCount"`
	}
}

func (instance *DescribeInstanceReponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DescribeInstanceReponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Delete Instance Request
type DeleteInstanceRequest struct {
	*cdshttp.BaseRequest
	InstanceIds []*string
}

// Delete Instance Reponse
type DeleteInstanceResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *DeleteInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Create Disk
type CreateDiskRequest struct {
	*cdshttp.BaseRequest
	InstanceId *string     `json:"InstanceId,omitempty" name:"InstanceId"`
	DataDisks  []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

func (instance *CreateDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *CreateDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type CreateDiskResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *CreateDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *CreateDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Resize Disk
type ResizeDiskRequest struct {
	*cdshttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DiskId     *string `json:"DiskId,omitempty" name:"DiskId"`
	DataSize   *int    `json:"DataSize,omitempty" name:"DataSize"`
}

func (instance *ResizeDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ResizeDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type ResizeDiskResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *ResizeDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Resize Disk
type DeleteDiskRequest struct {
	*cdshttp.BaseRequest
	InstanceId *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	DiskIds    []*string `json:"DiskIds,omitempty" name:"DiskIds"`
}

func (instance *DeleteDiskRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DeleteDiskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type DeleteDiskResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *DeleteDiskResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *DeleteDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

// Modify Ip Address
type ModifyIpRequest struct {
	*cdshttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	InterfaceId *string `json:"InterfaceId,omitempty" name:"InterfaceId"`
	Address     *string `json:"Address,omitempty" name:"Address"`
}

func (instance *ModifyIpRequest) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ModifyIpRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}

type ModifyIpResponse struct {
	*cdshttp.BaseResponse
	Code   *string
	TaskId *string
}

func (instance *ModifyIpResponse) ToJsonString() string {
	b, _ := json.Marshal(instance)
	return string(b)
}

func (instance *ModifyIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &instance)
}
