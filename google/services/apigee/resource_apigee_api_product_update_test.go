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
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/apigee/resource_apigee_api_product_update_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package apigee_test

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"testing"
)

func TestAccApigeeApiProduct_apigeeApiProduct_full(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccCheckApigeeApiProductDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApigeeApiProduct_apigeeApiProduct_full(context),
			},
			{
				ResourceName:            "google_apigee_api_product.apigee_api_product",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"org_id"},
			},
			{
				Config: testAccApigeeApiProduct_apigeeApiProduct_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_apigee_api_product.apigee_api_product", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_apigee_api_product.apigee_api_product",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"org_id"},
			},
		},
	})
}

func testAccApigeeApiProduct_apigeeApiProduct_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
  deletion_policy = "DELETE"
}
resource "time_sleep" "wait_60_seconds" {
  create_duration = "60s"
  depends_on = [google_project.project]
}
resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
  depends_on = [time_sleep.wait_60_seconds]
}
resource "google_project_service" "compute" {
  project = google_project.project.project_id
  service = "compute.googleapis.com"
  depends_on = [google_project_service.apigee]
}
resource "google_project_service" "servicenetworking" {
  project = google_project.project.project_id
  service = "servicenetworking.googleapis.com"
  depends_on = [google_project_service.compute]
}
resource "time_sleep" "wait_120_seconds" {
  create_duration = "120s"
  depends_on = [google_project_service.servicenetworking]
}
resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [time_sleep.wait_120_seconds]
}
resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}
resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
  depends_on              = [google_project_service.servicenetworking]
}
resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id
  authorized_network = google_compute_network.apigee_network.id
  depends_on         = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee,
  ]
}
resource "google_apigee_instance" "apigee_instance" {
  name               = "tf-test%{random_suffix}"
  location           = "us-central1"
  org_id             = google_apigee_organization.apigee_org.id
  peering_cidr_range = "SLASH_22"
}
resource "google_apigee_api_product" "apigee_api_product" {
  org_id        = google_apigee_organization.apigee_org.id
  name              = "tf-test%{random_suffix}"
  display_name  = "My full API Product"

  approval_type = "auto"

  description   = "This is a sample API Product created with Terraform."

  quota               = "10000"
  quota_interval      = "1"
  quota_time_unit     = "day"
  quota_counter_scope = "PROXY"

  environments = ["dev", "hom"]
  scopes = [
    "read:weather",
    "write:reports"
  ]

  attributes {
    name  = "access"
    value = "private"
  }

  attributes {
    name  = "custom"
    value = "value"
  }

  operation_group {
    operation_config_type = "proxy"

    operation_configs {
      api_source = "anoter-proxy"

      operations {
        resource = "/"
        methods  = ["POST", "GET"]
      }

      quota {
        limit     = "1000"
        interval  = "5"
        time_unit = "minute"
      }

      attributes {
        name  = "custom"
        value = "value"
      }
    }

    operation_configs {
      api_source = "hello-world"

      operations {
        resource = "/test"
        methods  = ["POST", "GET"]
      }

      quota {
        limit     = "10"
        interval  = "30"
        time_unit = "second"
      }

      attributes {
        name  = "custom"
        value = "value"
      }
    }
  }

  graphql_operation_group {
    operation_config_type = "proxy"

    operation_configs {
      api_source = "hello-world"

      quota {
        limit     = "30"
        interval  = "50"
        time_unit = "second"
      }

      operations {
        operation_types = ["QUERY"]
        operation       = "test"
      }

      attributes {
        name  = "custom"
        value = "value"
      }
    }

    operation_configs {
      api_source = "another-proxy"

      quota {
        limit     = "50000"
        interval  = "12"
        time_unit = "hour"
      }

      operations {
        operation_types = ["MUTATION"]
        operation       = "test"
      }

      attributes {
        name  = "custom"
        value = "value"
      }
    }
  }

  grpc_operation_group {

    operation_configs {
      api_source = "another-proxy"
      service    = "grpc another test"
      methods    = ["method3", "method4"]

      quota {
        limit     = "1000000"
        interval  = "1"
        time_unit = "month"
      }

      attributes {
        name  = "graph"
        value = "value"
      }
    }

    operation_configs {
      api_source = "hello-world"
      service    = "grpc test"
      methods    = ["method1", "method2"]

      quota {
        limit     = "5"
        interval  = "1"
        time_unit = "second"
      }

      attributes {
        name  = "graph"
        value = "value"
      }
    }
  }

  depends_on = [
    google_apigee_instance.apigee_instance
  ]
}
`, context)
}

func testAccApigeeApiProduct_apigeeApiProduct_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
  deletion_policy = "DELETE"
}
resource "time_sleep" "wait_60_seconds" {
  create_duration = "60s"
  depends_on = [google_project.project]
}
resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
  depends_on = [time_sleep.wait_60_seconds]
}
resource "google_project_service" "compute" {
  project = google_project.project.project_id
  service = "compute.googleapis.com"
  depends_on = [google_project_service.apigee]
}
resource "google_project_service" "servicenetworking" {
  project = google_project.project.project_id
  service = "servicenetworking.googleapis.com"
  depends_on = [google_project_service.compute]
}
resource "time_sleep" "wait_120_seconds" {
  create_duration = "120s"
  depends_on = [google_project_service.servicenetworking]
}
resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [time_sleep.wait_120_seconds]
}
resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}
resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
  depends_on              = [google_project_service.servicenetworking]
}
resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id
  authorized_network = google_compute_network.apigee_network.id
  depends_on         = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee,
  ]
}
resource "google_apigee_instance" "apigee_instance" {
  name               = "tf-test%{random_suffix}"
  location           = "us-central1"
  org_id             = google_apigee_organization.apigee_org.id
  peering_cidr_range = "SLASH_22"
}
resource "google_apigee_developer" "apigee_developer" {
  email      = "tf-test%{random_suffix}@acme.com"
  first_name = "John"
  last_name  = "Doe"
  user_name  = "john.doe"
  org_id     = google_apigee_organization.apigee_org.id
  depends_on = [
    google_apigee_instance.apigee_instance
  ]
}
resource "google_apigee_api_product" "apigee_api_product" {
  org_id        = google_apigee_organization.apigee_org.id
  name              = "tf-test%{random_suffix}"
  display_name  = "My full API Product"

  approval_type = "auto"

  description   = "This is a sample API Product created with Terraform."

  quota               = "5000"
  quota_interval      = "2"
  quota_time_unit     = "day"
  quota_counter_scope = "PROXY"

  environments = ["dev"]
  scopes = [
    "read:weather"
  ]

  attributes {
    name  = "access"
    value = "private"
  }

  attributes {
    name  = "custom"
    value = "value_changed"
	}

  operation_group {
    operation_config_type = "proxy"

    operation_configs {
      api_source = "anoter-proxy"

      operations {
        resource = "/changed"
        methods  = ["POST", "GET", "PUT"]
      }

      quota {
        limit     = "500"
        interval  = "6"
        time_unit = "minute"
      }

      attributes {
        name  = "custom"
        value = "value_changed"
      }
    }

    operation_configs {
      api_source = "hello-world"

      operations {
        resource = "/test_changed"
        methods  = ["POST"]
      }

      quota {
        limit     = "7"
        interval  = "20"
        time_unit = "second"
      }

      attributes {
        name  = "custom"
        value = "value_changed"
      }
    }
  }

  graphql_operation_group {
    operation_config_type = "proxy"

    operation_configs {
      api_source = "hello-world"

      quota {
        limit     = "20"
        interval  = "40"
        time_unit = "second"
      }

      operations {
        operation_types = ["MUTATION"]
        operation       = "test_changed"
      }

      attributes {
        name  = "custom"
        value = "value_changed"
      }
    }

    operation_configs {
      api_source = "another-proxy"

      quota {
        limit     = "5000"
        interval  = "10"
        time_unit = "hour"
      }

      operations {
        operation_types = ["QUERY"]
        operation       = "test_changed"
      }

      attributes {
        name  = "custom"
        value = "value_changed"
      }
    }
  }

  grpc_operation_group {

    operation_configs {
      api_source = "another-proxy"
      service    = "grpc another test"
      methods    = ["method3_changed", "method4_changed"]

      quota {
        limit     = "10000"
        interval  = "10"
        time_unit = "month"
      }

      attributes {
        name  = "graph"
        value = "value_changed"
      }
    }

    operation_configs {
      api_source = "hello-world"
      service    = "grpc test"
      methods    = ["method1_changed", "method2_changed"]

      quota {
        limit     = "50"
        interval  = "5"
        time_unit = "hour"
      }

      attributes {
        name  = "graph"
        value = "value_changed"
      }
    }
  }

  depends_on = [
    google_apigee_instance.apigee_instance
  ]
}
`, context)
}
