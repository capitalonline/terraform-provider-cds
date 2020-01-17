// VPC r for Module
resource "cds_vdc" "my_vdc" {
  vdc_name = var.vdc_name
  public_network = {
    "IPNum"          = 4
    "Qos"            = 20
    "Name"           = "test-accPubNet002"
    "FloatBandwidth" = 200
    "BillingMethod"  = "BandwIdth"
    "AutoRenew"      = 1
    "Type"           = "Bandwidth_BGP"
  }
}
resource "cds_private_subnet" "my_private_subnet_1" {
  vdc_id = cds_vdc.my_vdc.id
  name   = "private_1"
  type   = "private"
  addres = ""
  mask   = "26"
}

