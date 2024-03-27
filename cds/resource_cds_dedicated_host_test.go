package cds

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDedicatedHost(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccHaproxyStrategyDestory,
		Steps: []resource.TestStep{
			{
				Config: testAccDedicatedHostConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cds_dedicated_host.dedicated_host", "dedicated_host_name", "测试宿主机"),
				),
			},
		},
	})
}

const testAccDedicatedHostConfig = `
resource "cds_dedicated_host" "dedicated_host" {
    dedicated_host_cpu     = 16
    dedicated_host_good_id = 17147
    dedicated_host_limit   = 1
    dedicated_host_name    = "测试宿主机"
    dedicated_host_ram     = 28
    dedicated_host_type    = "ff52d30d-e0bc-4adb-98ff-898ee2528090"
    description_num        = true
    //prepaid_month          = 1
    auto_renew=1
    region_id              = "CN_Beijing_H"
}
`
