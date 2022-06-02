package cds

import (
	"context"
	"errors"
	"fmt"
	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/mysql"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCdsMySQLAccount() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCdsMySQLAccount,
		Read:   readResourceCdsMySQLAccount,
		Update: updateResourceCdsMySQLAccount,
		Delete: deleteResourceCdsMySQLAccount,
		Schema: map[string]*schema.Schema{
			"instance_uuid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"operations": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "modify db privilege",
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"privilege": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func createResourceCdsMySQLAccount(data *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_mysql_account.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	mysqlService := MySQLService{client: meta.(*CdsClient).apiConn}
	request := mysql.NewCreatePrivilegedAccountRequest()
	instanceUuid := data.Get("instance_uuid")
	request.InstanceUuid = common.StringPtr(instanceUuid.(string))

	accountName := data.Get("account_name")
	request.AccountName = common.StringPtr(accountName.(string))

	password := data.Get("password")
	request.Password = common.StringPtr(password.(string))

	accountType := data.Get("account_type")
	request.AccountType = common.StringPtr(accountType.(string))

	description, ok := data.GetOk("description")
	if ok {
		request.Description = common.StringPtr(description.(string))
	}

	response, err := mysqlService.CreatePrivilegedAccount(ctx, request)
	if err != nil {
		return err
	}
	if *response.Code != "Success" {
		fmt.Errorf("create mysql account request,api response:%v", response)
	}
	if err := waitMysqlRunning(ctx, mysqlService, *request.InstanceUuid); err != nil {
		return err
	}
	optionsResource, ok := data.GetOk("operations")
	if ok && optionsResource != nil {
		optionList := optionsResource.([]interface{})
		req := mysql.NewModifyDbPrivilegeRequest()
		req.InstanceUuid = request.InstanceUuid
		req.AccountName = request.AccountName
		options := make([]*mysql.ModifyDbPrivilegeOperation, 0, len(optionList))
		for _, item := range optionList {
			option := item.(map[string]interface{})

			options = append(options, &mysql.ModifyDbPrivilegeOperation{
				DBName:    common.StringPtr(option["db_name"].(string)),
				Privilege: common.StringPtr(option["privilege"].(string)),
			})
		}
		req.Operations = options
		if len(options) > 0 {
			resp, err := mysqlService.ModifyDbPrivilege(ctx, req)
			if err != nil {
				return err
			}
			if *resp.Code != "Success" {
				fmt.Errorf("modify db privilege request,api response:%v", resp)
			}
		}
	}
	return readResourceCdsMySQLAccount(data, meta)
}

func readResourceCdsMySQLAccount(data *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_mysql_account.read")()
	userId := data.Get("account_name")
	data.SetId(userId.(string))
	return nil
}

func updateResourceCdsMySQLAccount(data *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cds_mysql_account.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	mysqlService := MySQLService{client: meta.(*CdsClient).apiConn}
	instanceUuid := data.Get("instance_uuid")
	accountName := data.Get("account_name")
	if instanceUuid == nil || accountName == nil {
		return errors.New("instance_uuid and account_name cannot be nil")
	}
	if data.HasChange("operations") {
		_, newData := data.GetChange("operations")
		optionList := newData.([]interface{})
		req := mysql.NewModifyDbPrivilegeRequest()

		req.InstanceUuid = common.StringPtr(instanceUuid.(string))
		req.AccountName = common.StringPtr(accountName.(string))
		options := make([]*mysql.ModifyDbPrivilegeOperation, 0, len(optionList))
		for _, item := range optionList {
			option := item.(map[string]interface{})

			options = append(options, &mysql.ModifyDbPrivilegeOperation{
				DBName:    common.StringPtr(option["db_name"].(string)),
				Privilege: common.StringPtr(option["privilege"].(string)),
			})
		}
		req.Operations = options
		if len(options) > 0 {
			resp, err := mysqlService.ModifyDbPrivilege(ctx, req)
			if err != nil {
				return err
			}
			if *resp.Code != "Success" {
				fmt.Errorf("modify db privilege request,api response:%v", resp)
			}
		}
	}
	return nil
}

func deleteResourceCdsMySQLAccount(data *schema.ResourceData, meta interface{}) error {
	return nil
}
