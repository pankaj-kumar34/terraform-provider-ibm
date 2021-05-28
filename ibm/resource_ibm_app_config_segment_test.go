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

func TestAccIbmAppConfigSegmentBasic(t *testing.T) {
	var conf appconfigurationv1.Segment
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	segmentID := fmt.Sprintf("tf_segment_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	tags := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	tagsUpdate := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIbmAppConfigSegmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigSegmentConfig(instanceName, name, segmentID, description, tags),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmAppConfigSegmentExists("ibm_app_config_segment.app_config_segment_resource1", conf),
					resource.TestCheckResourceAttrSet("ibm_app_config_segment.app_config_segment_resource1", "name"),
					resource.TestCheckResourceAttrSet("ibm_app_config_segment.app_config_segment_resource1", "tags"),
					resource.TestCheckResourceAttrSet("ibm_app_config_segment.app_config_segment_resource1", "href"),
					resource.TestCheckResourceAttrSet("ibm_app_config_segment.app_config_segment_resource1", "rules.#"),
					// resource.TestCheckResourceAttrSet("ibm_app_config_segment.app_config_segment_resource1", "rules.0.values"),
					// resource.TestCheckResourceAttrSet("ibm_app_config_segment.app_config_segment_resource1", "rules.0.attribute_name"),
					// resource.TestCheckResourceAttrSet("ibm_app_config_segment.app_config_segment_resource1", "rules.0.operator"),
					resource.TestCheckResourceAttrSet("ibm_app_config_segment.app_config_segment_resource1", "updated_time"),
					resource.TestCheckResourceAttrSet("ibm_app_config_segment.app_config_segment_resource1", "created_time"),
					resource.TestCheckResourceAttrSet("ibm_app_config_segment.app_config_segment_resource1", "segment_id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_segment.app_config_segment_resource1", "description"),
				),
			},
			{
				Config: testAccCheckIbmAppConfigSegmentConfig(instanceName, nameUpdate, segmentID, descriptionUpdate, tagsUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_config_segment.app_config_segment_resource1", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_segment.app_config_segment_resource1", "tags", tagsUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_segment.app_config_segment_resource1", "description", descriptionUpdate),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigSegmentConfig(instanceName, name, segmentID, description, tags string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "app_config_terraform_test474"{
		name     = "%s"
		location = "us-south"
		service  = "apprapp"
		plan     = "lite"
	}
	resource "ibm_app_config_segment" "app_config_segment_resource1" {
		name 					= "%s"
		segment_id 		= "%s"
		description 	= "%s"
		tags 					= "%s"
		rules {
			attribute_name 	= "email"
			operator 				= "endsWith"
			values 					= ["@in.mnc.com"]
		}
		guid 					= ibm_resource_instance.app_config_terraform_test474.guid
		}
	`, instanceName, name, segmentID, description, tags)
}

func testAccCheckIbmAppConfigSegmentExists(n string, obj appconfigurationv1.Segment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		appconfigClient, err := getAppConfigClient(testAccProvider.Meta(), parts[0])
		if err != nil {
			return err
		}
		options := &appconfigurationv1.GetSegmentOptions{}

		options.SetSegmentID(parts[1])

		segment, _, err := appconfigClient.GetSegment(options)
		if err != nil {
			return err
		}

		obj = *segment
		return nil
	}
}

func testAccCheckIbmAppConfigSegmentDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "app_config_segment_resource1" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		appconfigClient, err := getAppConfigClient(testAccProvider.Meta(), parts[0])
		if err != nil {
			return err
		}
		optionis := &appconfigurationv1.GetSegmentOptions{}

		optionis.SetSegmentID(parts[1])

		// Try to find the key
		_, response, err := appconfigClient.GetSegment(optionis)

		if err == nil {
			return fmt.Errorf("Segment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Segment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
