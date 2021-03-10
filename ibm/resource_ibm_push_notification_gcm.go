package ibm

import (
	"log"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMPnApplicationGCM() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationGCMCreate,
		Update: resourceApplicationGCMUpdate,
		Delete: resourceApplicationGCMDelete,
		Read:   resourceApplicationGCMRead,

		Schema: map[string]*schema.Schema{
			"service_instance_guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance guid of the push notifications instance",
			},
			"sender_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Sender ID or Project Number from the Google Developer Console",
			},
			"server_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Server key/Legacy server key for the Sender ID",
			},
		},
	}
}

func resourceApplicationGCMCreate(d *schema.ResourceData, meta interface{}) error {
	pnClient, err := meta.(ClientSession).PushNotificationsV1API()
	if err != nil {
		return err
	}

	apiKey := d.Get("server_key").(string)
	senderID := d.Get("sender_id").(string)
	serviceInstanceGUID := d.Get("service_instance_guid").(string)

	_, _, e := pnClient.SaveGCMConf(&pushservicev1.SaveGCMConfOptions{
		ApplicationID: &serviceInstanceGUID,
		ApiKey:        &apiKey,
		SenderID:      &senderID,
	})

	if e != nil {
		log.Fatal(e)
	}

	return resourceApplicationGCMRead(d, meta)
}

func resourceApplicationGCMUpdate(d *schema.ResourceData, meta interface{}) error {

	if d.HasChanges("server_key", "sender_id") {
		return resourceApplicationGCMCreate(d, meta)
	}
	return nil
}

func resourceApplicationGCMRead(d *schema.ResourceData, meta interface{}) error {
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

func resourceApplicationGCMDelete(d *schema.ResourceData, meta interface{}) error {
	pnClient, err := meta.(ClientSession).PushNotificationsV1API()
	if err != nil {
		return err
	}
	serviceInstanceGUID := d.Get("service_instance_guid").(string)

	_, e := pnClient.DeleteGCMConf(&pushservicev1.DeleteGCMConfOptions{
		ApplicationID: &serviceInstanceGUID,
	})

	if e != nil {
		return e
	}

	d.SetId("")

	return nil

}
