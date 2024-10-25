package cds

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
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

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		fmt.Println(rs.Primary.ID)

		return nil
	}
}

func testAccInstanceDestroy(s *terraform.State) error {
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
  host_name = "123"
  private_ip           {
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

func TestName(t *testing.T) {
	s := []string{
		"www.baidu.com",
		"x.19hx",
		"x..19hx",
		"x-.19hx",
		"x--19hx",
		"x.19h-x",
		"x.19h-x_",
		"x.19h-xl.",
		".x.19h-x",
		"_x.19h-xl",
	}
	for _, value := range s {
		fmt.Println(value+"    ", v(value))
	}
}

func v(value string) (errs []error) {
	k := "hostname"
	if len(value) < 1 {
		errs = append(errs, fmt.Errorf("length of \"%s\"  should be more than or equal %d, the length of the current input value is %d", k, 1, len(value)))
	}
	if len(value) > 64 {
		errs = append(errs, fmt.Errorf("length of \"%s\"  should be less than or equal %d, the length of the current input value is %d", k, 64, len(value)))
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9\-.]+$`)
	if !re.MatchString(value) {
		errs = append(errs, errors.New(`static and transient hostnames should only contain a-z, A-Z, 0-9, "-", and "."`))
	}
	if strings.Contains(value, "..") ||
		strings.Contains(value, "--") ||
		strings.HasSuffix(value, ".") ||
		strings.HasSuffix(value, "-") ||
		strings.HasPrefix(value, "_") ||
		strings.HasPrefix(value, ".") {
		errs = append(errs, errors.New(`static and transient hostnames should cannot start or end with "." and "-". Consecutive "." or "-" are not allowed`))

	}
	return errs
}
