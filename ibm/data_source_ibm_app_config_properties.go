// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func dataSourceIbmAppConfigProperties() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmAppConfigPropertiesRead,

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
			"sort": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sort the feature details based on the specified attribute.",
			},
			"tags": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter the resources to be returned based on the associated tags. Specify the parameter as a list of comma separated tags. Returns resources associated with any of the specified tags.",
			},
			"collections": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter features by a list of comma separated collections.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"segments": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter features by a list of comma separated segments.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"expand": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to `true`, returns expanded view of the resource details.",
			},
			"includes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Include the associated collections or targeting rules details in the response.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.",
			},
			"properties": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of properties.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Property name.",
						},
						"property_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Property id.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Property description.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the Property (BOOLEAN, STRING, NUMERIC).",
						},
						"value": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Value of the Property. The value can be Boolean, String or a Numeric value as per the `type` attribute.",
						},
						"tags": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Tags associated with the property.",
						},
						"segment_rules": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specify the targeting rules that is used to set different property values for different segments.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rules": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Rules array.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"segments": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "List of segment ids that are used for targeting using the rule.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"value": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "Value to be used for evaluation for this rule. The value can be Boolean, String or a Numeric value as per the `type` attribute.",
									},
									"order": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation.",
									},
								},
							},
						},
						"segment_exists": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Denotes if the targeting rules are specified for the property.",
						},
						"collections": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of collection id representing the collections that are associated with the specified property.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"collection_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Collection id.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the collection.",
									},
								},
							},
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
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of records returned in the current response.",
			},
			"next": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the next list of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"first": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the first page of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"previous": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the previous list of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"last": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the last page of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmAppConfigPropertiesRead(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return err
	}

	options := &appconfigurationv1.ListPropertiesOptions{}

	options.SetEnvironmentID(d.Get("environment_id").(string))

	if _, ok := d.GetOk("expand"); ok {
		options.SetExpand(d.Get("expand").(bool))
	}
	if _, ok := d.GetOk("sort"); ok {
		options.SetSort(d.Get("sort").(string))
	}
	if _, ok := d.GetOk("tags"); ok {
		options.SetTags(d.Get("tags").(string))
	}
	if _, ok := d.GetOk("collections"); ok {
		collections := []string{}
		for _, item := range d.Get("collections").([]interface{}) {
			collections = append(collections, item.(string))
		}
		options.SetCollections(collections)
	}
	if _, ok := d.GetOk("segments"); ok {
		segments := []string{}
		for _, item := range d.Get("segments").([]interface{}) {
			segments = append(segments, item.(string))
		}
		options.SetSegments(segments)
	}
	if _, ok := d.GetOk("includes"); ok {
		includes := []string{}
		for _, item := range d.Get("includes").([]interface{}) {
			includes = append(includes, item.(string))
		}
		options.SetInclude(includes)
	}

	var propertiesList *appconfigurationv1.PropertiesList
	var offset int64
	var limit int64 = 10
	var isLimit bool
	finalList := []appconfigurationv1.Property{}
	if _, ok := d.GetOk("limit"); ok {
		isLimit = true
		limit = int64(d.Get("limit").(int))
	}
	options.SetLimit(limit)
	if _, ok := d.GetOk("offset"); ok {
		offset = int64(d.Get("offset").(int))
	}
	for {
		options.Offset = &offset
		result, response, err := appconfigClient.ListProperties(options)
		propertiesList = result
		if err != nil {
			log.Printf("[DEBUG] ListProperties failed %s\n%s", err, response)
			return err
		}
		if isLimit {
			offset = 0
		} else {
			offset = dataSourceAppConnfigGetNext(result.Next)
		}
		finalList = append(finalList, result.Properties...)
		if offset == 0 {
			break
		}
	}

	propertiesList.Properties = finalList

	d.SetId(fmt.Sprintf("%s/%s", guid, *options.EnvironmentID))

	if propertiesList.Properties != nil {
		err = d.Set("properties", getAppConfigPropertiesResponse(propertiesList.Properties))
		if err != nil {
			return fmt.Errorf("error setting properties %s", err)
		}
	}
	if propertiesList.TotalCount != nil {
		if err = d.Set("total_count", propertiesList.TotalCount); err != nil {
			return fmt.Errorf("error setting total_count: %s", err)
		}
	}
	if propertiesList.Limit != nil {
		if err = d.Set("limit", propertiesList.Limit); err != nil {
			return fmt.Errorf("error setting limit: %s", err)
		}
	}
	if propertiesList.Offset != nil {
		if err = d.Set("offset", propertiesList.Offset); err != nil {
			return fmt.Errorf("error setting offset: %s", err)
		}
	}
	if propertiesList.First != nil {
		err = d.Set("first", dataSourceAppConfigFlattenPagination(*propertiesList.First))
		if err != nil {
			return fmt.Errorf("error setting first %s", err)
		}
	}

	if propertiesList.Previous != nil {
		err = d.Set("previous", dataSourceAppConfigFlattenPagination(*propertiesList.Previous))
		if err != nil {
			return fmt.Errorf("error setting previous %s", err)
		}
	}

	if propertiesList.Last != nil {
		err = d.Set("last", dataSourceAppConfigFlattenPagination(*propertiesList.Last))
		if err != nil {
			return fmt.Errorf("error setting last %s", err)
		}
	}
	if propertiesList.Next != nil {
		err = d.Set("next", dataSourceAppConfigFlattenPagination(*propertiesList.Next))
		if err != nil {
			return fmt.Errorf("error setting next %s", err)
		}
	}

	return nil
}

func getAppConfigPropertiesResponse(result []appconfigurationv1.Property) (properties []map[string]interface{}) {
	for _, item := range result {
		properties = append(properties, handleAppConfigPropertyResponse(item))
	}
	return properties
}

func handleAppConfigPropertyResponse(property appconfigurationv1.Property) (propertyMap map[string]interface{}) {
	propertyMap = map[string]interface{}{}

	if property.Name != nil {
		propertyMap["name"] = property.Name
	}
	if property.PropertyID != nil {
		propertyMap["property_id"] = property.PropertyID
	}
	if property.Description != nil {
		propertyMap["description"] = property.Description
	}
	if property.Type != nil {
		propertyMap["type"] = property.Type
	}
	if property.Value != nil {
		propertyMap["value"] = property.Value
	}
	if property.Tags != nil {
		propertyMap["tags"] = property.Tags
	}
	if property.SegmentExists != nil {
		propertyMap["segment_exists"] = property.SegmentExists
	}
	if property.UpdatedTime != nil {
		propertyMap["updated_time"] = property.UpdatedTime
	}
	if property.CreatedTime != nil {
		propertyMap["created_time"] = property.CreatedTime
	}
	if property.Href != nil {
		propertyMap["href"] = property.Href
	}
	if property.Collections != nil {
		propertyMap["collections"] = getAppConfigCollectionResponse(property.Collections)
	}
	if property.SegmentRules != nil {
		propertyMap["segment_rules"] = getAppConfigSegmentResponse(property.SegmentRules)
	}
	return propertyMap
}
