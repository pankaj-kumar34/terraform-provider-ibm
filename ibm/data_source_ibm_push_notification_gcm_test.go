package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDataSourcePNApplicationGCM_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_PN_%d", acctest.RandIntRange(10, 100))
	key := fmt.Sprint(acctest.RandString(25)) // dummy value
	senderID := fmt.Sprint(acctest.RandInt()) // dummy value
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDataSourcePNApplicationGCMConfig(name, senderID, key),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pn_application_gcm.gcm", "api_key"),
					resource.TestCheckResourceAttrSet("data.ibm_pn_application_gcm.gcm", "sender_id"),
				),
			},
		},
	})
}

func testAccCheckIBMDataSourcePNApplicationGCMConfig(name, senderID, key string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "push_notification"{
			name     = "%s"
			location = "us-south"
			service  = "imfpush"
			plan     = "lite"
		}
		resource "ibm_pn_application_gcm" "application_gcm" {
			server_key            = "%s"
			sender_id             = "%s"
			service_instance_guid = ibm_resource_instance.push_notification.guid
		}
		data "ibm_pn_application_gcm" "gcm" {
			service_instance_guid = ibm_pn_application_gcm.application_gcm.id
		}`, name, key, senderID)
}
