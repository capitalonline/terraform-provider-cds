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
