resource "cds_vdc" "my_vdc" {
  vdc_name  = var.vdc_name
  region_id = var.region_id
  public_network = {
    "ipnum"          = var.ipnum
    "qos"            = var.qos
    "name"           = var.public_network_name
    "floatbandwidth" = var.floatbandwidth
    "billingmethod"  = var.billingmethod
    "autorenew"      = var.autorenew
    "type"           = var.public_network_type
  }
  #add_public_ip = 8
  #delete_public_ip  = var.delete_public_ip
}

data "cds_data_source_vdc" "my_vdc_data" {
    id                 = cds_vdc.my_vdc.id
    vdc_name           = cds_vdc.my_vdc.vdc_name
    result_output_file = "data.json" 
    #vdc_list    computed by terraform apply
}