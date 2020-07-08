package cds

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/vdc"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCdsCloudVdc(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCdsCheckVdcDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCdsCheckVdcExists("cds_vdc.my_vdc"),
					resource.TestCheckResourceAttr("cds_vdc.my_vdc", "vdc_name", "terraform-test"),
				),
			},
		},
	})
}

func TestAccCdsCloudVdc_update(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCdsCheckVdcDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCdsCheckVdcExists("cds_vdc.my_vdc"),
					resource.TestCheckResourceAttr("cds_vdc.my_vdc", "vdc_name", "terraform-test"),
				),
			},
			{
				Config: testAccVpcConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cds_vdc.my_vdc", "vdc_name", "terraform-test"),
				),
			},
		},
	})
}

func testAccCdsCheckVdcExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), "logId", logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		service := VdcService{client: testAccProvider.Meta().(*CdsClient).apiConn}

		request := vdc.DescribeVdcRequest()
		request.VdcId = common.StringPtr(rs.Primary.ID)
		has, err := service.DescribeVdc(ctx, request)
		if err != nil {
			return err
		}
		if len(has.Data) > 0 {
			return nil
		}

		return fmt.Errorf("vpc not exists.")
	}
}

func testAccCdsCheckVdcDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	service := VdcService{client: testAccProvider.Meta().(*CdsClient).apiConn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cds_vdc" {
			continue
		}
		time.Sleep(5 * time.Second)
		request := vdc.DescribeVdcRequest()
		request.VdcId = common.StringPtr(rs.Primary.ID)
		has, err := service.DescribeVdc(ctx, request)
		if err != nil {
			return err
		}
		if len(has.Data) == 0 {
			return nil
		}
		return fmt.Errorf("vpc not delete ok")
	}
	return nil
}

const testAccVpcConfig = `
resource "cds_vdc" "my_vdc" {
  vdc_name = "terraform-test"
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
`

const testAccVpcConfigUpdate = `
resource "cds_vdc" "my_vdc" {
  vdc_name = "terraform-test"
  region_id = "CN_Beijing_A"
  public_network = {
    "ipnum" = 5
    "qos" = 50
    "name" = "test-accPubNet002"
    "floatbandwidth" = 200
    "billingmethod" = "BandwIdth"
    "autorenew" = 1
    "type" = "Bandwidth_BGP"
  }
}
`
