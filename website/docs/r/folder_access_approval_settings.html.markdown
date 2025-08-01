---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/accessapproval/FolderSettings.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Access Approval"
description: |-
  Access Approval enables you to require your explicit approval whenever Google support and engineering need to access your customer content.
---

# google_folder_access_approval_settings

Access Approval enables you to require your explicit approval whenever Google support and engineering need to access your customer content.


To get more information about FolderSettings, see:

* [API documentation](https://cloud.google.com/access-approval/docs/reference/rest/v1/folders)

## Example Usage - Folder Access Approval Full


```hcl
resource "google_folder" "my_folder" {
  display_name = "my-folder"
  parent       = "organizations/123456789"
  deletion_protection = false
}

resource "google_folder_access_approval_settings" "folder_access_approval" {
  folder_id           = google_folder.my_folder.folder_id
  notification_emails = ["testuser@example.com", "example.user@example.com"]

  enrolled_services {
  	cloud_product = "all"
  }
}
```
## Example Usage - Folder Access Approval Active Key Version


```hcl
resource "google_folder" "my_folder" {
  display_name = "my-folder"
  parent       = "organizations/123456789"
  deletion_protection = false
}

resource "google_project" "my_project" {
  name       = "My Project"
  project_id = "your-project-id"
  folder_id  = google_folder.my_folder.name
  deletion_policy = "DELETE"
}

resource "google_kms_key_ring" "key_ring" {
  name     = "key-ring"
  location = "global"
  project  = google_project.my_project.project_id
}

resource "google_kms_crypto_key" "crypto_key" {
  name = "crypto-key"
  key_ring = google_kms_key_ring.key_ring.id
  purpose = "ASYMMETRIC_SIGN"

  version_template {
    algorithm = "EC_SIGN_P384_SHA384"
  }
}

data "google_access_approval_folder_service_account" "service_account" {
  folder_id = google_folder.my_folder.folder_id
}

resource "google_kms_crypto_key_iam_member" "iam" {
  crypto_key_id = google_kms_crypto_key.crypto_key.id
  role          = "roles/cloudkms.signerVerifier"
  member        = "serviceAccount:${data.google_access_approval_folder_service_account.service_account.account_email}"
}

data "google_kms_crypto_key_version" "crypto_key_version" {
  crypto_key = google_kms_crypto_key.crypto_key.id
}

resource "google_folder_access_approval_settings" "folder_access_approval" {
  folder_id           = google_folder.my_folder.folder_id
  active_key_version  = data.google_kms_crypto_key_version.crypto_key_version.name

  enrolled_services {
  	cloud_product = "all"
  }

  depends_on = [google_kms_crypto_key_iam_member.iam]
}
```

## Argument Reference

The following arguments are supported:


* `enrolled_services` -
  (Required)
  A list of Google Cloud Services for which the given resource has Access Approval enrolled.
  Access requests for the resource given by name against any of these services contained here will be required
  to have explicit approval. Enrollment can only be done on an all or nothing basis.
  A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
  Structure is [documented below](#nested_enrolled_services).

* `folder_id` -
  (Required)
  ID of the folder of the access approval settings.


* `notification_emails` -
  (Optional)
  A list of email addresses to which notifications relating to approval requests should be sent.
  Notifications relating to a resource will be sent to all emails in the settings of ancestor
  resources of that resource. A maximum of 50 email addresses are allowed.

* `active_key_version` -
  (Optional)
  The asymmetric crypto key version to use for signing approval requests.
  Empty active_key_version indicates that a Google-managed key should be used for signing.
  This property will be ignored if set by an ancestor of the resource, and new non-empty values may not be set.



<a name="nested_enrolled_services"></a>The `enrolled_services` block supports:

* `cloud_product` -
  (Required)
  The product for which Access Approval will be enrolled. Allowed values are listed (case-sensitive):
    * all
    * App Engine
    * BigQuery
    * Cloud Bigtable
    * Cloud Key Management Service
    * Compute Engine
    * Cloud Dataflow
    * Cloud Identity and Access Management
    * Cloud Pub/Sub
    * Cloud Storage
    * Persistent Disk
  Note: These values are supported as input, but considered a legacy format:
    * all
    * appengine.googleapis.com
    * bigquery.googleapis.com
    * bigtable.googleapis.com
    * cloudkms.googleapis.com
    * compute.googleapis.com
    * dataflow.googleapis.com
    * iam.googleapis.com
    * pubsub.googleapis.com
    * storage.googleapis.com

* `enrollment_level` -
  (Optional)
  The enrollment level of the service.
  Default value is `BLOCK_ALL`.
  Possible values are: `BLOCK_ALL`.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `folders/{{folder_id}}/accessApprovalSettings`

* `name` -
  The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"

* `enrolled_ancestor` -
  If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors of the Folder.

* `ancestor_has_active_key_version` -
  If the field is true, that indicates that an ancestor of this Folder has set active_key_version.

* `invalid_key_version` -
  If the field is true, that indicates that there is some configuration issue with the active_key_version
  configured on this Folder (e.g. it doesn't exist or the Access Approval service account doesn't have the
  correct permissions on it, etc.) This key version is not necessarily the effective key version at this level,
  as key versions are inherited top-down.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


FolderSettings can be imported using any of these accepted formats:

* `folders/{{folder_id}}/accessApprovalSettings`
* `{{folder_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import FolderSettings using one of the formats above. For example:

```tf
import {
  id = "folders/{{folder_id}}/accessApprovalSettings"
  to = google_folder_access_approval_settings.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), FolderSettings can be imported using one of the formats above. For example:

```
$ terraform import google_folder_access_approval_settings.default folders/{{folder_id}}/accessApprovalSettings
$ terraform import google_folder_access_approval_settings.default {{folder_id}}
```
