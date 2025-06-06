---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/d/pubsub_subscription_iam_policy.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud Pub/Sub"
description: |-
  A datasource to retrieve the IAM policy state for a Pubsub subscription.
---


# `google_pubsub_subscription_iam_policy`
Retrieves the current IAM policy data for a Pubsub subscription.

## example

```hcl
data "google_pubsub_subscription_iam_policy" "policy" {
  subscription = google_pubsub_subscription.subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `subscription` - (Required) The subscription name or id to bind to attach IAM policy to.

* `project` - (Optional) The project in which the resource belongs. If it
    is not provided, the provider project is used.

## Attributes Reference

The attributes are exported:

* `etag` - (Computed) The etag of the IAM policy.

* `policy_data` - (Computed) The policy data
