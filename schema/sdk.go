package schema

type ComponentSchema struct {
	Resources string
	Variables string
	DependsOn []string
	Providers []string
}

type ValidationSchema struct {
	GlobalConfigSchema             map[string]any
	RemoteStateSchema              map[string]any
	SiteConfigSchema               map[string]any
	SiteComponentConfigSchema      map[string]any
	SiteEndpointsConfig            map[string]any
	ComponentConfigSchema          map[string]any
	ComponentEndpointsConfigSchema map[string]any
}

type PluginSchema struct {
	Identifier string

	Configure             func(environment, provider string) error
	IsEnabled             func() bool
	GetValidationSchema   func() (*ValidationSchema, error)
	SetRemoteStateBackend func(data map[string]any) error

	SetGlobalConfig             func(data map[string]any) error
	SetSiteConfig               func(site string, data map[string]any) error
	SetSiteComponentConfig      func(site string, component string, data map[string]any) error
	SetSiteEndpointsConfig      func(site string, data map[string]any) error
	SetComponentConfig          func(component string, data map[string]any) error
	SetComponentEndpointsConfig func(component string, endpoints map[string]string) error

	RenderTerraformStateBackend func(site string) (string, error)
	RenderTerraformProviders    func(site string) (string, error)
	RenderTerraformResources    func(site string) (string, error)
	RenderTerraformComponent    func(site string, component string) (*ComponentSchema, error)
}
