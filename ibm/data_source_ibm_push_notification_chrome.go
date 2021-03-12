package ibm

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMPNApplicationChrome() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplicationChromeRead,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Unique guid of the application using the push service.",
			},
			"sender_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An senderId that gives the push service an authorized access to Google services that is used for Chrome Web Push.",
			},
			"web_site_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the WebSite / WebApp that should be permitted to subscribe to WebPush.",
			},
		},
	}
}

func dataSourceApplicationChromeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pushServiceClient, err := meta.(ClientSession).PushServiceV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getChromeWebConfOptions := &pushservicev1.GetChromeWebConfOptions{}

	applicationID := d.Get("application_id").(string)
	getChromeWebConfOptions.SetApplicationID(applicationID)

	chromeWebPushCredendialsModel, response, err := pushServiceClient.GetChromeWebConfWithContext(context, getChromeWebConfOptions)
	if err != nil {
		log.Printf("[DEBUG] GetChromeWebConfWithContext failed %s\n%d", err, response.StatusCode)
		return diag.FromErr(err)
	}

	d.SetId(applicationID)
	if err = d.Set("sender_id", chromeWebPushCredendialsModel.ApiKey); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting sender_id: %s", err))
	}
	if err = d.Set("web_site_url", chromeWebPushCredendialsModel.WebSiteURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting web_site_url: %s", err))
	}

	return nil
}
