---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cds_security_group Resource - terraform-provider-cds"
subcategory: ""
description: |-
  Security Group.

---

# cds_security_group (Resource)

Security Group. 

## Example usage

```hcl

resource "cds_security_group" "security_group_2" {
  name        = "test_tf_new_zz25"
  description = "New security group25"
  type        = "private"
  rule {
    action        = "1"
    description   = "tf_rule_test"
    targetaddress = "120.78.170.188/28;120.78.170.188/28;120.78.170.188/28"
    targetport    = "70;90;8"
    localport     = "800"
    direction     = "all"
    priority      = "11"
    protocol      = "TCP"
    ruletype      = "ip"
  }
}

```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) Description of the security group.
- `name` (String) Name of the security group.
- `type` (String) Type of the security group.

### Optional

- `rule` (Block Set, Max: 15) Security group rule. [View Document](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1createsecuritygroup) (see [below for nested schema](#nestedblock--rule))

### Read-Only

- `id` (String) The ID of this resource.
- `rule_current` (Set of Object) Current rules. (see [below for nested schema](#nestedatt--rule_current))

<a id="nestedblock--rule"></a>
### Nested Schema for `rule`

Optional:

- `action` (String) Action. Network connectivity policy within a security group. 0 represents deny, 1 represents allow. [View Document](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1createsecuritygroup)
- `description` (String) Description of the security group rule. Less than 256 characters. [View Document](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1createsecuritygroup)
- `direction` (String) Description of the security group rule.
- `id` (String) Id of the rule.
- `localport` (String) Local port. The source port that is open or blocked on this machine. [View Document](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1createsecuritygroup)
- `priority` (String) Priority. Set the rule priority in the security group, ranging from 1 to 100. [View Document](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1createsecuritygroup)
- `protocol` (String) Protocol. Transport layer protocol. Optional parameters are: ICMP, TCP, UDP. Note: If the protocol is ICMP, there is no need to pass the TargetPort and LocalPort parameters. [View Document](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1createsecuritygroup)
- `ruletype` (String) Rule type. Set rule type. Options: mac/ip.
- `targetaddress` (String) Target address. Less than 200 characters. [View Document](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1createsecuritygroup)
- `targetport` (String) Target port. [View Document](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#1createsecuritygroup)


<a id="nestedatt--rule_current"></a>
### Nested Schema for `rule_current`

Read-Only:

- `action` (String)
- `description` (String)
- `direction` (String)
- `id` (String)
- `localport` (String)
- `priority` (String)
- `protocol` (String)
- `ruletype` (String)
- `targetaddress` (String)
- `targetport` (String)
