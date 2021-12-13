# create certificate

resource cds_certificate my_cds_certificate {
  certificate_name  = "my_cert"
  certificate       = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  private_key       = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}

# data source certificate
data cds_data_source_certificate "my_certificate" {
     result_output_file = "data.json"
     #ha_cert_list    computed by terraform apply
}