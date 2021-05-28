// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigPropertyDataSourceBasic(t *testing.T) {
	propertyEnvironmentID := fmt.Sprintf("tf_environment_id_%d", acctest.RandIntRange(10, 100))
	propertyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	propertyPropertyID := fmt.Sprintf("tf_property_id_%d", acctest.RandIntRange(10, 100))
	propertyType := "BOOLEAN"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigPropertyDataSourceConfigBasic(propertyEnvironmentID, propertyName, propertyPropertyID, propertyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "environment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "property_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "value"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "tags"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "segment_rules.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "segment_exists"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "collections.#"),
					resource.TestCheckResourceAttr("data.ibm_app_config_property.app_config_property", "collections.0.name", propertyName),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "evaluation_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "href"),
				),
			},
		},
	})
}

func TestAccIbmAppConfigPropertyDataSourceAllArgs(t *testing.T) {
	propertyEnvironmentID := fmt.Sprintf("tf_environment_id_%d", acctest.RandIntRange(10, 100))
	propertyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	propertyPropertyID := fmt.Sprintf("tf_property_id_%d", acctest.RandIntRange(10, 100))
	propertyType := "BOOLEAN"
	propertyDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	propertyTags := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigPropertyDataSourceConfig(propertyEnvironmentID, propertyName, propertyPropertyID, propertyType, propertyDescription, propertyTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "environment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "property_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "value"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "tags"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "segment_rules.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "segment_rules.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "segment_rules.0.order"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "segment_exists"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "collections.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "collections.0.collection_id"),
					resource.TestCheckResourceAttr("data.ibm_app_config_property.app_config_property", "collections.0.name", propertyName),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "evaluation_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_property.app_config_property", "href"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigPropertyDataSourceConfigBasic(propertyEnvironmentID string, propertyName string, propertyPropertyID string, propertyType string) string {
	return fmt.Sprintf(`
		resource "ibm_app_config_property" "app_config_property" {
			environment_id = "%s"
			name = "%s"
			property_id = "%s"
			type = "%s"
			value = "FIXME"
		}

		data "ibm_app_config_property" "app_config_property" {
			environment_id = ibm_app_config_property.app_config_property.environment_id
			property_id = ibm_app_config_property.app_config_property.property_id
		}
	`, propertyEnvironmentID, propertyName, propertyPropertyID, propertyType)
}

func testAccCheckIbmAppConfigPropertyDataSourceConfig(propertyEnvironmentID string, propertyName string, propertyPropertyID string, propertyType string, propertyDescription string, propertyTags string) string {
	return fmt.Sprintf(`
		resource "ibm_app_config_property" "app_config_property" {
			environment_id = "%s"
			name = "%s"
			property_id = "%s"
			type = "%s"
			value = "FIXME"
			description = "%s"
			tags = "%s"
			segment_rules {
				rules {
					segments = ["betausers","premiumusers"]
				}
				value = true
				order = 1
			}
			collections {
				collection_id = "ghzinc"
			}
		}

		data "ibm_app_config_property" "app_config_property" {
			environment_id = ibm_app_config_property.app_config_property.environment_id
			property_id = ibm_app_config_property.app_config_property.property_id
		}
	`, propertyEnvironmentID, propertyName, propertyPropertyID, propertyType, propertyDescription, propertyTags)
}
