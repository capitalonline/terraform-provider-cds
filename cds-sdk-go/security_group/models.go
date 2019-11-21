package security_group

import (
	"encoding/json"
	cdshttp "terraform-provider-cds/cds-sdk-go/common/http"
)

type AddSecurityGroupRequest struct {
	*cdshttp.BaseRequest
	SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
	Description       *string `json:"Description" name:"Description"`
	SecurityGroupType *string `json:"SecurityGroupType" name:"SecurityGroupType"`
}

func (sg *AddSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *AddSecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type AddSecurityGroupResponse struct {
	*cdshttp.BaseResponse
	Code    *string                 `json:"Code" name:"Code"`
	Message *string                 `json:"Message" name:"Message"`
	Data    *map[string]interface{} `json:"Data" name:"Data"`
	TaskId  *string                 `json:"TaskId" name:"TaskId"`
}

func (sg *AddSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *AddSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type DescribeSecurityGroupRuleRequest struct {
	*cdshttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
	RuleId          *string `json:"RuleId,omitempty" name:"RuleId"`
}

func (sg *DescribeSecurityGroupRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *DescribeSecurityGroupRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type DescribeSecurityGroupRuleResponse struct {
	*cdshttp.BaseResponse
	Code    *string `json:"Code" name:"Code"`
	Message *string `json:"Message" name:"Message"`
	TaskId  *string `json:"TaskId" name:"TaskId"`
	Data    *struct {
		SecurityGroupType      *string               `json:"SecurityGroupType"`
		SecurityGroupRuleCount *int                  `json:"SecurityGroupRuleCount"`
		SecurityGroupRules     []*SecurityGroupRules `json:"SecurityGroupRules" name:"SecurityGroupRules" list`
	} `json:"Data,omitempty"`
}

type SecurityGroupRules struct {
	Status        *string `json:"Status"`
	TargetAddress *string `json:"TargetAddress"`
	Direction     *string `json:"Direction"`
	Protocol      *string `json:"Protocol"`
	LocalPort     *string `json:"LocalPort"`
	RuleId        *string `json:"RuleId"`
	Priority      *int    `json:"Priority"`
	Action        *int    `json:"Action"`
	TargetPort    *string `json:"TargetPort"`
	Type          *string `json:"Type"`
	CreateTime    *string `json:"CreateTime"`
	Description   *string `json:"Description"`
}

func (sg *DescribeSecurityGroupRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *DescribeSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type DeleteSecurityGroupRequest struct {
	*cdshttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId"`
}

func (sg *DeleteSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *DeleteSecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type DeleteSecurityGroupRespone struct {
	*cdshttp.BaseResponse
	Code    *string                 `json:"Code" name:"Code"`
	Message *string                 `json:"Message" name:"Message"`
	Data    *map[string]interface{} `json:"Data" name:"Data"`
	TaskId  *string                 `json:"TaskId" name:"TaskId"`
}

func (sg *DeleteSecurityGroupRespone) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *DeleteSecurityGroupRespone) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type DescribeSecurityGroupRequest struct {
	*cdshttp.BaseRequest
	SecurityGroupType *string `json:"SecurityGroupType,omitempty" name:"SecurityGroupType"`
	Keyword           *string `json:"keyword,omitempty" name:"Keyword"`
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (sg *DescribeSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *DescribeSecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type DescribeSecurityGroupResponse struct {
	*cdshttp.BaseResponse
	Message *string `json:"Message"`
	Code    *string `json:"Code"`
	Data    *struct {
		SecurityGroupCount *int             `json:"SecurityGroupCount"`
		SecurityGroup      []*SecurityGroup `json:"SecurityGroup"`
	} `json:"Data"`
	TaskID *string `json:"TaskId"`
}

type SecurityGroup struct {
	UpdateTime         *string `json:"UpdateTime"`
	Description        *string `json:"Description"`
	SecurityGroupName  *string `json:"SecurityGroupName"`
	SecurityGroupID    *string `json:"SecurityGroupId"`
	BindInstanceCount  *int    `json:"BindInstanceCount"`
	SecurityGroupType  *string `json:"SecurityGroupType"`
	BindVdcCount       *int    `json:"BindVdcCount"`
	CreateTime         *string `json:"CreateTime"`
	BindInterfaceCount *int    `json:"BindInterfaceCount"`
}

func (sg *DescribeSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *DescribeSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type BindData struct {
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	PrivateId  *string `json:"PrivateId,omitempty" name:"PrivateId"`
	PublicId   *string `json:"PublicId,omitempty" name:"PublicId"`
}
type JoinSecurityGroupRequest struct {
	*cdshttp.BaseRequest
	SecurityGroupId *string     `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	BindData        []*BindData `json:"BindData,omitempty" name:"BindData"`
}

func (sg *JoinSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *JoinSecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type JoinSecurityGroupResponse struct {
	*cdshttp.BaseResponse
	Code    *string                 `json:"Code" name:"Code"`
	Message *string                 `json:"Message" name:"Message"`
	Data    *map[string]interface{} `json:"Data" name:"Data"`
	TaskId  *string                 `json:"TaskId" name:"TaskId"`
}

func (sg *JoinSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *JoinSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type LeaveSecurityGroupRequest struct {
	*cdshttp.BaseRequest
	SecurityGroupId *string     `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	BindData        []*BindData `json:"BindData,omitempty" name:"BindData"`
}

func (sg *LeaveSecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *LeaveSecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type LeaveSecurityGroupResponse struct {
	*cdshttp.BaseResponse
	Code    *string                 `json:"Code" name:"Code"`
	Message *string                 `json:"Message" name:"Message"`
	Data    *map[string]interface{} `json:"Data" name:"Data"`
	TaskId  *string                 `json:"TaskId" name:"TaskId"`
}

func (sg *LeaveSecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *LeaveSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type ModifySecurityGroupRequest struct {
	*cdshttp.BaseRequest
	SecurityGroupName *string `json:"SecurityGroupName" name:"SecurityGroupName"`
	Description       *string `json:"Description" name:"Description"`
	SecurityGroupId   *string `json:"SecurityGroupId" name:"SecurityGroupId"`
}

func (sg *ModifySecurityGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *ModifySecurityGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type ModifySecurityGroupResponse struct {
	*cdshttp.BaseResponse
	Code    *string                 `json:"Code" name:"Code"`
	Message *string                 `json:"Message" name:"Message"`
	Data    *map[string]interface{} `json:"Data" name:"Data"`
	TaskId  *string                 `json:"TaskId" name:"TaskId"`
}

func (sg *ModifySecurityGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *ModifySecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}
