resource "cds_mysql" "mysql_example" {
    region_id       = var.region_id
    vdc_id          = var.vdc_id
    base_pipe_id    = var.base_pipe_id
    instance_name   = var.instance_name
    cpu             = var.cpu
    ram             = var.ram
    disk_type       = var.disk_type
    disk_value      = var.disk_value
}

data cds_data_source_mysql "mysql_data" {
    region_id = "CN_Beijing_A"
    result_output_file = "data.json" // availableDB, instances, regions
}