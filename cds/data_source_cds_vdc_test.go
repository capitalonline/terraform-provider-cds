package cds

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCdsVdcDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCdsCheckVdcDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCdsVdcDataSource_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCdsCheckVdcExists("cds_vdc.my_vdc"),
				),
			},
		},
	})
}

func testAccCdsVdcDataSource_basic() string {
	return `
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

data "cds_data_source_vdc" "cds_data_source_vdc" {
  id = "${cds_vdc.my_vdc.id}"
  result_output_file="my_test_path"
}
`
}
