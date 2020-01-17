package cds

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccPrivateSubnet(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPrivateSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPrivateSubnetConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccPrivateSubnetExists("cds_private_subnet.my_private_subnet_1"),
					resource.TestCheckResourceAttr("cds_private_subnet.my_private_subnet_1", "name", "private_1"),
				),
			},
		},
	})
}

func testAccPrivateSubnetExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		fmt.Println(rs.Primary.ID)

		return nil
	}
}

func testAccPrivateSubnetDestroy(s *terraform.State) error {
	return nil
}

const testAccPrivateSubnetConfig = `
resource "cds_vdc" "my_vdc" {
  vdc_name = "test_tf_bew_zz"
  region_id = "CN_Beijing_A"
  public_network = {
    "ipnum" = 4
    "qos" = 20
    "name" = "test-accPubNet002"
    "floatbandwidth" = 200
    "billingmethod" = "BandwIdth"
    "autorenew" = 1
    "type" = "Bandwidth_BGP"
  }
}
resource "cds_private_subnet" "my_private_subnet_1" {
  vdc_id = cds_vdc.my_vdc.id
  name = "private_1"
  type = "auto"
  address = "192.168.0.0"
  mask = 16
}

`
