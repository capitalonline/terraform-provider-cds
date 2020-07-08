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

// list vdc
data "cds_data_source_vdc" "vdclist" {
  vdc_id = ""
  vdc_name = ""
  // this is optional
  result_output_file = "somewhere_to_save"
}

// output the vdc list to the console
output "list_vdc" {
  value = data.cds_data_source_vdc.vdclist
}
