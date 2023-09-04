package protocol

import (
	"github.com/mach-composer/mach-composer-plugin-sdk/v1/schema"
)

type ConfigureInput struct {
	Environment string
	Provider    string
}

type GetValidationSchemaOutput struct {
	Result schema.ValidationSchema
	Err    *PluginError
}

type SetGlobalConfigInput struct {
	Meta schema.Meta
	Data map[string]any
}

type SetSiteComponentConfigInput struct {
	Version   string
	Site      string
	Component string
	Meta      schema.Meta
	Data      map[string]any
}

type SetSiteComponentConfigOutput struct {
	Err *PluginError
}

type SetSiteEndpointsConfigInput struct {
	Site string
	Name string
	Meta schema.Meta
	Data map[string]any
}

type SetSiteEndpointsConfigOutput struct {
	Err *PluginError
}

type SetComponentConfigInput struct {
	Component string
	Meta      schema.Meta
	Data      map[string]any
}

type SetComponentConfigOutput struct {
	Err *PluginError
}

type SetComponentEndpointsConfigInput struct {
	Component string
	Meta      schema.Meta
	Endpoints map[string]string
}

type SetComponentEndpointsConfigOutput struct {
	Err *PluginError
}

type SetSiteConfigInput struct {
	Name string
	Meta schema.Meta
	Data map[string]any
}

type RenderTerraformProvidersInput struct {
	Site string
}

type RenderTerraformProvidersOutput struct {
	Result string
	Err    *PluginError
}

type RenderTerraformResourcesInput struct {
	Site string
}

type RenderTerraformResourcesOutput struct {
	Result string
	Err    *PluginError
}

type RenderTerraformComponentInput struct {
	Site      string
	Component string
}

type RenderTerraformComponentOutput struct {
	Result *schema.ComponentSchema
	Err    *PluginError
}

type ErrorOutput struct {
	Err *PluginError
}
