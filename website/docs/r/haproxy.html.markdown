---
layout: "cds"
page_title: "cds: haproxy"
sidebar_current: "docs-cds-resource-haproxy"
description: |-
  Provide resources to create or delete HaProxys.
---

# HaProxy

Provide resources to create, update or delete Haproxys.

## Example Usage

```hcl
# create haproxy
resource "cds_haproxy" "haproxy_example" {
    instance_uuid       = var.instance_uuid
    instance_name       = var.instance_name
    region_id           = var.region_id
    vdc_id              = var.vdc_id
    base_pipe_id        = var.base_pipe_id
    cpu                 = 1
    ram                 = 2
    ips = [
      {
        pipe_type  = "private"
        #PrivateNetwork PrivateId from data source vdc
        pipe_id    = var.pipe_id
        segment_id = var.segment.id
      },
      #This parameter is required if you want to create a public network
      {
        pipe_type  = "public"
        #PublicNetwork PublicId from data source vdc
        pipe_id    = var.public_pipe_id
        segment_id = var.segment_id
      }
    ]
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
      server_timeout_unit = ""
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
      server_timeout_unit = ""
      listener_mode = ""
      listener_name = ""
      listener_port = 8080
      max_conn = 2000
      scheduler = ""
    }] 
}

# create certificate
resource cds_certificate my_cds_certificate {
  certificate_name = "my_cert"
  certificate = "XXXXXXXXXX"
  private_key = "XXXXXXXXXX"
}
```
## Argument Reference
The following arguments are supported
### Haproxy
* `instance_uuid` - (Optional) After creation, you need to provide ID manually to support update and deletion
* `instance_name` - (Required,Unmodifiable) The name of the instance.
* `region_id` - (Required,Unmodifiable) The Region of the instance, refer to [All Region](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#1describezones).
* `vdc_id` - (Required,Unmodifiable) Instance belongs to the virtual data center.
* `base_pipe_id` - (Required,Unmodifiable) Vdc private network id, the haproxy instance will create id by this [Get PipeId](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1describevdc)
* `ips` - (Required,Unmodifiable) The network used by haproxy [All Instance Type](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#%E4%B8%BB%E6%9C%BA%E7%B1%BB%E5%9E%8B).
  * `pipe_type` - (Required) The network of the haproxy type. The options are public and private.Public and private network information is required to create a public network.When creating a private network, only the private network information is required.
  * `pipe_id` - (Required) The netwrok of the haproxy id. [Get PipeId](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1describevdc)
  * `segment_id` - (Optional) When the haproxy type is public, it needs to be provided. [Get SegmentId](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1describevdc)
* `paas_goods_id` - (Required,Unmodifiable) Product ID that support haproxy in specific region [List of product id that support haproxy in specific regions](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#1describezones)
* `http_listeners` - (Optional) HTTP configuration list.
  * `acl_white_list` - (Required) White list setting, this is a string use , to split.
  * `backend_server` - (Required) Backend server configuration.
    * `ip` - (Required) Backend server IP address.
    * `port` - (Required) Backend server port.
    * `weight` - (Required) Backend server weight. 1-256
    * `max_conn` - (Required) Maximum number of backend server connections.
  * `certificate_ids` - (Optional) Bind certificate, set empty array without binding.
     * `certificate_id` - (Required) The certificate id.
     * `certificate_name` - (Required)  The certificate name.
  * `client_timeout` - (Required) Set the time for client connection timeout.
  * `client_timeout_unit` - (Required) Set the time unit for client connection timeout ['m', 'ms'].
  * `connect_timeout` - (Required) Set the time for client connection timeout.
  * `connect_timeout_unit` - (Required) Set the time unit for client connection timeout.
  * `server_timeout` - (Required) Set the time for server connection timeout.
  * `server_timeout_unit` - (Required) Set the time unit for server connection timeout.
  * `listener_mode` - (Required) Listener mode.
  * `listener_name` - (Required) The name of the listening strategy is used to generate the configuration file. The name cannot be the same.
  * `listener_port` - (Required) Listener port.
  * `max_conn` - (Required) The maximum number of connections on the proxy side.
  * `scheduler` - (Required) Scheduling strategy [roundrobin, leastconn, static-rr, source]
  * `sticky_session` - (Required) Turn on call back hold [on, off]
  * `session_persistence` - (Optional) The persistence settings of the session.
    * `key` - (Optional) The key of Cookie.
    * `mode` - (Optional) The mode of the session persistence. [details](https://github.com/capitalonline/openapi/blob/master/%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%E6%A6%82%E8%A7%88.md#13modifyloadbalancerstrategys)
    * `timer` - (Optional) Set the duration of the session persistence.
      * `max_idle` - (Optional) If the hold time is exceeded and there is no new request in the connection, the session persistence is automatically broken.The default value is 0, indicating that the session is maintained even if the session is idle. In range of [0-7200]
      * `max_life` - (Optional) Set the maximum duration for which a session can be held.
  * `option` - (Optional)  Advanced configuration.
    * `httpchk` - (Optional) The settings of health check.
      * `method` - (Optional) The method for health check. Default value is 'GET'. Valid values ["GET","HEAD","OPTIONS"]
      * `uri` - (uri) The uri for health check.
* `tcp_listeners` - (Optional) TCP configuration list.
  * `acl_white_list` - (Required) White list setting, this is a string use , to split.
  * `backend_server` - (Required) Backend server configuration.
    * `ip` - (Required) Backend server IP address.
    * `port` - (Required) Backend server port.
    * `weight` - (Required) Backend server weight. 1-256
    * `max_conn` - (Required) Maximum number of backend server connections.
  * `client_timeout` - (Required) Set the time for client connection timeout.
  * `client_timeout_unit` - (Required) Set the time unit for client connection timeout ['m', 'ms'].
  * `connect_timeout` - (Required) Set the time for client connection timeout.
  * `connect_timeout_unit` - (Required) Set the time unit for client connection timeout.
  * `server_timeout` - (Required) Set the time for server connection timeout.
  * `server_timeout_unit` - (Required) Set the time unit for server connection timeout.
  * `listener_mode` - (Required) Listener mode.
  * `listener_name` - (Required) The name of the listening strategy is used to generate the configuration file. The name cannot be the same.
  * `listener_port` - (Required) Listener port.
  * `max_conn` - (Required) The maximum number of connections on the proxy side.
  * `scheduler` - (Required) Scheduling strategy [roundrobin, leastconn, static-rr, source]
### Certificate
* `certificate_name` - (Required) Certificate name
* `certificate` - (Required) Certificate content such as BEGIN CERTIFICATE-----\nXXXXXXXXX\n-----END CERTIFICATE-----\n(must contain \n)
* `private_key` - (Required) Private key content such as -----BEGIN RSA PRIVATE KEY-----\nXXXXXXXXX\n-----END RSA PRIVATE KEY-----\n(must contain \n)
