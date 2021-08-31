resource "cds_vdc" "my_vdc" {
  vdc_name  = "Terraform(using1)"
  region_id = "CN_Beijing_A"
  public_network = {
    "ipnum"          = 4
    "qos"            = 10
    "name"           = "test-accPubNet"
    "floatbandwidth" = 200
    "billingmethod"  = "BandwIdth"
    "autorenew"      = 1
    "type"           = "Bandwidth_China_Telecom"
  }
  add_public_ip = 2
}