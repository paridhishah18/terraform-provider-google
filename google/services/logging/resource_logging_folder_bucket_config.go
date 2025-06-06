// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/logging/resource_logging_folder_bucket_config.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package logging

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

var loggingFolderBucketConfigSchema = map[string]*schema.Schema{
	"folder": {
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: `The parent resource that contains the logging bucket.`,
	},
}

func folderBucketConfigID(d *schema.ResourceData, config *transport_tpg.Config) (string, error) {
	folder := d.Get("folder").(string)
	location := d.Get("location").(string)
	bucketID := d.Get("bucket_id").(string)

	if !strings.HasPrefix(folder, "folder") {
		folder = "folders/" + folder
	}

	id := fmt.Sprintf("%s/locations/%s/buckets/%s", folder, location, bucketID)
	return id, nil
}

// Create Logging Bucket config
func ResourceLoggingFolderBucketConfig() *schema.Resource {
	return ResourceLoggingBucketConfig("folder", loggingFolderBucketConfigSchema, folderBucketConfigID)
}
