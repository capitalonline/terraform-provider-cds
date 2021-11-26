resource "cds_mysql" "mysql_example" {
    region_id         = var.region_id
    vdc_id            = var.vdc_id
    base_pipe_id      = var.base_pipe_id
    instance_name     = var.instance_name
    cpu               = var.cpu
    ram               = var.ram
    disk_type         = var.disk_type
    disk_value        = var.disk_value
    mysql_version     = var.mysql_version
    architecture_type = var.architecture_type
    compute_type      = var.compute_type
}


data cds_data_source_mysql "mysql_data" {
    region_id           = cds_mysql.mysql_example.region_id
    instance_uuid       = cds_mysql.mysql_example.id
    instance_name       = cds_mysql.mysql_example.instance_name
    ip                  = cds_mysql.mysql_example.ip
    result_output_file  = "data.json" // availableDB, instances, regions
}
