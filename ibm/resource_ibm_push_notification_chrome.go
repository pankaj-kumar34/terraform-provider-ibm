package ibm

import (
	"fmt"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMPNApplicationChrome() *schema.Resource {
	return &schema.Resource{
		Read:   resourceApplicationChromeRead,
		Create: resourceApplicationChromeCreate,
		Update: resourceApplicationChromeUpdate,
		Delete: resourceApplicationChromeDelete,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Unique guid of the application using the push service.",
			},
			"sender_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "An senderId that gives the push service an authorized access to Google services that is used for Chrome Web Push.",
			},
			"web_site_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL of the WebSite / WebApp that should be permitted to subscribe to WebPush.",
			},
		},
	}
}

func resourceApplicationChromeCreate(d *schema.ResourceData, meta interface{}) error {
	pnClient, err := meta.(ClientSession).PushServiceV1()
	if err != nil {
		return err
	}

	serverKey := d.Get("sender_id").(string)
	websiteURL := d.Get("web_site_url").(string)
	applicationID := d.Get("application_id").(string)

	_, response, err := pnClient.SaveChromeWebConf(&pushservicev1.SaveChromeWebConfOptions{
		ApplicationID: &applicationID,
		ApiKey:        &serverKey,
		WebSiteURL:    &websiteURL,
	})

	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error configuring chrome web platform: %s with responce code  %d", err, response.StatusCode)
	}
	d.SetId(applicationID)

	return resourceApplicationChromeRead(d, meta)
}

func resourceApplicationChromeUpdate(d *schema.ResourceData, meta interface{}) error {

	if d.HasChanges("sender_id", "web_site_url") {
		return resourceApplicationChromeCreate(d, meta)
	}
	return nil
}

func resourceApplicationChromeRead(d *schema.ResourceData, meta interface{}) error {
	pnClient, err := meta.(ClientSession).PushServiceV1()
	if err != nil {
		return err
	}

	applicationID := d.Id()

	result, response, err := pnClient.GetChromeWebConf(&pushservicev1.GetChromeWebConfOptions{
		ApplicationID: &applicationID,
	})

	if err != nil {
		if response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error fetching chrome web platform configuration: %s with responce code  %d", err, response.StatusCode)
	}

	d.SetId(applicationID)

	if response.StatusCode == 200 {
		d.Set("sender_id", *result.ApiKey)
		d.Set("web_site_url", *result.WebSiteURL)
	}
	return nil
}

func resourceApplicationChromeDelete(d *schema.ResourceData, meta interface{}) error {
	pnClient, err := meta.(ClientSession).PushServiceV1()
	if err != nil {
		return err
	}
	applicationID := d.Get("application_id").(string)

	response, err := pnClient.DeleteChromeWebConf(&pushservicev1.DeleteChromeWebConfOptions{
		ApplicationID: &applicationID,
	})

	if err != nil {
		if response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error deleting chrome web platform configuration: %s with responce code  %d", err, response.StatusCode)
	}

	d.SetId("")

	return nil

}
