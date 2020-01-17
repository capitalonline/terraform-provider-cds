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

