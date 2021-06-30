package cds

import (
	"context"
	"errors"
	"testing"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/haproxy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccHaproxyStrategy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccHaproxyStrategyDestory,
		Steps: []resource.TestStep{
			{
				Config: testAccHaproxyStrategyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cds_haproxy.my_haproxy", "instance_name", "my_terraform_haproxy"),
				),
			},
		},
	})
}

const testAccHaproxyStrategyConfig = `
resource cds_haproxy my_haproxy {
	region_id = "CN_Beijing_A"
	vdc_id = "XXXXXXXXX"
	base_pipe_id = "XXXXXXXXX"
	instance_name = "my_terraform_haproxy"
	paas_goods_id = 13721
	ips = [
		{
			pipe_id = "XXXXXXXXX"
			pipe_type = "private"
			segment_id = ""
		}
	]
	instance_uuid = "XXXXXXXXX"

	strategies = [
		{
			http_listeners = [
				{
				  server_timeout_unit = "s"
				  server_timeout = 120
				  sticky_session = "on"
				  acl_white_list = "192.168.4.1"
				  listener_mode = "http"
				  max_conn = 2021
				  connect_timeout_unit = "s"
				  scheduler = "roundrobin"
				  connect_timeout = 1000
				  client_timeout = 1000
				  listener_name = "terraform"
				  client_timeout_unit = "ms"
				  listener_port = 24353
				  backend_server = [
					{
					  ip = "192.168.3.1"
					  max_conn = 2021
					  port = 12313
					  weight = 1
					}
				  ]
				  certificate_ids = []
				}
			]
		}
	]
  }
`

func testAccHaproxyStrategyDestory(s *terraform.State) error {
	defer logElapsed("data_source.haproxy_strategy.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cds_haproxy_strategy" {
			continue
		}

		haproxyService := HaproxyService{client: testAccProvider.Meta().(*CdsClient).apiConn}

		request := haproxy.NewDescribeLoadBalancerStrategysRequest()

		request.InstanceUuid = common.StringPtr(rs.Primary.Attributes["id"])

		response, err := haproxyService.DescribeLoadBalancerStrategys(ctx, request)
		if err != nil {
			return err
		}

		if response.Data.HttpListeners != nil || response.Data.TcpListeners != nil {
			return errors.New("instance is not destroy")
		}
	}

	return nil
}
