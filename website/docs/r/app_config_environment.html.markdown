---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : app_config_environment'
description: |-
  Manages environment.
---

# ibm_app_config_feature

Provides a resource for environment. This allows environment to be created, updated and deleted.

## Example Usage

```hcl
resource "app_config_feature" "app_config_feature" {
  guid = "guid"
  environment_id = "environment_id"
  name = "name"
  description = "description"
  tags = "tags"
  color_code = "color_code"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, string) Environment Id.
- `name` - (Required, string) Feature name.
- `description` - (Optional, string) Feature description.
- `tags` - (Optional, string) Tags associated with the feature.
- `color_code` - (Optional, string) Color code to distinguish the environment.

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the environment resource.
