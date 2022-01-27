resource "cds_redis" "redis_example" {
    region_id         = var.region_id
    vdc_id            = var.vdc_id
    base_pipe_id      = var.base_pipe_id
    instance_name     = var.instance_name
    architecture_type = var.architecture_type
    ram               = var.ram
    redis_version     = var.redis_version
    password          = var.password
}


data cds_data_source_redis "redis_data" {
    region_id           = cds_redis.redis_example.region_id
    instance_uuid       = cds_redis.redis_example.id
    instance_name       = cds_redis.redis_example.instance_name
    ip                  = cds_redis.redis_example.ip
    result_output_file  = "data.json" // availableDB, instances, regions
}
