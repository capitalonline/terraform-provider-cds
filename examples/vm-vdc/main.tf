
// VPC r for Module
resource "cds_vdc" "cds_vdc" {
  vdc_name       = var.vdc_name
  public_network = {
    "IPNum"           = 4
    "Qos"             = 20
    "Name"            ="test-accPubNet002"
    "FloatBandwidth"  =200
    "BillingMethod"   ="BandwIdth"
    "AutoRenew"       =  1
    "Type"            ="Bandwidth_BGP"
  }
}

