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
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance uuid. [View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#5createprivilegedaccount)",
			},
			"account_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Account name. [View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#5createprivilegedaccount)",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password. [View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#5createprivilegedaccount)",
			},
			"account_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Account type, supports creating high privilege users and normal users. Possible values: High privilege user: \"Super\" 、Normal user: \"Normal\". Note: An instance can only have one high privilege account.[View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#5createprivilegedaccount)",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Description. [View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#5createprivilegedaccount)",
			},
			"operations": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Modify db privilege.[View Document](https://github.com/capitalonline/openapi/blob/master/MySQL%E6%A6%82%E8%A7%88.md#OperationsObj)",
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Database name.",
						},
						"privilege": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Privilege. ReadWrite: Read and write permission. DMLOnly: Only DML (Data Manipulation Language) permission. ReadOnly: Read-only permission. DDLOnly: Only DDL (Data Definition Language) permission",
						},
					},
				},
			},
		},
		Description: "Mysql account.\n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
resource "cds_mysql_account" "user1" {
    instance_uuid= cds_mysql.mysql_example.id
    account_name = "testuser"
    password = "xxxxxxxx"
    account_type = "Normal"
    description = "test"
#  to give permission
#  The db must be created in advance,and openapi does not support db creation. You need to create a DB manually
    operations = [{
        db_name = "db1"
        privilege = "DMLOnly"
    }
    ]
}
` +
			"\n```",
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
