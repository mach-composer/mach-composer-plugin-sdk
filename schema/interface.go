package schema

type MachComposerPlugin interface {
	Identifier() string

	IsEnabled() bool

	Configure(environment string, provider string) error

	GetValidationSchema() (*ValidationSchema, error)

	//Deprecated: this method is not used anymore in CLI 2.10.0 or newer. Instead, state is managed by the CLI
	SetRemoteStateBackend(data map[string]any) error

	SetGlobalConfig(data map[string]any) error

	SetSiteConfig(site string, data map[string]any) error

	SetSiteComponentConfig(site string, component string, data map[string]any) error

	SetSiteEndpointConfig(site string, name string, data map[string]any) error

	SetComponentConfig(component string, data map[string]any) error

	SetComponentEndpointsConfig(component string, data map[string]string) error

	//Deprecated: this method is not used anymore in CLI 2.10.0 or newer. Instead, state is managed by the CLI
	RenderTerraformStateBackend(site string) (string, error)

	RenderTerraformProviders(site string) (string, error)

	RenderTerraformResources(site string) (string, error)

	RenderTerraformComponent(site string, component string) (*ComponentSchema, error)
}
