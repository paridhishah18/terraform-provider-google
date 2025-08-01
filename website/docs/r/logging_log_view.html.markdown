---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/logging/LogView.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud (Stackdriver) Logging"
description: |-
  Describes a view over log entries in a bucket.
---

# google_logging_log_view

Describes a view over log entries in a bucket.


To get more information about LogView, see:

* [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.locations.buckets.views)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/logging/docs/apis)

## Example Usage - Logging Log View Basic


```hcl
resource "google_logging_project_bucket_config" "logging_log_view" {
    project        = "my-project-name"
    location       = "global"
    retention_days = 30
    bucket_id      = "_Default"
}

resource "google_logging_log_view" "logging_log_view" {
  name        = "my-view"
  bucket      = google_logging_project_bucket_config.logging_log_view.id
  description = "A logging view configured with Terraform"
  filter      = "SOURCE(\"projects/myproject\") AND resource.type = \"gce_instance\" AND LOG_ID(\"stdout\")"
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  The resource name of the view. For example: \`projects/my-project/locations/global/buckets/my-bucket/views/my-view\`

* `bucket` -
  (Required)
  The bucket of the resource


* `description` -
  (Optional)
  Describes this view.

* `filter` -
  (Optional)
  Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: - originating project/folder/organization/billing account. - resource type - log id For example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")

* `parent` -
  (Optional)
  The parent of the resource.

* `location` -
  (Optional)
  The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}`

* `create_time` -
  Output only. The creation timestamp of the view.

* `update_time` -
  Output only. The last update timestamp of the view.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


LogView can be imported using any of these accepted formats:

* `{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import LogView using one of the formats above. For example:

```tf
import {
  id = "{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}"
  to = google_logging_log_view.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), LogView can be imported using one of the formats above. For example:

```
$ terraform import google_logging_log_view.default {{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}
```
