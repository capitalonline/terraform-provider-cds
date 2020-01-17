// create vdc
resource "cds_vdc" "my_vdc" {
  vdc_name  = "Terraform(using)"
  region_id = "CN_Beijing_A"
  public_network = {
    "ipnum"          = 4
    "qos"            = 20
    "name"           = "test-accPubNet"
    "floatbandwidth" = 200
    "billingmethod"  = "BandwIdth"
    "autorenew"      = 1
    "type"           = "Bandwidth_BGP"
  }
}
resource "cds_private_subnet" "my_private_subnet_1" {
  vdc_id  = cds_vdc.my_vdc.id
  name    = "private_1"
  type    = "private"
  address = ""
  mask    = "26"
}

