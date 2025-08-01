---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/iamworkforcepool/WorkforcePool.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/datasource_iam.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud IAM"
description: |-
  A datasource to retrieve the IAM policy state for Cloud IAM WorkforcePool
---


# google_iam_workforce_pool_iam_policy

Retrieves the current IAM policy data for workforcepool


## Example Usage


```hcl
data "google_iam_workforce_pool_iam_policy" "policy" {
  location = google_iam_workforce_pool.example.location
  workforce_pool_id = google_iam_workforce_pool.example.workforce_pool_id
}
```

## Argument Reference

The following arguments are supported:

* `location` - (Optional) The location for the resource. Used to find the parent resource to bind the IAM policy to. If not specified,
  the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
  location is specified, it is taken from the provider configuration.
* `workforce_pool_id` - (Required) Used to find the parent resource to bind the IAM policy to

## Attributes Reference

The attributes are exported:

* `etag` - (Computed) The etag of the IAM policy.

* `policy_data` - (Required only by `google_iam_workforce_pool_iam_policy`) The policy data generated by
  a `google_iam_policy` data source.
