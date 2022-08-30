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
    # Set mysql instance parameters
    parameters        = [
        {
          name  = "back_log"
          value = "8888"
        }
    ]
    # set mysql instance time_zone
    time_zone = "+08:00"

    #  Set  backup
    backup = {
        backup_type = var.backup_type
        desc = var.backup_desc
        db_list = var.backup_db_list
    }

    #  Set auto backup policy
    data_backups = {
        time_slot="00:00-01:00"
    #   Split databases with ","
        date_list="da1,db2,db3"
        sign = 0
    }
}

## Resource cds_mysql_account allows you to create a user
resource "cds_mysql_account" "user1" {
    instance_uuid= cds_mysql.mysql_example.id
    account_name = "testuser"
    password = "xxxxxxxx"
    account_type = "Normal"
    description = "测试账号"
#  to give permission
#  The db must be created in advance,and openapi does not support db creation. You need to create a DB manually
    operations = [{
        db_name = "db1"
        privilege = "DMLOnly"
    }
    ]
}

## Resource cds_mysql_readonly allows you to create a readonly instance
resource "cds_mysql_readonly" "readonly1" {
    instance_uuid = cds_mysql.mysql_example.id
    instance_name = "readonly"
#    You can find paas_goods_id in data.json.
#    The field name is available_read_only_config
    paas_goods_id = 1680
#    test_group_id = 0
    disk_type = "high_disk"
    disk_value = "500"
    amount = 1
}


data cds_data_source_mysql "mysql_data" {
    region_id           = cds_mysql.mysql_example.region_id
    instance_uuid       = cds_mysql.mysql_example.id
    instance_name       = cds_mysql.mysql_example.instance_name
    ip                  = cds_mysql.mysql_example.ip
    result_output_file  = "data.json" // availableDB, instances, regions
}
