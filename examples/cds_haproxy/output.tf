output "test" {
#  output or use my_haproxy_data's ip
  value = data.cds_data_source_haproxy.my_haproxy_data.ip
}
output "segment_id" {
  #  output or use my_vdc's segment_id
  value = data.cds_data_source_vdc.my_vdc_data.vdc_list[0].public_network[0].segments[0].segment_id

}

