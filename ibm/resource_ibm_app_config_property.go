// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"strconv"

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

	options.SetName(d.Get("name").(string))
	options.SetType(d.Get("type").(string))
	options.SetEnvironmentID(d.Get("environment_id").(string))
	options.SetPropertyID(d.Get("property_id").(string))
	value := d.Get("value").(string)
	v, err := getAppConfigValueProperty(d, value)
	if err != nil {
		return err
	}
	options.SetValue(v)

	if _, ok := d.GetOk("description"); ok {
		options.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("tags"); ok {
		options.SetTags(d.Get("tags").(string))
	}
	if _, ok := d.GetOk("collections"); ok {
		data := getAppConfiigCollectionInput(d)
		options.SetCollections(data)
	}
	if _, ok := d.GetOk("segment_rules"); ok {
		data, err := getAppConfiigSegmentRuleInput(d)
		if err != nil {
			return err
		}
		options.SetSegmentRules(data)
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

	result, response, err := appconfigClient.GetProperty(options)
	if err != nil {
		log.Printf("[DEBUG] GetProperty failed %s\n%s", err, response)
		return err
	}

	d.Set("guid", parts[0])
	d.Set("environment_id", parts[1])

	if result.Name != nil {
		if err = d.Set("name", result.Name); err != nil {
			return fmt.Errorf("error setting name: %s", err)
		}
	}
	if result.PropertyID != nil {
		if err = d.Set("property_id", result.PropertyID); err != nil {
			return fmt.Errorf("error setting property_id: %s", err)
		}
	}
	if result.Type != nil {
		if err = d.Set("type", result.Type); err != nil {
			return fmt.Errorf("error setting type: %s", err)
		}
	}
	if result.Value != nil {
		value := result.Value
		switch value.(interface{}).(type) {
		case string:
			d.Set("value", value.(string))
		case float64:
			d.Set("value", fmt.Sprintf("%v", value))
		case bool:
			d.Set("value", strconv.FormatBool(value.(bool)))
		}
	}
	if result.Description != nil {
		if err = d.Set("description", result.Description); err != nil {
			return fmt.Errorf("error setting description: %s", err)
		}
	}
	if result.Tags != nil {
		if err = d.Set("tags", result.Tags); err != nil {
			return fmt.Errorf("error setting tags: %s", err)
		}
	}
	if result.SegmentExists != nil {
		if err = d.Set("segment_exists", result.SegmentExists); err != nil {
			return fmt.Errorf("error setting segment_exists: %s", err)
		}
	}
	if result.CreatedTime != nil {
		if err = d.Set("created_time", result.CreatedTime.String()); err != nil {
			return fmt.Errorf("error setting created_time: %s", err)
		}
	}
	if result.UpdatedTime != nil {
		if err = d.Set("updated_time", result.UpdatedTime.String()); err != nil {
			return fmt.Errorf("error setting updated_time: %s", err)
		}
	}
	if result.Href != nil {
		if err = d.Set("href", result.Href); err != nil {
			return fmt.Errorf("error setting href: %s", err)
		}
	}

	if result.SegmentRules != nil {
		segmentRules := getAppConfigSegmentResponse(result.SegmentRules)
		if err = d.Set("segment_rules", segmentRules); err != nil {
			return fmt.Errorf("error setting segment_rules: %s", err)
		}
	}
	if result.Collections != nil {
		collections := getAppConfigCollectionResponse(result.Collections)
		if err = d.Set("collections", collections); err != nil {
			return fmt.Errorf("error setting collections: %s", err)
		}
	}
	return nil
}

func resourceIbmAppConfigPropertyUpdate(d *schema.ResourceData, meta interface{}) error {
	if ok := d.HasChanges("name", "value", "description", "tags", "segment_rules", "collections"); ok {
		parts, err := idParts(d.Id())
		if err != nil {
			return nil
		}
		appconfigClient, err := getAppConfigClient(meta, parts[0])
		if err != nil {
			return err
		}
		options := &appconfigurationv1.UpdatePropertyOptions{}

		options.SetEnvironmentID(parts[1])
		options.SetPropertyID(parts[2])

		options.SetName(d.Get("name").(string))
		value := d.Get("value").(string)
		v, err := getAppConfigValueProperty(d, value)
		if err != nil {
			return err
		}
		options.SetValue(v)

		if _, ok := d.GetOk("description"); ok {
			options.SetDescription(d.Get("description").(string))
		}
		if _, ok := d.GetOk("tags"); ok {
			options.SetTags(d.Get("tags").(string))
		}
		if _, ok := d.GetOk("collections"); ok {
			data := getAppConfiigCollectionInput(d)
			options.SetCollections(data)
		}
		if _, ok := d.GetOk("segment_rules"); ok {
			data, err := getAppConfiigSegmentRuleInput(d)
			if err != nil {
				return err
			}
			options.SetSegmentRules(data)
		}
		_, response, err := appconfigClient.UpdateProperty(options)
		if err != nil {
			log.Printf("[DEBUG] PatchProperty failed %s\n%s", err, response)
			return err
		}

		return resourceIbmAppConfigPropertyRead(d, meta)

	}
	return nil
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
