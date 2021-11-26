resource "cds_vdc" "my_vdc" {
  vdc_name  = "Terraform(using)25"
  region_id = "CN_Beijing_A"
  public_network = {
    "ipnum"          = 4
    "qos"            = 10
    "name"           = "test-accPubNet"
    "floatbandwidth" = 200
    "billingmethod"  = "BandwIdth"
    "autorenew"      = 1
    "type"           = "Bandwidth_Multi_ISP_BGP"
  }
  add_public_ip = 4
}

// create private subnet for vdc
resource "cds_private_subnet" "my_private_subnet_1" {
  vdc_id  = cds_vdc.my_vdc.id
  name    = "private_25"
  type    = "auto"
  address = "192.168.0.0"
  mask    = 16
}

// create security group
resource "cds_security_group" "security_group_2" {
  name        = "test_tf_new_zz25"
  description = "New security group25"
  type        = "private"
  rule {
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

// create instance
resource "cds_instance" "my_instance2" {
  instance_name = "test_zz_04"
  region_id     = "CN_Beijing_A"
  image_id      = "Ubuntu_16.04_64"
  instance_type = "high_ccs"
  cpu           = 4
  ram           = 4
  vdc_id        = cds_vdc.my_vdc.id
  # password 和 public_key 二选一
  # public_key = file("/home/guest/.ssh/test.pub")
  password  = "123abc,.;"
  public_ip = "auto"
  private_ip {
    private_id = cds_private_subnet.my_private_subnet_1.id
    address    = "auto"
  }
  system_disk = {
    type = var.system_disk_type
    size = 324
    iops = 33
  }
  
  

  security_group_binding {
    type              = "private"
    subnet_id         = cds_private_subnet.my_private_subnet_1.id
    security_group_id = cds_security_group.security_group_2.id
  }
  #utc = true
}

data cds_data_source_instance "my_instance_data" {
    instance_id = cds_instance.my_instance2.id
    vdc_id =  cds_instance.my_instance2.vdc_id
    result_output_file = "data.json"
}