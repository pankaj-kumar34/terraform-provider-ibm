// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigSegmentsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmAppConfigSegmentsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "segments.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "count"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "first"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "previous"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "last"),
				),
			},
		},
	})
}

func TestAccIbmAppConfigSegmentsDataSourceAllArgs(t *testing.T) {
	segmentName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	segmentSegmentID := fmt.Sprintf("tf_segment_id_%d", acctest.RandIntRange(10, 100))
	segmentDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	segmentTags := fmt.Sprintf("tf_tags_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmAppConfigSegmentsDataSourceConfig(segmentName, segmentSegmentID, segmentDescription, segmentTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "segments.#"),
					resource.TestCheckResourceAttr("data.ibm_app_config_segments.app_config_segments", "segments.0.name", segmentName),
					resource.TestCheckResourceAttr("data.ibm_app_config_segments.app_config_segments", "segments.0.segment_id", segmentSegmentID),
					resource.TestCheckResourceAttr("data.ibm_app_config_segments.app_config_segments", "segments.0.description", segmentDescription),
					resource.TestCheckResourceAttr("data.ibm_app_config_segments.app_config_segments", "segments.0.tags", segmentTags),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "segments.0.created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "segments.0.updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "segments.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "count"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "first"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "previous"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_segments.app_config_segments", "last"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigSegmentsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_app_config_segment" "app_config_segment" {
		}

		data "ibm_app_config_segments" "app_config_segments" {
		}
	`)
}

func testAccCheckIbmAppConfigSegmentsDataSourceConfig(segmentName string, segmentSegmentID string, segmentDescription string, segmentTags string) string {
	return fmt.Sprintf(`
		resource "ibm_app_config_segment" "app_config_segment" {
			name = "%s"
			segment_id = "%s"
			description = "%s"
			tags = "%s"
			rules {
				attribute_name = "email"
				operator = "endsWith"
				values = ["@in.mnc.com","@us.mnc.com"]
			}
		}

		data "ibm_app_config_segments" "app_config_segments" {
		}
	`, segmentName, segmentSegmentID, segmentDescription, segmentTags)
}
