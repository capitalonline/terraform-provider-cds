output "test" {
  # reference to public networks attributes.
  value = data.cds_data_source_public_network.pbn.public_id
}