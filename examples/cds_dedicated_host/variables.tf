variable haproxy_name {
    type = string
    default = "terraform_haproxy"
}

# 可参考 data.json
variable haproxy_zones {
    type = list(string)
    default = [
        "US_LosAngeles_A",
        "EUR_Germany_A",
        "APAC_Tokyo_A",
        "CN_Hongkong_A",
        "APAC_Singapore_A",
        "EUR_Netherlands_A",
        "US_NewYork_A",
        "CN_Beijing_B",
        "CN_Beijing_E",
        "CN_Beijing_A",
        "US_Dallas_A",
        "CN_Taipei_A",
        "CN_Wuxi_A",
        "APAC_Seoul_A",
        "CN_Beijing_C",
        "CN_Guangzhou_A",
        "CN_Shanghai_A",
        "CN_Beijing_H",
        "US_Dallas_G",
        "US_Dallas_J",
        "EUR_Germany_B",
        "APAC_Singapore_D",
        "US_Virginia_A",
        "CN_Shanghai_C",
    ]
}

# 可参考 data.json
variable haproxy_goods {
    type = map(number)
    default = {
        "1C2G" = 13721,
        "2C4G" = 13724,
        "4C8G" = 13727,
        "8C16G" = 13730,
        "16C32G" = 19733
    }
}

variable "vdc_id" {
  default = "xxxxxxxxxxxxxxxxxxxxxxxx"
}

variable "base_pipe_id" {
  description = "PrivateNetwork id from data source vdc"
  default = "xxxxxxxxxxxxxxxxxxxxxxxx"
}

variable "instance_name" {
  default = "my_terraform_haproxy"
}

variable "cpu" {
  default = 1
}

variable "ram" {
  default = 2
}

variable "pipe_id" {
  description = "PrivateNetwork id from data source vdc"
  default = "xxxxxxxxxxxxxxxxxxxxxxxx"
}

variable "public_pipe_id" {
  description = "PublicNetwork PublicId from data source vdc"
  default = "xxxxxxxxxxxxxxxxxxxxxxxx"
}

variable "pipe_type" {
  description = "public|private"
  default = "private"
}

variable "segment_id" {
  description = "public network segment_id from data source vdc"
  default = "xxxxxxxxxxxxxxxxxxxxxxxx"
}

variable "server_timeout_unit" {
  default = "s"
}

variable "server_timeout" {
  default = 1300
}

variable "sticky_session" {
  default = "on"
}

variable "acl_white_list" {
  description = "Set a whitelist, for example 192.168.12.1,192.168.1.1/20"
  default = "192.168.9.1"
}

variable "listener_mode" {
  default = "http"
}

variable "max_conn" {
  default = 2022
}

variable "connect_timeout_unit" {
  default = "s"
}

variable "scheduler" {
  default = "roundrobin"
}

variable "connect_timeout" {
  default = 1300
}

variable "client_timeout" {
  default = 1002
}

variable "listener_name" {
  default = "terraform"
}

variable "client_timeout_unit" {
  default = "ms"
}

variable "listener_port" {
  default = 24354
}

variable "backend_server_ip" {
  default = "192.168.12.1"
}

variable "backend_server_max_conn" {
  default = 2022
}

variable "backend_server_port" {
  default = 12314
}

variable "backend_server_weight" {
  description = "between 1-256"
  default = 255
}

