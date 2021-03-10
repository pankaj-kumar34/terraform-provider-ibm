package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMResourcePNApplicationGCM_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_PN_%d", acctest.RandIntRange(10, 100))
	key := fmt.Sprintf(acctest.RandString(25))    // dummy value
	senderID := fmt.Sprint(acctest.RandInt())     // dummy value
	newKey := fmt.Sprintf(acctest.RandString(25)) // dummy value
	newSenderID := fmt.Sprint(acctest.RandInt())  // dummy value
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourcePNApplicationGCMConfig(name, senderID, key),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pn_application_gcm.application_gcm", "server_key", key),
					resource.TestCheckResourceAttr("ibm_pn_application_gcm.application_gcm", "sender_id", senderID),
				),
			},
			{
				Config: testAccCheckIBMResourcePNApplicationGCMUpdate(name, newSenderID, newKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pn_application_gcm.application_gcm", "server_key", newKey),
					resource.TestCheckResourceAttr("ibm_pn_application_gcm.application_gcm", "sender_id", newSenderID),
				),
			},
		},
	})
}

func testAccCheckIBMResourcePNApplicationGCMConfig(name string, senderID string, key string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "push_notification"{
		name     = "%s"
		location = "us-south"
		service  = "imfpush"
		plan     = "lite"
	}
	resource "ibm_pn_application_gcm" "application_gcm" {
		server_key               = "%s"
		sender_id             = "%s"
		service_instance_guid = ibm_resource_instance.push_notification.guid
	}`, name, key, senderID)
}

func testAccCheckIBMResourcePNApplicationGCMUpdate(name, newSenderID, newKey string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "push_notification"{
			name     = "%s"
			location = "us-south"
			service  = "imfpush"
			plan     = "lite"
		}
		resource "ibm_pn_application_gcm" "application_gcm" {
			server_key               = "%s"
			sender_id             = "%s"
			service_instance_guid = ibm_resource_instance.push_notification.guid
		}`, name, newKey, newSenderID)
}
