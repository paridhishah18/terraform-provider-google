// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycenter/FolderNotificationConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securitycenter

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceSecurityCenterFolderNotificationConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityCenterFolderNotificationConfigCreate,
		Read:   resourceSecurityCenterFolderNotificationConfigRead,
		Update: resourceSecurityCenterFolderNotificationConfigUpdate,
		Delete: resourceSecurityCenterFolderNotificationConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecurityCenterFolderNotificationConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"config_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `This must be unique within the organization.`,
			},
			"folder": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Numerical ID of the parent folder.`,
			},
			"pubsub_topic": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The Pub/Sub topic to send notifications to. Its format is
"projects/[project_id]/topics/[topic]".`,
			},
			"streaming_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `The config for triggering streaming-based notifications.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Expression that defines the filter to apply across create/update
events of assets or findings as specified by the event type. The
expression is a list of zero or more restrictions combined via
logical operators AND and OR. Parentheses are supported, and OR
has higher precedence than AND.

Restrictions have the form <field> <operator> <value> and may have
a - character in front of them to indicate negation. The fields
map to those defined in the corresponding resource.

The supported operators are:

* = for all value types.
* >, <, >=, <= for integer values.
* :, meaning substring matching, for strings.

The supported value types are:

* string literals in quotes.
* integer literals without quotes.
* boolean literals true and false without quotes.

See
[Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications)
for information on how to write a filter.`,
						},
					},
				},
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 1024),
				Description:  `The description of the notification config (max of 1024 characters).`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of this notification config, in the format
'folders/{{folder}}/notificationConfigs/{{config_id}}'.`,
			},
			"service_account": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The service account that needs "pubsub.topics.publish" permission to
publish to the Pub/Sub topic.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSecurityCenterFolderNotificationConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterFolderNotificationConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	pubsubTopicProp, err := expandSecurityCenterFolderNotificationConfigPubsubTopic(d.Get("pubsub_topic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("pubsub_topic"); !tpgresource.IsEmptyValue(reflect.ValueOf(pubsubTopicProp)) && (ok || !reflect.DeepEqual(v, pubsubTopicProp)) {
		obj["pubsubTopic"] = pubsubTopicProp
	}
	streamingConfigProp, err := expandSecurityCenterFolderNotificationConfigStreamingConfig(d.Get("streaming_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("streaming_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(streamingConfigProp)) && (ok || !reflect.DeepEqual(v, streamingConfigProp)) {
		obj["streamingConfig"] = streamingConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterBasePath}}folders/{{folder}}/notificationConfigs?configId={{config_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FolderNotificationConfig: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating FolderNotificationConfig: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "folders/{{folder}}/notificationConfigs/{{config_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FolderNotificationConfig %q: %#v", d.Id(), res)

	return resourceSecurityCenterFolderNotificationConfigRead(d, meta)
}

func resourceSecurityCenterFolderNotificationConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterBasePath}}folders/{{folder}}/notificationConfigs/{{config_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecurityCenterFolderNotificationConfig %q", d.Id()))
	}

	if err := d.Set("name", flattenSecurityCenterFolderNotificationConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderNotificationConfig: %s", err)
	}
	if err := d.Set("description", flattenSecurityCenterFolderNotificationConfigDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderNotificationConfig: %s", err)
	}
	if err := d.Set("pubsub_topic", flattenSecurityCenterFolderNotificationConfigPubsubTopic(res["pubsubTopic"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderNotificationConfig: %s", err)
	}
	if err := d.Set("service_account", flattenSecurityCenterFolderNotificationConfigServiceAccount(res["serviceAccount"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderNotificationConfig: %s", err)
	}
	if err := d.Set("streaming_config", flattenSecurityCenterFolderNotificationConfigStreamingConfig(res["streamingConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderNotificationConfig: %s", err)
	}

	return nil
}

func resourceSecurityCenterFolderNotificationConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterFolderNotificationConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	pubsubTopicProp, err := expandSecurityCenterFolderNotificationConfigPubsubTopic(d.Get("pubsub_topic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("pubsub_topic"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, pubsubTopicProp)) {
		obj["pubsubTopic"] = pubsubTopicProp
	}
	streamingConfigProp, err := expandSecurityCenterFolderNotificationConfigStreamingConfig(d.Get("streaming_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("streaming_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, streamingConfigProp)) {
		obj["streamingConfig"] = streamingConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterBasePath}}folders/{{folder}}/notificationConfigs/{{config_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FolderNotificationConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("pubsub_topic") {
		updateMask = append(updateMask, "pubsubTopic")
	}

	if d.HasChange("streaming_config") {
		updateMask = append(updateMask, "streamingConfig.filter")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating FolderNotificationConfig %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating FolderNotificationConfig %q: %#v", d.Id(), res)
		}

	}

	return resourceSecurityCenterFolderNotificationConfigRead(d, meta)
}

func resourceSecurityCenterFolderNotificationConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterBasePath}}folders/{{folder}}/notificationConfigs/{{config_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting FolderNotificationConfig %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "FolderNotificationConfig")
	}

	log.Printf("[DEBUG] Finished deleting FolderNotificationConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceSecurityCenterFolderNotificationConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^folders/(?P<folder>[^/]+)/notificationConfigs/(?P<config_id>[^/]+)$",
		"^(?P<folder>[^/]+)/(?P<config_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "folders/{{folder}}/notificationConfigs/{{config_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	idParts := strings.Split(d.Id(), "/")
	if len(idParts) != 4 {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected folders/{{folder}}/notificationConfigs/{{config_id}}", d.Id())
	}

	if err := d.Set("folder", idParts[1]); err != nil {
		return nil, fmt.Errorf("error setting folder: %s", err)
	}

	if err := d.Set("config_id", idParts[3]); err != nil {
		return nil, fmt.Errorf("error setting config_id: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenSecurityCenterFolderNotificationConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderNotificationConfigDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderNotificationConfigPubsubTopic(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderNotificationConfigServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterFolderNotificationConfigStreamingConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["filter"] =
		flattenSecurityCenterFolderNotificationConfigStreamingConfigFilter(original["filter"], d, config)
	return []interface{}{transformed}
}
func flattenSecurityCenterFolderNotificationConfigStreamingConfigFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecurityCenterFolderNotificationConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderNotificationConfigPubsubTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderNotificationConfigStreamingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFilter, err := expandSecurityCenterFolderNotificationConfigStreamingConfigFilter(original["filter"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilter); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["filter"] = transformedFilter
	}

	return transformed, nil
}

func expandSecurityCenterFolderNotificationConfigStreamingConfigFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
