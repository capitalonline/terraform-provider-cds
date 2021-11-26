variable "vdc_name" {
  description = "The vdc name used to launch a new vdc."
  default = "TF-testAccVdc"
}
variable "availability_zone" {
  default = "CN_Beijing_A"
}

variable "image_id" {
  default = "Ubuntu_16.04_64"
}

variable "instance_type" {
  default = "high_ccs"
}

variable "instance_name" {
  default = "terraform-testing-zzz"
}

variable "cpu" {
  default = 4
}

variable "ram" {
  default = 4
}

variable "password" {
  default = "123abc,.;"
}

variable "public_address" {
  default = "auto"
}

variable "group_name" {
  default = "tf_test_zz_1"
}

variable "group_type" {
  default = "public"
}

variable "system_disk_type" {
  default = "ssd_system_disk"
}

variable "system_disk_size" {
  default = 100
}

variable "system_disk_iops" {
  default = 5
}

variable "data_disks_type" {
  default = "ssd_disk"
}

variable "data_disks_size" {
  default = 150
}

variable "data_disks_iops" {
  default = 10
}


