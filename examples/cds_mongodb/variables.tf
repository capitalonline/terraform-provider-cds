variable "region_id" {
  description = "availability zone like CN_Beijing_A,CN_Beijing_E"
  default = "CN_Beijing_A"
}

variable "vdc_id" {
  description = "vdc id"
  default = "xxxxxxxxxxxx"
}

variable "base_pipe_id" {
  description = "vdc PrivateNetwork PrivateId"
  default = "xxxxxxxxxxxx"
}

variable "instance_name" {
  default = "my_mongodb_test_a"
}

variable "cpu" {
  default = 1
}

variable "ram" {
  default = 2
}

variable "mongodb_version" {
  description = "mongodb version select by 4.0.3|3.6.7|3.2.21"
  default = "4.0.3"
}

variable "password" {
  description = "the password of mongodb"
  default = "xxxxxxxxxxxx"
}


variable "disk_type" {
  default = "ssd_disk"
}

variable "disk_value" {
  default = "100"
}


