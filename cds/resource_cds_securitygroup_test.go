package cds

import (
	"context"
	"fmt"
	"testing"
	"time"

	"terraform-provider-cds/cds-sdk-go/common"
	"terraform-provider-cds/cds-sdk-go/security_group"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccSecurityGroup(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityGroupExists("cds_security_group.security_group_1"),
					resource.TestCheckResourceAttr("cds_security_group.security_group_1", "name", "terraform-test"),
				),
			},
		},
	})
}

func TestAccSecurityGroup_update(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccSecurityGroupExists("cds_security_group.security_group_1"),
					resource.TestCheckResourceAttr("cds_security_group.security_group_1", "name", "terraform-test"),
				),
			},
			{
				Config: testAccSecurityGroupUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cds_security_group.security_group_1", "name", "terraform-test"),
				),
			},
		},
	})
}

func testAccSecurityGroupExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		//logId := getLogId(contextNil)
		//ctx := context.WithValue(context.TODO(), "logId", logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		//service := VdcService{client: testAccProvider.Meta().(*CdsClient).apiConn}
		//
		//request := vdc.DescribeVdcRequest()
		//request.VdcId = common.StringPtr(rs.Primary.ID)
		//has, err := service.DescribeVdc(ctx, request)
		//if err != nil {
		//	return err
		//}
		//if len(has.Data) > 0 {
		//	return nil
		//}
		fmt.Println(rs.Primary.ID)

		return nil
	}
}

func testAccSecurityGroupDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	service := SecurityGroupService{client: testAccProvider.Meta().(*CdsClient).apiConn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cds_security_group" {
			continue
		}
		time.Sleep(5 * time.Second)
		request := security_group.NewDescribeSecurityGroupRequest()
		request.SecurityGroupId = common.StringPtr(rs.Primary.ID)
		has, err := service.DescribeSecurityGroup(ctx, request)
		if err != nil {
			return err
		}
		if len(has.Data.SecurityGroup) == 0 {
			return nil
		}
		return fmt.Errorf("vpc not delete ok")
	}
	return nil
}

const testAccSecurityGroupConfig = `
resource "cds_security_group" "security_group_1" {
  name = "terraform-test"
  description = "New security group"
  type ="public"
  rule  {
    action        = "1"
    description   = "tf_rule_test"
    targetaddress = "120.78.170.188/28;120.78.170.188/28;120.78.170.188/28"
    targetport    = "70;90;8"
    localport     = "800"
    direction     = "all"
    priority      = "11"
    protocol      = "TCP"
    ruletype      = "ip"
  }
}

`

const testAccSecurityGroupUpdate = `
resource "cds_security_group" "security_group_1" {
  name = "terraform-test"
  description = "New security group"
  type ="public"
  rule  {
    action        = "1"
    description   = "tf_rule_test_1"
    targetaddress = "120.78.170.188/28;120.78.170.188/28;120.78.170.188/28"
    targetport    = "70;90;8"
    localport     = "80"
    direction     = "all"
    priority      = "11"
    protocol      = "TCP"
    ruletype      = "ip"
  }
  rule  {
    action        = "1"
    description   = "tf_rule_test_1"
    targetaddress = "120.78.170.188/28;120.78.170.188/28;120.78.170.188/28"
    targetport    = "70;90;8"
    localport     = "80"
    direction     = "all"
    priority      = "12"
    protocol      = "TCP"
    ruletype      = "ip"
  }
}
`
