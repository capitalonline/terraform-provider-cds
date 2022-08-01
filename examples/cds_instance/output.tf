output "public_id" {
#  output or use my_vdc's public_id
  value = data.cds_data_source_vdc.my_vdc_data.vdc_list[0].public_network[0].public_id
}

output "segment_id" {
  #  output or use my_vdc's segment_id
  value = data.cds_data_source_vdc.my_vdc_data.vdc_list[0].public_network[0].segments[0].segment_id

}