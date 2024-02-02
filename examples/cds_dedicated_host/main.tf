


resource cds_dedicated_host dedicated_host {
	region_id 		= "CN_Beijing_B"
	dedicated_host_type = "xxx"
	dedicated_host_good_id = 1
	dedicated_host_name= "测试宿主机"
	dedicated_host_cpu = 16
	dedicated_host_ram = 32
	dedicated_host_limit= 1

	prepaid_month = 1
	auto_renew = 1
	amount = 1
	description_num = true
	subject_id=101
}

data "cds_data_source_dedicated_host" "test" {
	host_id = "xxx"
}
