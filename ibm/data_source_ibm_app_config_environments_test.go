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

func TestAccIbmAppConfigEnvironmentsDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	colorCode := "#e23433"
	tags := fmt.Sprintf("tags_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	envName := fmt.Sprintf("env_%d", acctest.RandIntRange(10, 100))
	environmentID := fmt.Sprintf("environment_id_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigEnvironmentsDataSourceConfigBasic(name, envName, environmentID, description, colorCode, tags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environments.app_config_environments_data2", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environments.app_config_environments_data2", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environments.app_config_environments_data2", "last.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environments.app_config_environments_data2", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environments.app_config_environments_data2", "environments.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_environments.app_config_environments_data2", "environments.0.name"),
					resource.TestCheckResourceAttr("data.ibm_app_config_environments.app_config_environments_data2", "environments.0.name", envName),
					resource.TestCheckResourceAttr("data.ibm_app_config_environments.app_config_environments_data2", "environments.0.environment_id", environmentID),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigEnvironmentsDataSourceConfigBasic(name, envName, environmentID, description, colorCode, tags string) string {
	return fmt.Sprintf(`
		 resource "ibm_resource_instance" "app_config_terraform_test48"{
			 name     = "%s"
			 location = "us-south"
			 service  = "apprapp"
			 plan     = "standard"
		 }
		 resource "ibm_app_config_environment" "app_config_environment_resource2" {
			 name          		= "%s"
			 environment_id    = "%s"
			 description       = "%s"
			 color_code        = "%s"
			 tags              = "%s"
			 guid = ibm_resource_instance.app_config_terraform_test48.guid
		 }
		 data "ibm_app_config_environments" "app_config_environments_data2" {
			 expand            = true
			 guid 							= ibm_app_config_environment.app_config_environment_resource2.guid
		 }`, name, envName, environmentID, description, colorCode, tags)
}
