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
resource "cds_haproxy" "haproxy_example" {
    instance_uuid       = var.instance_uuid
    instance_name       = var.instance_name
    region_id           = var.region_id
    vdc_id              = var.vdc_id
    base_pipe_id        = var.base_pipe_id
    paas_goods_id       = var.paas_goods_id
    ips = [{
        pipe_type  = var.pipe_type
        pipe_id    = var.pipe_id
        segment_id = var.segment.id
    }]
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
	}]
}

data cds_data_source_haproxy "my_haproxy_data" {
    ip = ""
    instance_name = ""
    start_time = ""
    end_time = ""
    region_id = ""
}

data cds_data_source_certificate "my_certificate" {
    result_output_file = data.json
}
```

## Argument Reference

The following arguments are supported:

# Haproxy
* `ip` - (Optional) The HaProxy instance IP to filter.
* `instance_name` - (Optional) The HaProxy instance name to filter.
* `start_time` - (Optional) The HaProxy instance create time to filter.
* `end_time` - (Optional) The HaProxy instance end time to filter.
* `result_output_file` - (Required) Save all instance information under the VDC to the path.

# Certificate
* no parameter
