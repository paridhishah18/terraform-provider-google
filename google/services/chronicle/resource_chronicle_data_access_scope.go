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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/chronicle/DataAccessScope.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package chronicle

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceChronicleDataAccessScope() *schema.Resource {
	return &schema.Resource{
		Create: resourceChronicleDataAccessScopeCreate,
		Read:   resourceChronicleDataAccessScopeRead,
		Update: resourceChronicleDataAccessScopeUpdate,
		Delete: resourceChronicleDataAccessScopeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceChronicleDataAccessScopeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"data_access_scope_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Required. The user provided scope id which will become the last part of the name
of the scope resource.
Needs to be compliant with https://google.aip.dev/122`,
			},
			"instance": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The unique identifier for the Chronicle instance, which is the same as the customer ID.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or "europe-west2".`,
			},
			"allow_all": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Optional. Whether or not the scope allows all labels, allow_all and
allowed_data_access_labels are mutually exclusive and one of them must be
present. denied_data_access_labels can still be used along with allow_all.
When combined with denied_data_access_labels, access will be granted to all
data that doesn't have labels mentioned in denied_data_access_labels. E.g.:
A customer with scope with denied labels A and B and allow_all will be able
to see all data except data labeled with A and data labeled with B and data
with labels A and B.`,
			},
			"allowed_data_access_labels": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The allowed labels for the scope. There has to be at
least one label allowed for the scope to be valid.
The logical operator for evaluation of the allowed labels is OR.
Either allow_all or allowed_data_access_labels needs to be provided.
E.g.: A customer with scope with allowed labels A and B will be able
to see data with labeled with A or B or (A and B).`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"asset_namespace": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The asset namespace configured in the forwarder
of the customer's events.`,
						},
						"data_access_label": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The name of the data access label.`,
						},
						"ingestion_label": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Representation of an ingestion label type.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ingestion_label_key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Required. The key of the ingestion label. Always required.`,
									},
									"ingestion_label_value": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `Optional. The value of the ingestion label. Optional. An object
with no provided value and some key provided would match
against the given key and ANY value.`,
									},
								},
							},
						},
						"log_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The name of the log type.`,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `Output only. The display name of the label.
Data access label and log types's name
will match the display name of the resource.
The asset namespace will match the namespace itself.
The ingestion key value pair will match the key of the tuple.`,
						},
					},
				},
				AtLeastOneOf: []string{"allowed_data_access_labels", "allow_all"},
			},
			"denied_data_access_labels": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Optional. The denied labels for the scope.
The logical operator for evaluation of the denied labels is AND.
E.g.: A customer with scope with denied labels A and B won't be able
to see data labeled with A and data labeled with B
and data with labels A and B.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"asset_namespace": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The asset namespace configured in the forwarder
of the customer's events.`,
						},
						"data_access_label": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The name of the data access label.`,
						},
						"ingestion_label": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Representation of an ingestion label type.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ingestion_label_key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Required. The key of the ingestion label. Always required.`,
									},
									"ingestion_label_value": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `Optional. The value of the ingestion label. Optional. An object
with no provided value and some key provided would match
against the given key and ANY value.`,
									},
								},
							},
						},
						"log_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The name of the log type.`,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `Output only. The display name of the label.
Data access label and log types's name
will match the display name of the resource.
The asset namespace will match the namespace itself.
The ingestion key value pair will match the key of the tuple.`,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Optional. A description of the data access scope for a human reader.`,
			},
			"author": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The user who created the data access scope.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time at which the data access scope was created.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The name to be used for display to customers of the data access scope.`,
			},
			"last_editor": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The user who last updated the data access scope.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique full name of the data access scope. This unique identifier is generated using values provided for the URL parameters.
Format:
projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{data_access_scope_id}`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time at which the data access scope was last updated.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceChronicleDataAccessScopeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	allowedDataAccessLabelsProp, err := expandChronicleDataAccessScopeAllowedDataAccessLabels(d.Get("allowed_data_access_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allowed_data_access_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(allowedDataAccessLabelsProp)) && (ok || !reflect.DeepEqual(v, allowedDataAccessLabelsProp)) {
		obj["allowedDataAccessLabels"] = allowedDataAccessLabelsProp
	}
	allowAllProp, err := expandChronicleDataAccessScopeAllowAll(d.Get("allow_all"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allow_all"); !tpgresource.IsEmptyValue(reflect.ValueOf(allowAllProp)) && (ok || !reflect.DeepEqual(v, allowAllProp)) {
		obj["allowAll"] = allowAllProp
	}
	deniedDataAccessLabelsProp, err := expandChronicleDataAccessScopeDeniedDataAccessLabels(d.Get("denied_data_access_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("denied_data_access_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(deniedDataAccessLabelsProp)) && (ok || !reflect.DeepEqual(v, deniedDataAccessLabelsProp)) {
		obj["deniedDataAccessLabels"] = deniedDataAccessLabelsProp
	}
	descriptionProp, err := expandChronicleDataAccessScopeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ChronicleBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance}}/dataAccessScopes?dataAccessScopeId={{data_access_scope_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DataAccessScope: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DataAccessScope: %s", err)
	}
	billingProject = project

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
		return fmt.Errorf("Error creating DataAccessScope: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/instances/{{instance}}/dataAccessScopes/{{data_access_scope_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating DataAccessScope %q: %#v", d.Id(), res)

	return resourceChronicleDataAccessScopeRead(d, meta)
}

func resourceChronicleDataAccessScopeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ChronicleBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance}}/dataAccessScopes/{{data_access_scope_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DataAccessScope: %s", err)
	}
	billingProject = project

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ChronicleDataAccessScope %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DataAccessScope: %s", err)
	}

	if err := d.Set("name", flattenChronicleDataAccessScopeName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataAccessScope: %s", err)
	}
	if err := d.Set("allowed_data_access_labels", flattenChronicleDataAccessScopeAllowedDataAccessLabels(res["allowedDataAccessLabels"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataAccessScope: %s", err)
	}
	if err := d.Set("allow_all", flattenChronicleDataAccessScopeAllowAll(res["allowAll"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataAccessScope: %s", err)
	}
	if err := d.Set("denied_data_access_labels", flattenChronicleDataAccessScopeDeniedDataAccessLabels(res["deniedDataAccessLabels"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataAccessScope: %s", err)
	}
	if err := d.Set("display_name", flattenChronicleDataAccessScopeDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataAccessScope: %s", err)
	}
	if err := d.Set("create_time", flattenChronicleDataAccessScopeCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataAccessScope: %s", err)
	}
	if err := d.Set("author", flattenChronicleDataAccessScopeAuthor(res["author"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataAccessScope: %s", err)
	}
	if err := d.Set("last_editor", flattenChronicleDataAccessScopeLastEditor(res["lastEditor"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataAccessScope: %s", err)
	}
	if err := d.Set("description", flattenChronicleDataAccessScopeDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataAccessScope: %s", err)
	}
	if err := d.Set("update_time", flattenChronicleDataAccessScopeUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading DataAccessScope: %s", err)
	}

	return nil
}

func resourceChronicleDataAccessScopeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DataAccessScope: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	allowedDataAccessLabelsProp, err := expandChronicleDataAccessScopeAllowedDataAccessLabels(d.Get("allowed_data_access_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allowed_data_access_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, allowedDataAccessLabelsProp)) {
		obj["allowedDataAccessLabels"] = allowedDataAccessLabelsProp
	}
	allowAllProp, err := expandChronicleDataAccessScopeAllowAll(d.Get("allow_all"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allow_all"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, allowAllProp)) {
		obj["allowAll"] = allowAllProp
	}
	deniedDataAccessLabelsProp, err := expandChronicleDataAccessScopeDeniedDataAccessLabels(d.Get("denied_data_access_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("denied_data_access_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, deniedDataAccessLabelsProp)) {
		obj["deniedDataAccessLabels"] = deniedDataAccessLabelsProp
	}
	descriptionProp, err := expandChronicleDataAccessScopeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ChronicleBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance}}/dataAccessScopes/{{data_access_scope_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DataAccessScope %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("allowed_data_access_labels") {
		updateMask = append(updateMask, "allowedDataAccessLabels")
	}

	if d.HasChange("allow_all") {
		updateMask = append(updateMask, "allowAll")
	}

	if d.HasChange("denied_data_access_labels") {
		updateMask = append(updateMask, "deniedDataAccessLabels")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
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
			return fmt.Errorf("Error updating DataAccessScope %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating DataAccessScope %q: %#v", d.Id(), res)
		}

	}

	return resourceChronicleDataAccessScopeRead(d, meta)
}

func resourceChronicleDataAccessScopeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DataAccessScope: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ChronicleBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance}}/dataAccessScopes/{{data_access_scope_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting DataAccessScope %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "DataAccessScope")
	}

	log.Printf("[DEBUG] Finished deleting DataAccessScope %q: %#v", d.Id(), res)
	return nil
}

func resourceChronicleDataAccessScopeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/instances/(?P<instance>[^/]+)/dataAccessScopes/(?P<data_access_scope_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<instance>[^/]+)/(?P<data_access_scope_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<instance>[^/]+)/(?P<data_access_scope_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/instances/{{instance}}/dataAccessScopes/{{data_access_scope_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenChronicleDataAccessScopeName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeAllowedDataAccessLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"data_access_label": flattenChronicleDataAccessScopeAllowedDataAccessLabelsDataAccessLabel(original["dataAccessLabel"], d, config),
			"log_type":          flattenChronicleDataAccessScopeAllowedDataAccessLabelsLogType(original["logType"], d, config),
			"asset_namespace":   flattenChronicleDataAccessScopeAllowedDataAccessLabelsAssetNamespace(original["assetNamespace"], d, config),
			"ingestion_label":   flattenChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabel(original["ingestionLabel"], d, config),
			"display_name":      flattenChronicleDataAccessScopeAllowedDataAccessLabelsDisplayName(original["displayName"], d, config),
		})
	}
	return transformed
}
func flattenChronicleDataAccessScopeAllowedDataAccessLabelsDataAccessLabel(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeAllowedDataAccessLabelsLogType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeAllowedDataAccessLabelsAssetNamespace(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabel(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["ingestion_label_key"] =
		flattenChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabelIngestionLabelKey(original["ingestionLabelKey"], d, config)
	transformed["ingestion_label_value"] =
		flattenChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabelIngestionLabelValue(original["ingestionLabelValue"], d, config)
	return []interface{}{transformed}
}
func flattenChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabelIngestionLabelKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabelIngestionLabelValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeAllowedDataAccessLabelsDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeAllowAll(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeDeniedDataAccessLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"display_name":      flattenChronicleDataAccessScopeDeniedDataAccessLabelsDisplayName(original["displayName"], d, config),
			"data_access_label": flattenChronicleDataAccessScopeDeniedDataAccessLabelsDataAccessLabel(original["dataAccessLabel"], d, config),
			"log_type":          flattenChronicleDataAccessScopeDeniedDataAccessLabelsLogType(original["logType"], d, config),
			"asset_namespace":   flattenChronicleDataAccessScopeDeniedDataAccessLabelsAssetNamespace(original["assetNamespace"], d, config),
			"ingestion_label":   flattenChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel(original["ingestionLabel"], d, config),
		})
	}
	return transformed
}
func flattenChronicleDataAccessScopeDeniedDataAccessLabelsDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeDeniedDataAccessLabelsDataAccessLabel(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeDeniedDataAccessLabelsLogType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeDeniedDataAccessLabelsAssetNamespace(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["ingestion_label_key"] =
		flattenChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelIngestionLabelKey(original["ingestionLabelKey"], d, config)
	transformed["ingestion_label_value"] =
		flattenChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelIngestionLabelValue(original["ingestionLabelValue"], d, config)
	return []interface{}{transformed}
}
func flattenChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelIngestionLabelKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelIngestionLabelValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeAuthor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeLastEditor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleDataAccessScopeUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandChronicleDataAccessScopeAllowedDataAccessLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDataAccessLabel, err := expandChronicleDataAccessScopeAllowedDataAccessLabelsDataAccessLabel(original["data_access_label"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDataAccessLabel); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["dataAccessLabel"] = transformedDataAccessLabel
		}

		transformedLogType, err := expandChronicleDataAccessScopeAllowedDataAccessLabelsLogType(original["log_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLogType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["logType"] = transformedLogType
		}

		transformedAssetNamespace, err := expandChronicleDataAccessScopeAllowedDataAccessLabelsAssetNamespace(original["asset_namespace"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAssetNamespace); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["assetNamespace"] = transformedAssetNamespace
		}

		transformedIngestionLabel, err := expandChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabel(original["ingestion_label"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIngestionLabel); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ingestionLabel"] = transformedIngestionLabel
		}

		transformedDisplayName, err := expandChronicleDataAccessScopeAllowedDataAccessLabelsDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandChronicleDataAccessScopeAllowedDataAccessLabelsDataAccessLabel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeAllowedDataAccessLabelsLogType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeAllowedDataAccessLabelsAssetNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIngestionLabelKey, err := expandChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabelIngestionLabelKey(original["ingestion_label_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngestionLabelKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ingestionLabelKey"] = transformedIngestionLabelKey
	}

	transformedIngestionLabelValue, err := expandChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabelIngestionLabelValue(original["ingestion_label_value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngestionLabelValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ingestionLabelValue"] = transformedIngestionLabelValue
	}

	return transformed, nil
}

func expandChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabelIngestionLabelKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabelIngestionLabelValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeAllowedDataAccessLabelsDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeAllowAll(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeDeniedDataAccessLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandChronicleDataAccessScopeDeniedDataAccessLabelsDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		transformedDataAccessLabel, err := expandChronicleDataAccessScopeDeniedDataAccessLabelsDataAccessLabel(original["data_access_label"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDataAccessLabel); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["dataAccessLabel"] = transformedDataAccessLabel
		}

		transformedLogType, err := expandChronicleDataAccessScopeDeniedDataAccessLabelsLogType(original["log_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLogType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["logType"] = transformedLogType
		}

		transformedAssetNamespace, err := expandChronicleDataAccessScopeDeniedDataAccessLabelsAssetNamespace(original["asset_namespace"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAssetNamespace); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["assetNamespace"] = transformedAssetNamespace
		}

		transformedIngestionLabel, err := expandChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel(original["ingestion_label"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIngestionLabel); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ingestionLabel"] = transformedIngestionLabel
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandChronicleDataAccessScopeDeniedDataAccessLabelsDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeDeniedDataAccessLabelsDataAccessLabel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeDeniedDataAccessLabelsLogType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeDeniedDataAccessLabelsAssetNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIngestionLabelKey, err := expandChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelIngestionLabelKey(original["ingestion_label_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngestionLabelKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ingestionLabelKey"] = transformedIngestionLabelKey
	}

	transformedIngestionLabelValue, err := expandChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelIngestionLabelValue(original["ingestion_label_value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngestionLabelValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ingestionLabelValue"] = transformedIngestionLabelValue
	}

	return transformed, nil
}

func expandChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelIngestionLabelKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelIngestionLabelValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleDataAccessScopeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
