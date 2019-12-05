package cds

import (
	"context"
	"github.com/terraform-providers/terraform-provider-tencentcloud/tencentcloud/ratelimit"
	"log"
	"terraform-provider-cds/cds-sdk-go/common"
	"terraform-provider-cds/cds-sdk-go/security_group"
	"terraform-provider-cds/cds-sdk-go/security_group_rule"
	"terraform-provider-cds/cds/connectivity"
	u "terraform-provider-cds/cds/utils"
	"time"
)

type SecurityGroupService struct {
	client *connectivity.CdsClient
}

// ////////api
func (me *SecurityGroupService) CreateSecurityGroup(
	ctx context.Context,
	name string,
	description string,
	securityGroupType string,
	rules []map[string]interface{}) (id string, errRet error) {

	logId := getLogId(ctx)
	request := security_group.NewAddSecurityGroupRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	// add security group
	request.SecurityGroupName = &name
	request.SecurityGroupType = &securityGroupType
	request.Description = &description
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseSecurityGroupClient().CreateSecurityGroup(request)
	if err != nil {
		errRet = err
		return
	}
	// get security group id
	time.Sleep(1 * time.Second)
	securityGroupId := common.StringPtr("")
	getIdRequest := security_group.NewDescribeSecurityGroupRequest()
	getIdRequest.Keyword = common.StringPtr(name)
	getIdRequest.SecurityGroupType = common.StringPtr(securityGroupType)
	groupResponse, getIdErr := me.client.UseSecurityGroupClient().DescribeSecurityGroup(getIdRequest)
	if getIdErr != nil {
		return "", getIdErr
	}
	for _, value := range groupResponse.Data.SecurityGroup {
		if *value.SecurityGroupName == name {
			securityGroupId = value.SecurityGroupID
		}
	}

	// add rule
	if len(rules) > 0 {
		for _, value := range rules {
			var pn = security_group_rule.NewAddSecurityGroupRuleRequest()
			pn.SecurityGroupId = securityGroupId
			decodeErr := u.Mapstructure(value, pn)
			if decodeErr != nil {
				errRet = decodeErr
				return
			}
			ruleResponse, MakeRuleErr := me.client.UseSecurityRuleClient().CreateSecurityGroupRule(pn)
			if MakeRuleErr != nil {
				log.Printf("[ERROR]%s api[%s] , request body [%s], response body[%s]\n",
					logId, pn.GetAction(), pn.ToJsonString(), ruleResponse.ToJsonString())
				errRet = MakeRuleErr
				return
			}
		}

	}

	log.Printf("[INFO]%s api[%s] , request body [%s], response body[%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	id = *securityGroupId
	return
}

func (me *SecurityGroupService) JoinSecurityGroup(
	ctx context.Context,
	request *security_group.JoinSecurityGroupRequest) (taskId string, errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	response, errRet := me.client.UseSecurityGroupClient().JoinSecurityGroup(request)
	if errRet != nil {
		return "", errRet
	}
	taskId = *response.TaskId
	return
}
func (me *SecurityGroupService) LeaveSecurityGroup(
	ctx context.Context,
	request *security_group.LeaveSecurityGroupRequest) (taskId string, errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	response, errRet := me.client.UseSecurityGroupClient().LeaveSecurityGroup(request)
	if errRet != nil {
		return
	}
	taskId = *response.TaskId
	return
}
func (me *SecurityGroupService) DeleteSecurityGroup(
	ctx context.Context,
	request *security_group.DeleteSecurityGroupRequest) (taskId string, errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	response, errRet := me.client.UseSecurityGroupGetClient().DeleteSecurityGroup(request)
	if errRet != nil {
		return "", errRet
	}
	taskId = *response.TaskId
	return
}

func (me *SecurityGroupService) DescribeSecurityGroup(
	ctx context.Context,
	request *security_group.DescribeSecurityGroupRequest) (response security_group.DescribeSecurityGroupResponse, errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	responseAd, errRet := me.client.UseSecurityGroupClient().DescribeSecurityGroup(request)
	response = *responseAd
	return
}

func (me *SecurityGroupService) ModifySecurityGroup(
	ctx context.Context,
	request *security_group.ModifySecurityGroupRequest) (taskId string, errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	reponse, errRet := me.client.UseSecurityGroupClient().ModifySecurityGroup(request)
	taskId = *reponse.TaskId
	return
}

func (me *SecurityGroupService) DescribeSecurityGroupRule(
	ctx context.Context,
	request *security_group.DescribeSecurityGroupRuleRequest) (response security_group.DescribeSecurityGroupRuleResponse, errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	result, errRet := me.client.UseSecurityGroupGetClient().DescribeSecurityGroupAttribute(request)
	response = *result
	return
}

func (me *SecurityGroupService) DeleteSecurityGroupRule(
	ctx context.Context,
	request *security_group_rule.DeleteSecurityGroupRuleRequest) (taskId string, errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	_, errRet = me.client.UseSecurityRuleClient().DeleteSecurityGroupRule(request)
	taskId = ""
	//taskId = *result.TaskId
	return
}
