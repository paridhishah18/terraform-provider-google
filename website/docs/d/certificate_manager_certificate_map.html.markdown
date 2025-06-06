---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/d/certificate_manager_certificate_map.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Certificate Manager"
description: |-
  Contains the data that describes a Certificate Map
---
# google_certificate_manager_certificate_map

Get info about a Google Certificate Manager Certificate Map resource.

## Example Usage

```tf
data "google_certificate_manager_certificate_map" "default" {
 name = "cert-map"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the certificate map.

- - -
* `project` - (Optional) The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.

## Attributes Reference

See [google_certificate_manager_certificate_map](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/certificate_manager_certificate_map) resource for details of the available attributes.
