---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : app_config_features'
description: |-
  Get information about Features flag
---

# ibm_app_config_features

Provides a read-only data source for all Features flag. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "app_config_features" "app_config_features" {
	guid = "guid"
	environment_id = "environment_id"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, string) Environment Id.

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the Features datasource.
- `features` - Array of Features. Nested `features` blocks have the following structure:

  - `name` - Feature name.
  - `feature_id` - Feature id.
  - `description` - Feature description.
  - `type` - Type of the feature (BOOLEAN, STRING, NUMERIC).
  - `enabled_value` - Value of the feature when it is enabled.
  - `disabled_value` - Value of the feature when it is disabled.
  - `tags` - Tags associated with the feature.
  - `segment_rules` - Segment Rules array. Nested `segment_rules` blocks have the following structure:
    - `rules` - Rules array. Nested `rules` blocks have the following structure:
      - `segments` - Segments array.
    - `value` - Value of the segment.
    - `order` - Order of the segment, used during evaluation.
  - `collections` - Collection array. Nested `collections` blocks have the following structure:
    - `collection_id` - Collection id.
    - `name` - Name of the collection.

- `page_info` Nested `page_info` blocks have the following structure:
  - `total_count` - total count of the records.
  - `count` - total page count.
  - `previous` - URL to navigate to the previous list of records.
  - `next` - URL to navigate to the next list of records.
