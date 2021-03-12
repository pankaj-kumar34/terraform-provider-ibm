package ibm

import (
	"fmt"
	"testing"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMResourcePNApplicationChrome_Basic(t *testing.T) {
	var conf pushservicev1.ChromeWebPushCredendialsModel
	name := fmt.Sprintf("terraform_PN_%d", acctest.RandIntRange(10, 100))
	senderID := fmt.Sprint(acctest.RandString(45))                // dummy value
	websiteURL := "http://webpushnotificatons.mybluemix.net"      // dummy url
	newWebsiteURL := "http://chromepushnotificaton.mybluemix.net" // dummy url
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourcePNApplicationChromeConfig(name, senderID, websiteURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMResourcePNApplicationChromeExists("ibm_pn_application_chrome.application_chrome", conf),
					resource.TestCheckResourceAttrSet("ibm_pn_application_chrome.application_chrome", "id"),
				),
			},
			{
				Config: testAccCheckIBMResourcePNApplicationChromeUpdate(name, senderID, newWebsiteURL),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pn_application_chrome.application_chrome", "web_site_url", newWebsiteURL),
				),
			},
		},
	})
}

func testAccCheckIBMResourcePNApplicationChromeConfig(name, senderID, websiteURL string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "push_notification"{
		name     = "%s"
		location = "us-south"
		service  = "imfpush"
		plan     = "lite"
	}
	resource "ibm_pn_application_chrome" "application_chrome" {
		sender_id            		= "%s"
		web_site_url           = "%s"
		application_id = ibm_resource_instance.push_notification.guid
	}`, name, senderID, websiteURL)
}

func testAccCheckIBMResourcePNApplicationChromeUpdate(name, senderID, newWebsiteURL string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "push_notification"{
			name     = "%s"
			location = "us-south"
			service  = "imfpush"
			plan     = "lite"
		}
		resource "ibm_pn_application_chrome" "application_chrome" {
			sender_id            	 = "%s"
			web_site_url           = "%s"
			application_id = ibm_resource_instance.push_notification.guid
		}`, name, senderID, newWebsiteURL)
}

func testAccCheckIBMResourcePNApplicationChromeExists(n string, obj pushservicev1.ChromeWebPushCredendialsModel) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		pushServiceClient, err := testAccProvider.Meta().(ClientSession).PushServiceV1()
		if err != nil {
			return err
		}

		getChromeWebConfOptions := &pushservicev1.GetChromeWebConfOptions{}

		applicationID := rs.Primary.ID

		getChromeWebConfOptions.SetApplicationID(applicationID)

		chromeConf, _, err := pushServiceClient.GetChromeWebConf(getChromeWebConfOptions)
		if err != nil {
			return err
		}

		obj = *chromeConf
		return nil
	}
}
