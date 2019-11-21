package connectivity

import (
	"terraform-provider-cds/cds-sdk-go/common"
	"terraform-provider-cds/cds-sdk-go/common/profile"
	"terraform-provider-cds/cds-sdk-go/instance"
	"terraform-provider-cds/cds-sdk-go/security_group"
	"terraform-provider-cds/cds-sdk-go/security_group_rule"
	"terraform-provider-cds/cds-sdk-go/task"
	"terraform-provider-cds/cds-sdk-go/vdc"
)

// client for all Capitalonline data service
type CdsClient struct {
	Region     string
	SecretId   string
	SecretKey  string
	vdcConn    *vdc.Client
	vdcGetConn *vdc.Client
	sgConn     *security_group.Client
	sgGetConn  *security_group.Client
	vmConn     *instance.Client
	taskConn   *task.Client
	sgrConn    *security_group_rule.Client
	sgrGetConn *security_group_rule.Client
}

func NewCdsClient(secretId, secretKey, region string) *CdsClient {
	var cdsClient CdsClient
	cdsClient.SecretId, cdsClient.SecretKey, cdsClient.Region = secretId, secretKey, region
	return &cdsClient
}

// get vdc client for service
func (me *CdsClient) UseVdcClient() *vdc.Client {
	if me.vdcConn != nil {
		return me.vdcConn
	}
	credential := common.NewCredential(me.SecretId, me.SecretKey)
	client, _ := vdc.NewClient(credential, me.Region, clientProfile("POST"))
	var round LogRoundTripper
	client.WithHttpTransport(&round)
	me.vdcConn = client
	return me.vdcConn
}

// get vdc client for service
func (me *CdsClient) UseVdcGetClient() *vdc.Client {
	if me.vdcGetConn != nil {
		return me.vdcGetConn
	}
	credential := common.NewCredential(me.SecretId, me.SecretKey)
	client, _ := vdc.NewClient(credential, me.Region, clientProfile("GET"))
	var round LogRoundTripper
	client.WithHttpTransport(&round)
	me.vdcGetConn = client
	return me.vdcGetConn
}

// get security group client for service
func (me *CdsClient) UseSecurityGroupClient() *security_group.Client {
	if me.sgConn != nil {
		return me.sgConn
	}
	credential := common.NewCredential(me.SecretId, me.SecretKey)
	client, _ := security_group.NewClient(credential, me.Region, clientProfile("POST"))
	var round LogRoundTripper
	client.WithHttpTransport(&round)
	me.sgConn = client
	return me.sgConn
}

func (me *CdsClient) UseSecurityGroupGetClient() *security_group.Client {
	if me.sgGetConn != nil {
		return me.sgGetConn
	}
	credential := common.NewCredential(me.SecretId, me.SecretKey)
	client, _ := security_group.NewClient(credential, me.Region, clientProfile("GET"))
	var round LogRoundTripper
	client.WithHttpTransport(&round)
	me.sgGetConn = client
	return me.sgGetConn
}

func (me *CdsClient) UseCvmClient() *instance.Client {
	if me.vmConn != nil {
		return me.vmConn
	}
	credential := common.NewCredential(me.SecretId, me.SecretKey)
	client, _ := instance.NewClient(credential, me.Region, clientProfile("POST"))
	var round LogRoundTripper
	client.WithHttpTransport(&round)
	me.vmConn = client
	return me.vmConn
}

func (me *CdsClient) UseCvmGetClient() *instance.Client {
	if me.vmConn != nil {
		return me.vmConn
	}
	credential := common.NewCredential(me.SecretId, me.SecretKey)
	client, _ := instance.NewClient(credential, me.Region, clientProfile("GET"))
	var round LogRoundTripper
	client.WithHttpTransport(&round)
	me.vmConn = client
	return me.vmConn
}

func (me *CdsClient) UseSecurityRuleClient() *security_group_rule.Client {
	if me.sgrConn != nil {
		return me.sgrConn
	}
	credential := common.NewCredential(me.SecretId, me.SecretKey)
	client, _ := security_group_rule.NewClient(credential, me.Region, clientProfile("POST"))
	var round LogRoundTripper
	client.WithHttpTransport(&round)
	me.sgrConn = client
	return me.sgrConn
}

func (me *CdsClient) UseSecurityRuleGetClient() *security_group_rule.Client {
	if me.sgrGetConn != nil {
		return me.sgrGetConn
	}
	credential := common.NewCredential(me.SecretId, me.SecretKey)
	client, _ := security_group_rule.NewClient(credential, me.Region, clientProfile("GET"))
	var round LogRoundTripper
	client.WithHttpTransport(&round)
	me.sgrGetConn = client
	return me.sgrGetConn
}

// get task client
func (me *CdsClient) UseTaskGetClient() *task.Client {
	if me.taskConn != nil {
		return me.taskConn
	}
	credential := common.NewCredential(me.SecretId, me.SecretKey)
	client, _ := task.NewClient(credential, me.Region, clientProfile("GET"))
	var round LogRoundTripper
	client.WithHttpTransport(&round)
	me.taskConn = client
	return me.taskConn
}

func clientProfile(method string) *profile.ClientProfile {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = method
	return cpf
}
