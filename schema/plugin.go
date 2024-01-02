package schema

// MachComposerPlugin is the interface that plugins must implement. Note that in combination with PluginSchema you can omit
// functions that you don't need
type MachComposerPlugin interface {

	// Identifier returns the identifier of the plugin
	Identifier() string

	// Configure is called when the plugin is loaded
	Configure(environment string, provider string) error

	// GetValidationSchema returns the validation schema for the plugin
	GetValidationSchema() (*ValidationSchema, error)

	// SetGlobalConfig is called when the global config is set
	SetGlobalConfig(data map[string]any) error

	// SetSiteConfig is called when the site config is set
	SetSiteConfig(site string, data map[string]any) error

	// SetSiteComponentConfig is called when the site component config is set
	SetSiteComponentConfig(site string, component string, data map[string]any) error

	// SetSiteEndpointConfig is called when the site endpoint config is set
	//
	// Deprecated: move endpoints to the terraform module instead
	SetSiteEndpointConfig(site string, name string, data map[string]any) error

	// SetComponentConfig is called when the component config is set
	SetComponentConfig(component, version string, data map[string]any) error

	// SetComponentEndpointsConfig is called when the component endpoints config is set
	//
	// Deprecated: move endpoints to the terraform module instead
	SetComponentEndpointsConfig(component string, data map[string]string) error

	// RenderTerraformProviders returns the terraform providers for the site
	RenderTerraformProviders(site string) (string, error)

	// RenderTerraformResources returns the terraform resources for the site
	RenderTerraformResources(site string) (string, error)

	// RenderTerraformComponent returns the terraform component for the site
	RenderTerraformComponent(site string, component string) (*ComponentSchema, error)
}
