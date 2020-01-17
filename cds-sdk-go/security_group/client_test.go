package security_group

import (
	"fmt"
	"testing"

	"terraform-provider-cds/cds-sdk-go/common"
	"terraform-provider-cds/cds-sdk-go/common/profile"
	"terraform-provider-cds/cds-sdk-go/common/regions"
	"terraform-provider-cds/cds-sdk-go/security_group_rule"
)

func TestClient_CreateSecurityGroup(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewAddSecurityGroupRequest()
	request.Description = common.StringPtr("Describe.")
	request.SecurityGroupName = common.StringPtr("name")
	request.SecurityGroupType = common.StringPtr("public")
	response, err := client.CreateSecurityGroup(request)
	fmt.Printf(">>>>> Response: %s, err: %s", response.ToJsonString(), err)

}

func TestClient_DescribeSecurityGroupAttribute(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDescribeSecurityGroupRuleRequest()
	//request.RuleId = common.StringPtr("")
	request.SecurityGroupId = common.StringPtr("id")
	response, err := client.DescribeSecurityGroupAttribute(request)
	fmt.Printf(">>>>> Response: %s, err: %s", response.ToJsonString(), err)

}

func TestClient_DescribeSecurityGroup(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDescribeSecurityGroupRequest()
	response, err := client.DescribeSecurityGroup(request)
	fmt.Printf(">>>>> Response: %s, err: %s", response.ToJsonString(), err)

}

func TestClient_DeleteSecurityGroup(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDeleteSecurityGroupRequest()
	request.SecurityGroupId = common.StringPtr("id")
	response, err := client.DeleteSecurityGroup(request)
	fmt.Printf(">>>>> Response: %s, err: %s", response.ToJsonString(), err)

}

func TestClient_JoinSecurityGroup(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewJoinSecurityGroupRequest()
	request.BindData = append(request.BindData, &BindData{InstanceId: common.StringPtr(""), PrivateId: common.StringPtr("")})
	request.SecurityGroupId = common.StringPtr("id")
	response, err := client.JoinSecurityGroup(request)
	fmt.Printf(">>>>> Response: %s, err: %s", response.ToJsonString(), err)

}

func TestClient_LeaveSecurityGroup(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewLeaveSecurityGroupRequest()
	request.SecurityGroupId = common.StringPtr("id")
	request.BindData = append(request.BindData, &BindData{InstanceId: common.StringPtr(""), PrivateId: common.StringPtr("")})
	response, err := client.LeaveSecurityGroup(request)
	fmt.Printf(">>>>> Response: %s, err: %s", response.ToJsonString(), err)

}

func TestClient_AddRuleSecurityGroup(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := security_group_rule.NewClient(credential, regions.Beijing, cpf)

	request := security_group_rule.NewAddSecurityGroupRuleRequest()
	request.SecurityGroupId = common.StringPtr("id")
	request.Action = common.IntPtr(1)
	request.Description = common.StringPtr("1")
	request.TargetPort = common.StringPtr("0")
	request.TargetAddress = common.StringPtr("120.78.170.188/28")
	request.LocalPort = common.StringPtr("80")
	request.Direction = common.StringPtr("all")
	request.Protocol = common.StringPtr("TCP")
	request.Priority = common.StringPtr("20")
	request.RuleType = common.StringPtr("ip")

	response, err := client.CreateSecurityGroupRule(request)
	fmt.Printf(">>>>> Response: %s, err: %s", response.ToJsonString(), err)

}
