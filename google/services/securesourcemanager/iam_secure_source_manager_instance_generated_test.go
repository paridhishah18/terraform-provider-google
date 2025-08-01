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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securesourcemanager/Instance.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/examples/base_configs/iam_test_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securesourcemanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccSecureSourceManagerInstanceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"role":            "roles/securesourcemanager.instanceManager",
		"admin_role":      "roles/securesourcemanager.instanceOwner",
		"deletion_policy": "DELETE",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerInstanceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_secure_source_manager_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s roles/securesourcemanager.instanceManager", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccSecureSourceManagerInstanceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_secure_source_manager_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s roles/securesourcemanager.instanceManager", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecureSourceManagerInstanceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"role":            "roles/securesourcemanager.instanceManager",
		"admin_role":      "roles/securesourcemanager.instanceOwner",
		"deletion_policy": "DELETE",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccSecureSourceManagerInstanceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_secure_source_manager_instance_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s roles/securesourcemanager.instanceManager user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSecureSourceManagerInstanceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	// This may skip test, so do it first
	sa := envvar.GetTestServiceAccountFromEnv(t)
	context := map[string]interface{}{
		"random_suffix":   acctest.RandString(t, 10),
		"role":            "roles/securesourcemanager.instanceManager",
		"admin_role":      "roles/securesourcemanager.instanceOwner",
		"deletion_policy": "DELETE",
	}
	context["service_account"] = sa

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerInstanceIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_secure_source_manager_instance_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_secure_source_manager_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSecureSourceManagerInstanceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_secure_source_manager_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSecureSourceManagerInstanceIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "default" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
    labels = {
      "foo" = "bar"
    }

    # Prevent accidental deletions.
    deletion_policy = "%{deletion_policy}"
}

resource "google_secure_source_manager_instance_iam_member" "foo" {
  project = google_secure_source_manager_instance.default.project
  location = google_secure_source_manager_instance.default.location
  instance_id = google_secure_source_manager_instance.default.instance_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccSecureSourceManagerInstanceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "default" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
    labels = {
      "foo" = "bar"
    }

    # Prevent accidental deletions.
    deletion_policy = "%{deletion_policy}"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
  binding {
    role = "%{admin_role}"
    members = ["serviceAccount:%{service_account}"]
  }
}

resource "google_secure_source_manager_instance_iam_policy" "foo" {
  project = google_secure_source_manager_instance.default.project
  location = google_secure_source_manager_instance.default.location
  instance_id = google_secure_source_manager_instance.default.instance_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_secure_source_manager_instance_iam_policy" "foo" {
  project = google_secure_source_manager_instance.default.project
  location = google_secure_source_manager_instance.default.location
  instance_id = google_secure_source_manager_instance.default.instance_id
  depends_on = [
    google_secure_source_manager_instance_iam_policy.foo
  ]
}
`, context)
}

func testAccSecureSourceManagerInstanceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "default" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
    labels = {
      "foo" = "bar"
    }

    # Prevent accidental deletions.
    deletion_policy = "%{deletion_policy}"
}

data "google_iam_policy" "foo" {
}

resource "google_secure_source_manager_instance_iam_policy" "foo" {
  project = google_secure_source_manager_instance.default.project
  location = google_secure_source_manager_instance.default.location
  instance_id = google_secure_source_manager_instance.default.instance_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccSecureSourceManagerInstanceIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "default" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
    labels = {
      "foo" = "bar"
    }

    # Prevent accidental deletions.
    deletion_policy = "%{deletion_policy}"
}

resource "google_secure_source_manager_instance_iam_binding" "foo" {
  project = google_secure_source_manager_instance.default.project
  location = google_secure_source_manager_instance.default.location
  instance_id = google_secure_source_manager_instance.default.instance_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccSecureSourceManagerInstanceIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "default" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
    labels = {
      "foo" = "bar"
    }

    # Prevent accidental deletions.
    deletion_policy = "%{deletion_policy}"
}

resource "google_secure_source_manager_instance_iam_binding" "foo" {
  project = google_secure_source_manager_instance.default.project
  location = google_secure_source_manager_instance.default.location
  instance_id = google_secure_source_manager_instance.default.instance_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
