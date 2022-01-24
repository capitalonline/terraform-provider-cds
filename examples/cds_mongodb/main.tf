resource "cds_mongodb" "mongodb_example" {
    region_id         = var.region_id
    vdc_id            = var.vdc_id
    base_pipe_id      = var.base_pipe_id
    instance_name     = var.instance_name
    cpu               = var.cpu
    ram               = var.ram
    disk_type         = var.disk_type
    disk_value        = var.disk_value
    password          = var.password
    mongodb_version   = var.mongodb_version
}


data cds_data_source_mongodb "mongodb_data" {
    region_id           = cds_mongodb.mongodb_example.region_id
    instance_uuid       = cds_mongodb.mongodb_example.id
    instance_name       = cds_mongodb.mongodb_example.instance_name
    ip                  = cds_mongodb.mongodb_example.ip
    result_output_file  = "data.json" // availableDB, instances, regions
}
