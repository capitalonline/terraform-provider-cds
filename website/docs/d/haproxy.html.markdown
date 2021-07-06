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
    strategies = [{
        http_listeners = [{
            acl_white_list = ""
            backend_server = [{
                ip = ""
                max_conn = 1000
                port = port
                weight = 1
            }]
            certificate_ids = [{
                certificate_id = ""
                certificate_name = ""
            }]
            client_timeout = ""
            client_timeout_unit = ""
            connect_timeout = ""
            connect_timeout_unit = ""
            server_timeout = ""
            server_timeout_unit
            listener_mode = ""
            listener_name = ""
            listener_port = 8080
            max_conn = 2000
            scheduler = ""
            sticky_session = ""
        }]
        tcp_listeners = [{
            acl_white_list = ""
            backend_server = [{
                ip = ""
                max_conn = 1000
                port = port
                weight = 1
            }]
            client_timeout = ""
            client_timeout_unit = ""
            connect_timeout = ""
            connect_timeout_unit = ""
            server_timeout = ""
            server_timeout_unit
            listener_mode = ""
            listener_name = ""
            listener_port = 8080
            max_conn = 2000
            scheduler = ""
        }] 
    }] 
}

data cds_data_source_haproxy "my_haproxy_data" {
    ip = ""
    instance_name = ""
    start_time = ""
    end_time = ""
    region_id = ""
}
```

## Argument Reference

The following arguments are supported:

* `ip` - (Optional) The HaProxy instance IP to filter.
* `instance_name` - (Optional) The HaProxy instance name to filter.
* `start_time` - (Optional) The HaProxy instance create time to filter.
* `end_time` - (Optional) The HaProxy instance end time to filter.
* `result_output_file` - (Required) Save all instance information under the VDC to the path.