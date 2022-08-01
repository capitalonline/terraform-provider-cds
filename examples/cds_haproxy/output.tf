output "test" {
#  output or use my_haproxy_data's ip
  value = data.cds_data_source_haproxy.my_haproxy_data.ip
}