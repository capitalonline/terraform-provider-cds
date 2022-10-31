
resource "cds_public_network" "pb1" {
  ip_num          = 4
  qos             = 10
  # To identify multiple different public networks, the 'name' field is required .
  name            = "terraform-copy"
  float_bandwidth = 200
  billing_method  = "BandwIdth"
  auto_renew      = 1
  type            = "Bandwidth_Multi_ISP_BGP"
  vdc_id          = "xxxxxxxx-xxxx"
}


data "cds_data_source_public_network" "pbn" {
  vdc_id    = cds_public_network.pb1.vdc_id
  public_id = cds_public_network.pb1.id
}