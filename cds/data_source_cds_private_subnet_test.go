package cds

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccPrivateSubnetDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPrivateSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCdsPrivateSubnetDataSource_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccPrivateSubnetExists("cds_private_subnet.my_private_subnet_1"),
					resource.TestCheckResourceAttr("cds_private_subnet.my_private_subnet_1", "name", "private_1"),
				),
			},
		},
	})
}

func testAccCdsPrivateSubnetDataSource_basic() string {
	return `
resource "cds_private_subnet" "my_private_subnet_1" {
  vdc_id = "9e53b1e0-c49a-4827-bea0-af8cb4857b30"
  name = "private_1"
  type = "auto"
  address = "192.168.0.0"
  mask = 16
}
data "cds_data_source_private_subnet" "cds_data_source" {
  vdc_id = "9e53b1e0-c49a-4827-bea0-af8cb4857b30"
  result_output_file="my_test_path"
}
`
}
