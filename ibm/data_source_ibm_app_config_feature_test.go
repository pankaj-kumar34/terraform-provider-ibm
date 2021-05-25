/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmAppConfigFeatureDataSource(t *testing.T) {
	environmentID := "dev"
	featureType := "BOOLEAN"
	tags := "development feature"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	featureID := fmt.Sprintf("tf_feature_id_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigFeatureDataSourceConfigBasic(instanceName, name, environmentID, featureID, featureType, description, tags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "tags"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "feature_id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "enabled_value"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "segment_exists"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "disabled_value"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_feature.ibm_app_config_feature_data1", "href"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigFeatureDataSourceConfigBasic(instanceName, name, environmentID, featureID, featureType, description, tags string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "app_config_terraform_test482" {
			name     = "%s"
			location = "us-south"
			service  = "apprapp"
			plan     = "standard"
		}
		
		resource "ibm_app_config_feature" "app_config_feature_resource1" {
			guid           	= ibm_resource_instance.app_config_terraform_test482.guid
			name           	= "%s"
			environment_id  = "%s"
			feature_id     	= "%s"
			type           	= "%s"
			enabled_value  	= "true"
			disabled_value 	= "false"
			description    	= "%s"
			tags    			 	= "%s"
		}
		
		data "ibm_app_config_feature" "ibm_app_config_feature_data1" {
			guid          = ibm_app_config_feature.app_config_feature_resource1.guid
			feature_id    = ibm_app_config_feature.app_config_feature_resource1.feature_id
			environment_id = ibm_app_config_feature.app_config_feature_resource1.environment_id
		}
		`, instanceName, name, environmentID, featureID, featureType, description, tags)
}
