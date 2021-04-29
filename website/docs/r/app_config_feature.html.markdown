---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : app_config_feature'
description: |-
  Manages Feature flag.
---

# ibm_app_config_feature

Provides a resource for Feature flag. This allows Feature flag to be created, updated and deleted.

## Example Usage

```hcl
resource "app_config_feature" "app_config_feature" {
  guid = "guid"
  environment_id = "environment_id"
  name = "name"
  feature_id = "feature_id"
  type = "type"
  enabled_value = true
  disabled_value = false
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, string) Environment Id.
- `name` - (Required, string) Feature name.
- `feature_id` - (Required, string) Feature id.
- `type` - (Required, string) Type of the feature (BOOLEAN, STRING, NUMERIC).
- `enabled_value` - (Required, bool) Value of the feature when it is enabled.
- `disabled_value` - (Required, bool) Value of the feature when it is disabled.
- `description` - (Optional, string) Feature description.
- `tags` - (Optional, string) Tags associated with the feature.
- `segment_rules` - (Optional, List) Segment Rules array.
  - `rules` - (Required, []interface{}) Rules array.
  - `value` - (Required, string) Value of the segment.
  - `order` - (Required, int) Order of the segment, used during evaluation.
- `collections` - (Optional, List) Collection array.
  - `collection_id` - (Required, string) Collection id.
  - `name` - (Required, string) Name of the collection.

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the Feature flag resource.
