


resource cds_dedicated_host dedicated_host {
	region_id 		= "CN_Beijing_h"
	dedicated_host_type = "ff52d30d-e0bc-4adb-98ff-898ee2528090"
	dedicated_host_good_id = 1
	dedicated_host_name= "测试宿主机"
	dedicated_host_cpu = 16
	dedicated_host_ram = 28
	dedicated_host_limit= 1

	prepaid_month = 1
	auto_renew = 1
	description_num = true
}

data "cds_data_source_dedicated_host" "test" {
	host_id = "xxx"
}
