package cds

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCdsSecurityGroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCdsSecurityGroupDataSource_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSecurityGroupExists("cds_security_group.security_group_1"),
				),
			},
		},
	})
}

func testAccCdsSecurityGroupDataSource_basic() string {
	return `
resource "cds_security_group" "security_group_1" {
  name = "test_tf_new_zz"
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

data "cds_data_source_security_group" "cds_data_source" {
  id = "${cds_security_group.security_group_1.id}"
  result_output_file="my_test_path"
}
`
}
