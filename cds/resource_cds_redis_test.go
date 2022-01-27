package cds

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/redis"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCdsRedis(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCdsCheckRedisDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRedisConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCdsCheckRedisExists("cds_redis.my_redis"),
					resource.TestCheckResourceAttr("cds_redis.my_redis", "instance_name", "my_redis_test123"),
				),
			},
		},
	})
}

func testAccCdsCheckRedisExists(r string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		fmt.Println("testAccCdsCheckRedisExists")
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), "logId", logId)

		rs, ok := s.RootModule().Resources[r]

		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		service := RedisService{client: testAccProvider.Meta().(*CdsClient).apiConn}

		request := redis.NewDescribeDBInstancesRequest()

		request.InstanceUuid = common.StringPtr(rs.Primary.ID)
		has, err := service.DescribeRedis(ctx, request)

		if err != nil {
			return err
		}

		if len(has.Data) > 0 {
			return nil
		}
		return fmt.Errorf("redis not exists.")
	}
}

func testAccCdsCheckRedisDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), "logId", logId)

	service := RedisService{client: testAccProvider.Meta().(*CdsClient).apiConn}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cds_redis" {
			continue
		}
		time.Sleep(5 * time.Second)
		request := redis.NewDescribeDBInstancesRequest()
		request.InstanceUuid = common.StringPtr(rs.Primary.ID)
		has, err := service.DescribeRedis(ctx, request)
		if err != nil {
			return err
		}

		if len(has.Data) == 0 {
			return nil
		}
		return fmt.Errorf("redis not delete ok")

	}
	return nil
}

const testAccRedisConfig = `
resource "cds_redis" "my_redis" {
    region_id         = "CN_Beijing_A"
    vdc_id            = "c26529b9-f455-47d7-b10c-7eb1f1f72bd0"
    base_pipe_id      = "9fd9bf3e-540a-11ec-9d8e-96e971c86150"
    instance_name     = "my_redis_test123"
    architecture_type = 3
    ram               = 4
    redis_version           = "2.8"
    password          = "PassW@ord11"
}
`

const testAccRedisConfigUpdate = `

`

func TestDescribeRedis(t *testing.T) {
}
