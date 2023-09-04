package schema

// PluginSchema is the schema implementation for a plugin. Together with protocol.Adapter this allows for the
// selective implementation of methods
type PluginSchema struct {
	Identifier string

	Configure           func(environment, provider string) error
	GetValidationSchema func() (*ValidationSchema, error)

	SetGlobalConfig        func(meta Meta, data map[string]any) error
	SetSiteConfig          func(site string, meta Meta, data map[string]any) error
	SetSiteComponentConfig func(site string, component string, meta Meta, data map[string]any) error
	SetSiteEndpointConfig  func(site string, name string, meta Meta, data map[string]any) error
	//TODO where do we want the version to go? (Specifically for sentry)
	SetComponentConfig          func(component, version string, meta Meta, data map[string]any) error
	SetComponentEndpointsConfig func(component string, meta Meta, endpoints map[string]string) error

	RenderTerraformProviders func(site string) (string, error)
	RenderTerraformResources func(site string) (string, error)
	RenderTerraformComponent func(site string, component string) (*ComponentSchema, error)
}
