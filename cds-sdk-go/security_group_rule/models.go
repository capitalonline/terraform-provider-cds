package security_group_rule

import "encoding/json"
import cdshttp "terraform-provider-cds/cds-sdk-go/common/http"

type RuleParam struct {
	Action        *string `json:"Action" name:"Action"`
	Description   *string `json:"Description" name:"Description"`
	TargetAddress *string `json:"TargetAddress" name:"TargetAddress"`
	TargetPort    *string `json:"TargetPort" name:"TargetPort"`
	LocalPort     *string `json:"LocalPort" name:"LocalPort"`
	Direction     *string `json:"Direction" name:"Direction"`
	Priority      *string `json:"Priority" name:"Priority"`
	Protocol      *string `json:"Protocol" name:"Protocol"`
	RuleType      *string `json:"RuleType" name:"RuleType"`
}

type AddSecurityGroupRuleRequest struct {
	*cdshttp.BaseRequest
	SecurityGroupId *string `json:"SecurityGroupId" name:"SecurityGroupId" tf:"securitygroupid"`
	Action          *int    `json:"Action" name:"Action"  tf:"action"`
	Description     *string `json:"Description" name:"Description"  tf:"description"`
	TargetAddress   *string `json:"TargetAddress" name:"TargetAddress"  tf:"targetaddress"`
	TargetPort      *string `json:"TargetPort" name:"TargetPort"  tf:"targetport"`
	LocalPort       *string `json:"LocalPort" name:"LocalPort" tf:"localport"`
	Direction       *string `json:"Direction" name:"Direction" tf:"direction"`
	Priority        *string `json:"Priority" name:"Priority" tf:"priority"`
	Protocol        *string `json:"Protocol" name:"Protocol" tf:"protocol"`
	RuleType        *string `json:"RuleType" name:"RuleType" tf:"ruletype"`
}

func (sg *AddSecurityGroupRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *AddSecurityGroupRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type AddSecurityGroupRuleResponse struct {
	*cdshttp.BaseResponse
	Code    *string                 `json:"Code" name:"Code"`
	Message *string                 `json:"Message" name:"Message"`
	Data    *map[string]interface{} `json:"Data" name:"Data"`
	TaskId  *string                 `json:"TaskId" name:"TaskId"`
}

func (sg *AddSecurityGroupRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

type DeleteSecurityGroupRuleRequest struct {
	*cdshttp.BaseRequest
	SecurityGroupId *string   `json:"SecurityGroupId" name:"SecurityGroupId"`
	RuleIds         []*string `json:"RuleIds" name:"RuleIds"`
}

func (sg *DeleteSecurityGroupRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}

func (sg *DeleteSecurityGroupRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &sg)
}

type DeleteSecurityGroupRuleResponse struct {
	*cdshttp.BaseResponse
	Code    *string                 `json:"Code" name:"Code"`
	Message *string                 `json:"Message" name:"Message"`
	Data    *map[string]interface{} `json:"Data" name:"Data"`
	TaskId  *string                 `json:"TaskId" name:"TaskId"`
}

func (sg *DeleteSecurityGroupRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(sg)
	return string(b)
}
