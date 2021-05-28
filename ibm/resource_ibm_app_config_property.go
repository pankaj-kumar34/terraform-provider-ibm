// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func resourceIbmAppConfigProperty() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIbmAppConfigPropertyCreate,
		Read:     resourceIbmAppConfigPropertyRead,
		Update:   resourceIbmAppConfigPropertyUpdate,
		Delete:   resourceIbmAppConfigPropertyDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Environment Id.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Property name.",
			},
			"property_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Property id.",
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_app_config_type", "type"),
				Description:  "Type of the feature (BOOLEAN, STRING, NUMERIC).",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Value of the Property. The value can be Boolean, String or a Numeric value as per the `type` attribute.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Property description.",
			},
			"tags": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Tags associated with the property.",
			},
			"segment_rules": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specify the targeting rules that is used to set different property values for different segments.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rules": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Rules array.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"segments": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of segment ids that are used for targeting using the rule.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Value to be used for evaluation for this rule. The value can be Boolean, String or a Numeric value as per the `type` attribute.",
						},
						"order": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation.",
						},
					},
				},
			},
			"collections": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of collection id representing the collections that are associated with the specified property.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"collection_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Collection id.",
						},
					},
				},
			},
			"segment_exists": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Denotes if the targeting rules are specified for the property.",
			},
			"created_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of the property.",
			},
			"updated_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last modified time of the property data.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Property URL.",
			},
		},
	}
}

func resourceIbmAppConfigPropertyCreate(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)
	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return err
	}

	options := &appconfigurationv1.CreatePropertyOptions{}

	options.SetEnvironmentID(d.Get("environment_id").(string))
	options.SetName(d.Get("name").(string))
	options.SetPropertyID(d.Get("property_id").(string))
	options.SetType(d.Get("type").(string))

	if _, ok := d.GetOk("description"); ok {
		options.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("tags"); ok {
		options.SetTags(d.Get("tags").(string))
	}
	if _, ok := d.GetOk("segment_rules"); ok {
		var segmentRules []appconfigurationv1.SegmentRule
		for _, e := range d.Get("segment_rules").([]interface{}) {
			value := e.(map[string]interface{})
			segmentRulesItem, err := resourceIbmAppConfigMapToSegmentRule(d, value)
			if err != nil {
				return err
			}
			segmentRules = append(segmentRules, segmentRulesItem)
		}
		options.SetSegmentRules(segmentRules)
	}
	if _, ok := d.GetOk("collections"); ok {
		var collections []appconfigurationv1.CollectionRef
		for _, e := range d.Get("collections").([]interface{}) {
			value := e.(map[string]interface{})
			collectionsItem := resourceIbmAppConfigMapToCollections(value)
			collections = append(collections, collectionsItem)
		}
		options.SetCollections(collections)
	}

	result, response, err := appconfigClient.CreateProperty(options)
	if err != nil {
		log.Printf("[DEBUG] CreateProperty failed %s\n%s", err, response)
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", guid, *options.EnvironmentID, *result.PropertyID))

	return resourceIbmAppConfigPropertyRead(d, meta)
}

func resourceIbmAppConfigPropertyRead(d *schema.ResourceData, meta interface{}) error {
	parts, err := idParts(d.Id())
	if err != nil {
		return nil
	}

	appconfigClient, err := getAppConfigClient(meta, parts[0])
	if err != nil {
		return err
	}

	options := &appconfigurationv1.GetPropertyOptions{}

	options.SetEnvironmentID(parts[1])
	options.SetPropertyID(parts[2])

	property, response, err := appconfigClient.GetProperty(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetProperty failed %s\n%s", err, response)
		return err
	}

	d.Set("guid", parts[0])
	d.Set("environment_id", parts[1])

	if property.Name != nil {
		if err = d.Set("name", property.Name); err != nil {
			return fmt.Errorf("error setting name: %s", err)
		}
	}
	if property.PropertyID != nil {
		if err = d.Set("property_id", property.PropertyID); err != nil {
			return fmt.Errorf("error setting property_id: %s", err)
		}
	}
	if property.Type != nil {
		if err = d.Set("type", property.Type); err != nil {
			return fmt.Errorf("error setting type: %s", err)
		}
	}
	if property.Value != nil {
		if err = d.Set("value", property.Value); err != nil {
			return fmt.Errorf("error setting value: %s", err)
		}
	}
	if property.Description != nil {
		if err = d.Set("description", property.Description); err != nil {
			return fmt.Errorf("error setting description: %s", err)
		}
	}
	if property.Tags != nil {
		if err = d.Set("tags", property.Tags); err != nil {
			return fmt.Errorf("error setting tags: %s", err)
		}
	}
	if property.SegmentExists != nil {
		if err = d.Set("segment_exists", property.SegmentExists); err != nil {
			return fmt.Errorf("error setting segment_exists: %s", err)
		}
	}
	if property.CreatedTime != nil {
		if err = d.Set("created_time", property.CreatedTime.String()); err != nil {
			return fmt.Errorf("error setting created_time: %s", err)
		}
	}
	if property.UpdatedTime != nil {
		if err = d.Set("updated_time", property.UpdatedTime.String()); err != nil {
			return fmt.Errorf("error setting updated_time: %s", err)
		}
	}
	if property.Href != nil {
		if err = d.Set("href", property.Href); err != nil {
			return fmt.Errorf("error setting href: %s", err)
		}
	}

	if property.SegmentRules != nil {
		segmentRules := []map[string]interface{}{}
		for _, segmentRulesItem := range property.SegmentRules {
			segmentRulesItemMap := resourceIbmAppConfigSegmentRuleToMap(segmentRulesItem)
			segmentRules = append(segmentRules, segmentRulesItemMap)
		}
		if err = d.Set("segment_rules", segmentRules); err != nil {
			return fmt.Errorf("error setting segment_rules: %s", err)
		}
	}
	if property.Collections != nil {
		collections := []map[string]interface{}{}
		for _, collectionsItem := range property.Collections {
			collectionsItemMap := resourceIbmAppConfigCollectionToMap(collectionsItem)
			collections = append(collections, collectionsItemMap)
		}
		if err = d.Set("collections", collections); err != nil {
			return fmt.Errorf("error setting collections: %s", err)
		}
	}
	return nil
}

func resourceIbmAppConfigPropertyUpdate(d *schema.ResourceData, meta interface{}) error {
	appConfigurationClient, err := meta.(ClientSession).AppConfigurationV1()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	options := &appconfigurationv1.UpdatePropertyOptions{}

	options.SetEnvironmentID(parts[0])
	options.SetPropertyID(parts[1])

	hasChange := false

	if d.HasChange("environment_id") {
		options.SetEnvironmentID(d.Get("environment_id").(string))
		hasChange = true
	}
	if d.HasChange("name") {
		options.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("property_id") {
		options.SetPropertyID(d.Get("property_id").(string))
		hasChange = true
	}
	if d.HasChange("value") {
		hasChange = true
	}
	if d.HasChange("description") {
		options.SetDescription(d.Get("description").(string))
		hasChange = true
	}
	if d.HasChange("tags") {
		options.SetTags(d.Get("tags").(string))
		hasChange = true
	}
	if d.HasChange("segment_rules") {
		// TODO: handle SegmentRules of type TypeList -- not primitive, not model
		hasChange = true
	}
	if d.HasChange("collections") {
		// TODO: handle Collections of type TypeList -- not primitive, not model
		hasChange = true
	}

	if hasChange {
		_, response, err := appConfigurationClient.UpdateProperty(options)
		if err != nil {
			log.Printf("[DEBUG] PatchProperty failed %s\n%s", err, response)
			return err
		}
	}

	return resourceIbmAppConfigPropertyRead(d, meta)
}

func resourceIbmAppConfigPropertyDelete(d *schema.ResourceData, meta interface{}) error {
	parts, err := idParts(d.Id())
	if err != nil {
		return nil
	}
	appconfigClient, err := getAppConfigClient(meta, parts[0])
	if err != nil {
		return err
	}

	options := &appconfigurationv1.DeletePropertyOptions{}

	options.SetEnvironmentID(parts[0])
	options.SetPropertyID(parts[1])

	response, err := appconfigClient.DeleteProperty(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] DeleteProperty failed %s\n%s", err, response)
		return err
	}

	d.SetId("")

	return nil
}
