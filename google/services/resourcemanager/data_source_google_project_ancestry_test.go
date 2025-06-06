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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/resourcemanager/data_source_google_project_ancestry_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package resourcemanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDataSourceGoogleProjectAncestry_basic(t *testing.T) {
	t.Parallel()

	// Common resource configuration
	staticPrefix := "tf-test-"
	randomSuffix := "-" + acctest.RandString(t, 10)
	orgID := envvar.GetTestOrgFromEnv(t)

	// Configuration of resources
	folderThanos := staticPrefix + "thanos" + randomSuffix
	folderLoki := staticPrefix + "loki" + randomSuffix
	folderUltron := staticPrefix + "ultron" + randomSuffix
	projectThor := staticPrefix + "thor" + randomSuffix
	projectIronMan := staticPrefix + "ironman" + randomSuffix
	projectCap := staticPrefix + "cap" + randomSuffix
	projectHulk := staticPrefix + "hulk" + randomSuffix

	// Configuration map used in test deployment
	context := map[string]interface{}{
		"org_id":          orgID,
		"folder_thanos":   folderThanos,
		"folder_loki":     folderLoki,
		"folder_ultron":   folderUltron,
		"project_thor":    projectThor,
		"project_ironman": projectIronMan,
		"project_cap":     projectCap,
		"project_hulk":    projectHulk,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleProjectAncestryConfig(context),
				Check: resource.ComposeTestCheckFunc(
					// Project thor under organization
					resource.TestCheckResourceAttr("data.google_project_ancestry.thor", "ancestors.#", "2"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.thor", "ancestors.0.type", "project"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.thor", "ancestors.1.type", "organization"),

					// Project ironman under organization and thanos
					resource.TestCheckResourceAttr("data.google_project_ancestry.ironman", "ancestors.#", "3"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.ironman", "ancestors.0.type", "project"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.ironman", "ancestors.1.type", "folder"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.ironman", "ancestors.2.type", "organization"),

					// Project cap under organization, thanos and loki
					resource.TestCheckResourceAttr("data.google_project_ancestry.cap", "ancestors.#", "4"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.cap", "ancestors.0.type", "project"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.cap", "ancestors.1.type", "folder"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.cap", "ancestors.2.type", "folder"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.cap", "ancestors.3.type", "organization"),

					// Project hulk under organization, thanos, loki and ultron
					resource.TestCheckResourceAttr("data.google_project_ancestry.hulk", "ancestors.#", "5"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.hulk", "ancestors.0.type", "project"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.hulk", "ancestors.1.type", "folder"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.hulk", "ancestors.2.type", "folder"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.hulk", "ancestors.3.type", "folder"),
					resource.TestCheckResourceAttr("data.google_project_ancestry.hulk", "ancestors.4.type", "organization"),
				),
			},
		},
	})
}

func testAccCheckGoogleProjectAncestryConfig(context map[string]interface{}) string {
	return fmt.Sprintf(`
locals {
  org_id          = "%s"
  folder_thanos   = "%s"
  folder_loki     = "%s"
  folder_ultron   = "%s"
  project_thor    = "%s"
  project_ironman = "%s"
  project_cap     = "%s"
  project_hulk    = "%s"
}

resource "google_folder" "thanos" {
  deletion_protection = false
  display_name        = local.folder_thanos
  parent              = "organizations/${local.org_id}"
}

resource "google_folder" "loki" {
  deletion_protection = false
  display_name        = local.folder_loki
  parent              = google_folder.thanos.name
}

resource "google_folder" "ultron" {
  deletion_protection = false
  display_name        = local.folder_ultron
  parent              = google_folder.loki.name
}

resource "google_project" "thor" {
  deletion_policy = "DELETE"
  name            = local.project_thor
  org_id          = local.org_id
  project_id      = local.project_thor
}

resource "google_project" "ironman" {
  deletion_policy = "DELETE"
  folder_id       = google_folder.thanos.id
  name            = local.project_ironman
  project_id      = local.project_ironman
}

resource "google_project" "cap" {
  deletion_policy = "DELETE"
  folder_id       = google_folder.loki.id
  name            = local.project_cap
  project_id      = local.project_cap
}

resource "google_project" "hulk" {
  deletion_policy = "DELETE"
  folder_id       = google_folder.ultron.id
  name            = local.project_hulk
  project_id      = local.project_hulk
}

data "google_project_ancestry" "thor" {
  project = google_project.thor.project_id
}

data "google_project_ancestry" "ironman" {
  project = google_project.ironman.project_id
}

data "google_project_ancestry" "cap" {
  project = google_project.cap.project_id
}

data "google_project_ancestry" "hulk" {
  project = google_project.hulk.project_id
}
`,
		context["org_id"].(string),
		context["folder_thanos"].(string),
		context["folder_loki"].(string),
		context["folder_ultron"].(string),
		context["project_thor"].(string),
		context["project_ironman"].(string),
		context["project_cap"].(string),
		context["project_hulk"].(string),
	)
}
