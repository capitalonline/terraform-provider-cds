# 因为HaProxy创建实例现在无法返回Instance_Uuid，在创建Haproxy实例后需要手动更新 Resource cds_haproxy.instance_uuid 字段后执行 terraform apply 使生效。  
# 如果不更新 Resource cds_haproxy.instance_uuid 字段, terraform将无法更新Haproxy实例及其对应的策略。  
# 执行 Data Resource cds_data_source_haproxy 将会获取 region_id 对应区域创建Haproxy及其策略的入参参数及HaProxy实例列表。 
# VDC 请求字段参考 https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#%E8%99%9A%E6%8B%9F%E6%95%B0%E6%8D%AE%E4%B8%AD%E5%BF%83%E7%9B%B8%E5%85%B3
# HaProxy 请求字段参考 https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#9describecacertificate


# 获取创建HaProxy实例所需要的参数数据包括（可用区，可用产品规格，已创建的HaProxy实例）
data cds_data_source_haproxy my_data {
    # 取对应区域的Haproxy实例列表
   region_id = "CN_Beijing_A"
   result_output_file = "data.json"
}

resource cds_haproxy my_haproxy {
	region_id = var.haproxy_zones[9]
    # vdc_id base_pipe_id 从 vdc 获取
	vdc_id = "XXXXXXXXXX"
	base_pipe_id = "XXXXXXXXXX"
	instance_name = "my_terraform_haproxy"
    # paas_goods_id 从data.json PaasGoodsId 获取
	paas_goods_id = 13721
	ips = [
		{
            # pipe_id pipe_type vdc网络
			pipe_id = "f02fa2fa-d57f-11eb-a77a-3eb98755b2a4"
			pipe_type = "private"
			segment_id = ""
		}
	]
	instance_uuid = "bc1fe739-e01e-43ad-9575-1080ee9565d6"
}

# terraform 因可选嵌套对象无法灵活设置对象字段的必选和可选性，所以可选字段必须保留字段赋空值 [参考](https://github.com/hashicorp/terraform-provider-google/issues/3928)
resource cds_haproxy_strategy my_haproxy_strategy {
	instance_uuid = "d2e7e2f2-15b4-4cdf-b1db-33ebdc7b01b3"
	http_listeners = [{
		server_timeout_unit = "s"
		server_timeout = 1200
		sticky_session = "on"
		acl_white_list = "192.168.4.1"
		listener_mode = "http"
		max_conn = 2021
		connect_timeout_unit = "s"
		scheduler = "roundrobin"
		connect_timeout = 1200
		client_timeout = 1001
		listener_name = "terraform"
		client_timeout_unit = "ms"
		listener_port = 24353
		backend_server = [{
			ip = "192.168.3.1"
			max_conn = 2021
			port = 12313
			weight = 1
		}]
		certificate_ids = []
	}]
}

resource cds_certificate "my_certificate" {
	certificate_name = "my_certificate"
	certificate = "XXXXXXXX"
	private_key = "XXXXXXXX"
}

data cds_data_source_certificate "my_certificate_data" {
    result_output_file = "data.json"
}

