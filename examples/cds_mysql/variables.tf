variable "region_id" {
  description = "availability zone like CN_Beijing_A,CN_Beijing_E"
  default = "CN_Beijing_E"
}

variable "vdc_id" {
  default = "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}

variable "base_pipe_id" {
  description = "vdc PrivateNetwork PrivateId"
  default = "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}

variable "instance_name" {
  default = "testMysql05"
}

variable "cpu" {
  default = 2
}

variable "ram" {
  default = 4
}

variable "disk_type" {
  description = "disk type ssd_disk|high_disk"
  default = "ssd_disk"
}

variable "disk_value" {
  description = "disk size default 100G "
  default = "100"
}

variable "mysql_version" {
  description = "selective 5.6,5.7,8.0"
  default = "5.7"
}

variable "architecture_type" {
  description = "architecture type 0:basic edition|1:master-slave edition"
  default = "1"
}
variable "compute_type" {
  description = "compute type 0:common type"  
  default = "0"
}

variable "backup_type" {
  type = string
  description = "backup type  backup_type|logical-backup"
  default = "logical-backup"
}

variable "backup_desc" {
  type = string
  default = ""
}

variable "backup_db_list" {
  type = string
  description = "db list,split with ','"
  default = "database1,database2"
}