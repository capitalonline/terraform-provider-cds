package cds

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccInstance(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInstanceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccInstanceExists("cds_instance.my_instance"),
					resource.TestCheckResourceAttr("cds_instance.my_instance", "instance_name", "test_zz_002"),
				),
			},
		},
	})
}

func TestAccInstance_update(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInstanceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccInstanceExists("cds_instance.my_instance"),
					resource.TestCheckResourceAttr("cds_instance.my_instance", "instance_name", "test_zz_002"),
				),
			},
			{
				Config: testAccInstanceUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cds_instance.my_instance", "instance_name", "test_zz_002"),
					resource.TestCheckResourceAttr("cds_instance.my_instance", "data_disks.0.size", "200"),
				),
			},
		},
	})
}

func testAccInstanceExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		//logId := getLogId(contextNil)
		//ctx := context.WithValue(context.TODO(), "logId", logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		fmt.Println(rs.Primary.ID)

		return nil
	}
}

func testAccInstanceDestroy(s *terraform.State) error {
	//logId := getLogId(contextNil)
	//ctx := context.WithValue(context.TODO(), "logId", logId)
	//
	//service := SecurityGroupService{client: testAccProvider.Meta().(*CdsClient).apiConn}
	//for _, rs := range s.RootModule().Resources {
	//	if rs.Type != "cds_instance" {
	//		continue
	//	}
	//	time.Sleep(5 * time.Second)
	//	request := security_group.NewDescribeSecurityGroupRequest()
	//	request.SecurityGroupId = common.StringPtr(rs.Primary.ID)
	//	has, err := service.DescribeSecurityGroup(ctx, request)
	//	if err != nil {
	//		return err
	//	}
	//	if len(has.Data.SecurityGroup) == 0 {
	//		return nil
	//	}
	//	return fmt.Errorf("vpc not delete ok")
	//}
	return nil
}

const testAccInstanceConfig = `
resource "cds_vdc" "my_vdc" {
  vdc_name = "Terraform(using)1"
  region_id = "CN_Beijing_A"
  public_network = {
    "ipnum" = 4
    "qos" = 5
    "name" = "test-acc"
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

resource "cds_security_group" "security_group_1" {
  name = "test_tf_new_zz"
  description = "New security group 1"
  type ="private"
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
resource "cds_security_group" "security_group_2" {
  name = "test_tf_new_2"
  description = "New security group 2"
  type ="private"
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

resource "cds_instance" "my_instance" {
  instance_name       = "test_zz_002"
  region_id           = "CN_Beijing_A"
  image_id            = "Ubuntu_16.04_64"
  instance_type       = "high_ccs"
  cpu                 = 4
  ram                 = 4
  vdc_id              = cds_vdc.my_vdc.id
  password            = "123abc,.;"
  public_ip           = "auto"
  private_ip          = {
    "private_id" = cds_private_subnet.my_private_subnet_1.id
    "address" = "auto"
  }
  data_disks {
    size  =  100
    type  =  "high_disk"
  }
  security_group_binding {
   type = "private"
   subnet_id = cds_private_subnet.my_private_subnet_1.id
   security_group_id = cds_security_group.security_group_1.id
  }
}
`

const testAccInstanceUpdate = `
resource "cds_vdc" "my_vdc" {
  vdc_name = "Terraform(using)1"
  region_id = "CN_Beijing_A"
  public_network = {
    "ipnum" = 4
    "qos" = 5
    "name" = "test-acc"
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

resource "cds_security_group" "security_group_1" {
  name = "test_tf_new_zz"
  description = "New security group 1"
  type ="private"
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

resource "cds_security_group" "security_group_2" {
  name = "test_tf_new_2"
  description = "New security group 2"
  type ="private"
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

resource "cds_instance" "my_instance" {
  instance_name       = "test_zz_002"
  region_id           = "CN_Beijing_A"
  image_id            = "Ubuntu_16.04_64"
  instance_type       = "high_ccs"
  cpu                 = 4
  ram                 = 4
  vdc_id              = cds_vdc.my_vdc.id
  password            = "123abc,.;"
  public_ip           = "auto"
  private_ip          = {
    "private_id" = cds_private_subnet.my_private_subnet_1.id
    "address" = "auto"
  }
  data_disks {
    size  =  200
    type  =  "high_disk"
  }
  security_group_binding {
    type = "private"
    subnet_id = cds_private_subnet.my_private_subnet_1.id
    security_group_id = cds_security_group.security_group_2.id
  }
}
`
