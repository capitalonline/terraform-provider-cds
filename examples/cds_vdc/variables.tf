variable "region_id" {
  description = "availability zone like CN_Beijing_A,CN_Beijing_E"
  default = "CN_Beijing_A"
}

variable "vdc_name" {
  default = "Terraform(using)test1"
}

variable "ipnum" {
  default = 4
}

variable "qos" {
  default = 10
}

variable "public_network_name" {
  default = "test-accPubNet"
}

variable "floatbandwidth" {
  default = 200
}

variable "billingmethod" {
  default = "BandwIdth"
}

variable "autorenew" {
  default = 1
}

variable "public_network_type" {
  default = "Bandwidth_China_Telecom"
}

variable "add_public_ip_num" {
  description = "add public ip num like 4,8,16,32,64,128,254"
  default = 8  
}

variable "delete_public_ip" {
  description = "public network SegmentId"
  default = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}