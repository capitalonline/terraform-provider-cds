# 因为HaProxy创建实例现在无法返回Instance_Uuid，在创建Haproxy实例后需要手动更新 Resource cds_haproxy.instance_uuid 字段后执行 terraform apply 使生效。  
# 如果不更新 Resource cds_haproxy.instance_uuid 字段, terraform将无法更新Haproxy实例及其对应的策略。  
# 执行 Data Resource cds_data_source_haproxy 将会获取 region_id 对应区域创建Haproxy及其策略的入参参数及HaProxy实例列表。 
# VDC 请求字段参考 https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#%E8%99%9A%E6%8B%9F%E6%95%B0%E6%8D%AE%E4%B8%AD%E5%BF%83%E7%9B%B8%E5%85%B3
# HaProxy 请求字段参考 https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#9describecacertificate



resource cds_haproxy my_haproxy {
	region_id 		= "XXXXXXXX"
	//region_id 	= var.haproxy_zones[8]
    # vdc_id base_pipe_id 从 vdc 获取
	//vdc_id 		= "XXXXXXXX"
	//base_pipe_id  = "XXXXXXXX"
	vdc_id          = var.vdc_id
    base_pipe_id    = var.base_pipe_id
	instance_name   = var.instance_name
    # paas_goods_id 从data.json PaasGoodsId 获取
	cpu = var.cpu
	ram = var.ram
	ips = [
		{
            # pipe_id pipe_type vdc网络
			# pipe_id is PrivateNetwork PrivateId from vdc info 
			pipe_id    = var.pipe_id
			pipe_type  = "private"
			segment_id = var.segment_id
		},
		#This parameter is required if you want to create a public network(如创建公网，则需要)
		{
			# pipe_id is PublicNetwork PublicId from vdc info 
			pipe_id    = var.public_pipe_id
			pipe_type  = "public"
			segment_id = var.segment_id
		}
	]
	http_listeners = [{
		server_timeout_unit = var.server_timeout_unit
		server_timeout      = var.server_timeout
		sticky_session      = var.sticky_session
		acl_white_list      = var.acl_white_list
		listener_mode       = var.listener_mode
		max_conn            = var.max_conn
		connect_timeout_unit = var.connect_timeout_unit
		scheduler           = var.scheduler
		connect_timeout     = var.connect_timeout
		client_timeout      = var.client_timeout
		listener_name       = var.listener_name
		client_timeout_unit = var.client_timeout_unit
		listener_port       = var.listener_port
		backend_server = [{
			ip       = var.backend_server_ip
			max_conn = var.backend_server_max_conn
			port     = var.backend_server_port
			weight   = var.backend_server_weight
		}]
		certificate_ids = []

#		The parameters option is a list,only one element at most
#		option = [{
#			httpchk = {
#				method = "GET"
#				uri = "/health"
#			}
#		}]
#		The parameters session_persistence is a list,only one element at most
#		session_persistence  = [
#			{
#				key = "test"
#				mode = 1
#				timer = {
#					max_idle=33
#					max_life=44
#				}
#			}
#		]
	}]
}

data cds_data_source_haproxy "my_haproxy_data" {
	instance_uuid      = "xxxxxxxxxxxxx" 
	instance_name      = cds_haproxy.my_haproxy.instance_name
	region_id          = cds_haproxy.my_haproxy.region_id
	result_output_file = "data.json"
	#ha_list  computed by terraform apply
}
