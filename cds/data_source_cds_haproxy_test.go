package cds

import (
	"context"
	"log"
	"testing"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/haproxy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testHaproxyResource = `
resource cds_haproxy my_haproxy {
    data cds_data_source_haproxy my_data {
	  region_id = "CN_Beijing_A"
	  result_output_file = "data.json"
	}
}
`

func TestHaproxySource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccHaproxyDestory,
		Steps: []resource.TestStep{
			{
				Config: testHaproxyResource,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cds_haproxy.my_haproxy", "instance_name", "my_haproxy"),
				),
			},
		},
	})
}

func testAccHaproxyDestory(s *terraform.State) error {
	defer logElapsed("data_source.haproxy.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cds_haproxy" {
			continue
		}

		haproxyService := HaproxyService{client: testAccProvider.Meta().(*CdsClient).apiConn}

		request := haproxy.NewDescribeLoadBalancersRequest()

		request.InstanceName = common.StringPtr(rs.Primary.Attributes["instance_name"])

		response, err := haproxyService.DescribeHaproxy(ctx, request)
		if err != nil {
			return err
		}
		for _, entry := range response.Data {
			log.Println(entry.InstanceName)
		}
	}

	return nil
}
