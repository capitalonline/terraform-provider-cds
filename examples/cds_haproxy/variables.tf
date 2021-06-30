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