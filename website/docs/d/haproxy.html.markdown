---
layout: "cds"
page_title: "cds: haproxy"
sidebar_current: "docs-cds-resource-haproxy"
description: |-
  Provide resources to create or delete HaProxy.
---

# HaProxy

This data source provides a json file of instances in a HaProxy  

## Example Usage

```hcl

data cds_data_source_haproxy "my_haproxy_data" {
    ip = ""
    instance_uuid = ""
    instance_name = ""
    start_time = ""
    end_time = ""
    region_id = ""
}

data cds_data_source_certificate "my_certificate" {
    result_output_file = data.json
}

output "test" {
  # output or use my_haproxy_data's ip
  value = data.cds_data_source_haproxy.my_haproxy_data.ip
}
```

## Argument Reference

The following arguments are supported:

# Haproxy
* `ip` - (Optional) The HaProxy instance IP to filter.
* `instance_uuid` - (Optional) The HaProxy instance id to filter.
* `instance_name` - (Optional) The HaProxy instance name to filter.
* `start_time` - (Optional) The HaProxy instance create time to filter.
* `end_time` - (Optional) The HaProxy instance end time to filter.
* `result_output_file` - (Required) Save all instance information under the VDC to the path.
* `ha_list` - (Computed) The list of haproxy info , By calculation , Don't have to fill out.
* `region_id` (Required) The region of HaProxy.
* `ha_list` - (Computed) List of HaProxy instances of the query .
  * `cpu` - (Computed) The cpu number of HaProxy.
  * `create_time` - (Computed) The created time  of HaProxy.
  * `display_name` - (Computed) The site name.
  * `ip` - (Computed) The ip of HaProxy.
  * `instance_name` - (Computed) HaProxy instance name.
  * `instance_uuid` - (Computed) HaProxy instance uuid.
  * `link_type` - (Computed) A maximum number of disks can be created at a time.
  * `link_type_str` - (Computed) Link Type Name.
  * `master_info` - (Computed) In a slave cluster, read-only services have a value.
  * `port` - (Computed) Port to connect.
  * `ram` - (Computed) The ram of HaProxy.
  * `region_id` - (Computed) The region_id of HaProxy.
  * `resource_id` - (Computed) Bill id, unique identifier for querying a bill.
  * `status` - (Computed) Instance status.
  * `status_str` - (Computed) Indicates the Chinese version of the status.
  * `sub_product_name` - (Computed)  The name of the sub-product.
  * `vdc_id` - (Computed) Data center ID.
  * `vdc_name` - (Computed) Data center name.
  * `version` - (Computed) The service version.
  * `project_name` - (Computed) The project name.
  * `vips` - (Computed) Load balancing multiple VIP addresses.

# Certificate
* `ha_cert_list` - (Computed) The list of the ha Certificate , By calculation , Don't have to fill out.
