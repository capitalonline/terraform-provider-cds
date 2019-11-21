variable "vdc_name" {
  //  description = "The vdc name used to launch a new vdc."
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

variable "vdc_id" {
  default = "6a0ff09f-8f54-4ae9-ad70-1673392af6d0"
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


