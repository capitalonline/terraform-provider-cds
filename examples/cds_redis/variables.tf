variable "region_id" {
  description = "availability zone like CN_Beijing_A,CN_Beijing_E"
  default = "CN_Beijing_A"
}

variable "vdc_id" {
  description = "vdc id"
  default = "c26529b9-f455-47d7-b10c-7eb1f1f72bd0"
}

variable "base_pipe_id" {
  description = "vdc PrivateNetwork PrivateId"
  default = "9fd9bf3e-540a-11ec-9d8e-96e971c86150"
}

variable "instance_name" {
  default = "my_redis_test_e"
}

variable "architecture_type" {
  default = 3
}

variable "ram" {
  default = 4
}

variable "redis_version" {
  description = "redis version select by 2.8|4.0|5.0"
  default = "2.8"
}

variable "password" {
  description = "the password of redis"
  default = "PassW@ord123"
}


