---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/d/container_cluster.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Kubernetes (Container) Engine"
description: |-
  Get info about a Google Kubernetes Engine cluster.
---

# google_container_cluster

Get info about a GKE cluster from its name and location.

## Example Usage

```tf
data "google_container_cluster" "my_cluster" {
  name     = "my-cluster"
  location = "us-east1-a"
}

output "endpoint" {
  value = data.google_container_cluster.my_cluster.endpoint
}

output "instance_group_urls" {
  value = data.google_container_cluster.my_cluster.node_pool[0].instance_group_urls
}

output "node_config" {
  value = data.google_container_cluster.my_cluster.node_config
}

output "node_pools" {
  value = data.google_container_cluster.my_cluster.node_pool
}
```

## Argument Reference

The following arguments are supported:

* `name` (Required) - The name of the cluster.

* `location` (Optional) - The location (zone or region) this cluster has been
created in. One of `location`, `region`, `zone`, or a provider-level `zone` must
be specified.

* `zone` (Optional) - The zone this cluster has been created in. Deprecated in
favour of `location`.

* `region` (Optional) - The region this cluster has been created in. Deprecated
in favour of `location`.

- - -

* `project` - (Optional) The project in which the resource belongs. If it
    is not provided, the provider project is used.

## Attributes Reference

See [google_container_cluster](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster) resource for details of the available attributes.
