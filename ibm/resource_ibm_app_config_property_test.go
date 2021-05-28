// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func TestAccIbmAppConfigPropertyBasic(t *testing.T) {
	var conf appconfigurationv1.Property
	environmentID := fmt.Sprintf("tf_environment_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	propertyID := fmt.Sprintf("tf_property_id_%d", acctest.RandIntRange(10, 100))
	typeVar := "BOOLEAN"
	environmentIDUpdate := fmt.Sprintf("tf_environment_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	propertyIDUpdate := fmt.Sprintf("tf_property_id_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "NUMERIC"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmAppConfigPropertyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmAppConfigPropertyConfigBasic(environmentID, name, propertyID, typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmAppConfigPropertyExists("ibm_app_config_property.app_config_property", conf),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "environment_id", environmentID),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "name", name),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "property_id", propertyID),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "type", typeVar),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmAppConfigPropertyConfigBasic(environmentIDUpdate, nameUpdate, propertyIDUpdate, typeVarUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "environment_id", environmentIDUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "property_id", propertyIDUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "type", typeVarUpdate),
				),
			},
		},
	})
}

func TestAccIbmAppConfigPropertyAllArgs(t *testing.T) {
	var conf appconfigurationv1.Property
	environmentID := fmt.Sprintf("tf_environment_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	propertyID := fmt.Sprintf("tf_property_id_%d", acctest.RandIntRange(10, 100))
	typeVar := "BOOLEAN"
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	tags := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))
	environmentIDUpdate := fmt.Sprintf("tf_environment_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	propertyIDUpdate := fmt.Sprintf("tf_property_id_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "NUMERIC"
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	tagsUpdate := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmAppConfigPropertyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmAppConfigPropertyConfig(environmentID, name, propertyID, typeVar, description, tags),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmAppConfigPropertyExists("ibm_app_config_property.app_config_property", conf),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "environment_id", environmentID),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "name", name),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "property_id", propertyID),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "description", description),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "tags", tags),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmAppConfigPropertyConfig(environmentIDUpdate, nameUpdate, propertyIDUpdate, typeVarUpdate, descriptionUpdate, tagsUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "environment_id", environmentIDUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "property_id", propertyIDUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_property.app_config_property", "tags", tagsUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_app_config_property.app_config_property",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmAppConfigPropertyConfigBasic(environmentID string, name string, propertyID string, typeVar string) string {
	return fmt.Sprintf(`

		resource "ibm_app_config_property" "app_config_property" {
			environment_id = "%s"
			name = "%s"
			property_id = "%s"
			type = "%s"
			value = "FIXME"
		}
	`, environmentID, name, propertyID, typeVar)
}

func testAccCheckIbmAppConfigPropertyConfig(environmentID string, name string, propertyID string, typeVar string, description string, tags string) string {
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
	`, environmentID, name, propertyID, typeVar, description, tags)
}

func testAccCheckIbmAppConfigPropertyExists(n string, obj appconfigurationv1.Property) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		appConfigurationClient, err := testAccProvider.Meta().(ClientSession).AppConfigurationV1()
		if err != nil {
			return err
		}

		getPropertyOptions := &appconfigurationv1.GetPropertyOptions{}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getPropertyOptions.SetEnvironmentID(parts[0])
		getPropertyOptions.SetPropertyID(parts[1])

		property, _, err := appConfigurationClient.GetProperty(getPropertyOptions)
		if err != nil {
			return err
		}

		obj = *property
		return nil
	}
}

func testAccCheckIbmAppConfigPropertyDestroy(s *terraform.State) error {
	appConfigurationClient, err := testAccProvider.Meta().(ClientSession).AppConfigurationV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_app_config_property" {
			continue
		}

		getPropertyOptions := &appconfigurationv1.GetPropertyOptions{}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getPropertyOptions.SetEnvironmentID(parts[0])
		getPropertyOptions.SetPropertyID(parts[1])

		// Try to find the key
		_, response, err := appConfigurationClient.GetProperty(getPropertyOptions)

		if err == nil {
			return fmt.Errorf("app_config_property still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for app_config_property (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
