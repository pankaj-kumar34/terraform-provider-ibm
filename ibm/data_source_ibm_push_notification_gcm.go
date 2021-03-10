package ibm

import (
	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMPnApplicationGCM() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplicationGCMRead,

		Schema: map[string]*schema.Schema{
			"service_instance_guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance guid of the push notifications instance",
			},
			"sender_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Sender ID or Project Number from the Google Developer Console",
			},
			"server_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server key/Legacy server key for the Sender ID",
			},
		},
	}
}

func dataSourceApplicationGCMRead(d *schema.ResourceData, meta interface{}) error {
	pnClient, err := meta.(ClientSession).PushNotificationsV1API()
	if err != nil {
		return err
	}

	serviceInstanceGUID := d.Get("service_instance_guid").(string)

	result, response, err := pnClient.GetGCMConf(&pushservicev1.GetGCMConfOptions{
		ApplicationID: &serviceInstanceGUID,
	})

	if err != nil {
		return err
	}

	if response.StatusCode == 200 {
		d.SetId(serviceInstanceGUID)
		d.Set("server_key", *result.ApiKey)
		d.Set("sender_id", *result.SenderID)
		d.Set("service_instance_guid", serviceInstanceGUID)
	}
	return nil
}
