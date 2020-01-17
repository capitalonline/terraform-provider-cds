package security_group_rule

import (
	"fmt"
	"testing"

	"terraform-provider-cds/cds-sdk-go/common"
	"terraform-provider-cds/cds-sdk-go/common/profile"
	"terraform-provider-cds/cds-sdk-go/common/regions"
)

func TestClient_CreateSecurityGroup(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewAddSecurityGroupRuleRequest()
	request.SecurityGroupId = common.StringPtr("id")
	request.Action = common.IntPtr(1)
	request.Description = common.StringPtr("desc")
	request.TargetAddress = common.StringPtr("120.78.170.188/28")
	request.TargetPort = common.StringPtr("70")
	request.LocalPort = common.StringPtr("800")
	request.Direction = common.StringPtr("all")
	request.Priority = common.StringPtr("11")
	request.Protocol = common.StringPtr("TCP")
	request.RuleType = common.StringPtr("ip")

	response, err := client.CreateSecurityGroupRule(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}

func TestClient_DeleteSecurityGroup(t *testing.T) {
	credential := common.NewCredential("", "")

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	client, _ := NewClient(credential, regions.Beijing, cpf)

	request := NewDeleteSecurityGroupRuleRequest()
	request.SecurityGroupId = common.StringPtr("id")
	request.RuleIds = common.StringPtrs([]string{"rule id"})

	response, err := client.DeleteSecurityGroupRule(request)
	fmt.Printf(">>>>> Resonponse: %s, err: %s", response.ToJsonString(), err)

}
