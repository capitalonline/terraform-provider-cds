---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cds_data_source_subjects Data Source - terraform-provider-cds"
subcategory: ""
description: |-
  Data source subjects
  Example usage
  ```hcl
  data "cdsdatasource_subjects" "subjects" {
  }
  ```
---

# cds_data_source_subjects (Data Source)

Data source subjects 

## Example usage

```hcl

data "cds_data_source_subjects" "subjects" {
}

```



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `result_output_file` (String) Output file path.

### Read-Only

- `id` (String) The ID of this resource.
- `subjects` (List of Object) Subjects. (see [below for nested schema](#nestedatt--subjects))

<a id="nestedatt--subjects"></a>
### Nested Schema for `subjects`

Read-Only:

- `balance` (String)
- `begin_time` (String)
- `bill_method` (String)
- `bill_method_display` (String)
- `end_time` (String)
- `goods_ids` (String)
- `goods_names` (String)
- `name` (String)
- `site_ids` (String)
- `site_names` (String)
- `subject_id` (Number)