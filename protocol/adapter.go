package protocol

import (
	"github.com/hashicorp/go-hclog"

	"github.com/mach-composer/mach-composer-plugin-sdk/schema"
)

type Adapter struct {
	Logger hclog.Logger
	fn     *schema.PluginSchema
}

func NewAdapter(s *schema.PluginSchema, logger hclog.Logger) *Adapter {
	return &Adapter{
		Logger: logger,
		fn:     s,
	}
}

func (c *Adapter) SetLogger(logger hclog.Logger) {
	c.Logger = logger
}

func (c *Adapter) Configure(environment, provider string) error {
	if c.fn.Configure != nil {
		return c.fn.Configure(environment, provider)
	}
	return nil
}

func (c *Adapter) Identifier() string {
	return c.fn.Identifier
}

func (p *Adapter) IsEnabled() bool {
	if p.fn.IsEnabled != nil {
		return p.fn.IsEnabled()
	}
	return true
}

func (p *Adapter) SetRemoteStateBackend(data map[string]any) error {
	if p.fn.SetRemoteStateBackend != nil {
		return p.fn.SetRemoteStateBackend(data)
	}
	return nil
}

func (p *Adapter) SetGlobalConfig(data map[string]any) error {
	if p.fn.SetGlobalConfig != nil {
		return p.fn.SetGlobalConfig(data)
	}
	return nil
}
func (p *Adapter) SetSiteConfig(site string, data map[string]any) error {
	if p.fn.SetSiteConfig != nil {
		return p.fn.SetSiteConfig(site, data)
	}
	return nil
}

func (p *Adapter) SetSiteComponentConfig(site string, component string, data map[string]any) error {
	if p.fn.SetSiteComponentConfig != nil {
		return p.fn.SetSiteComponentConfig(site, component, data)
	}
	return nil
}

func (p *Adapter) SetSiteEndpointsConfig(site string, data map[string]any) error {
	if p.fn.SetSiteEndpointsConfig != nil {
		return p.fn.SetSiteEndpointsConfig(site, data)
	}
	return nil
}

func (p *Adapter) SetComponentConfig(component string, data map[string]any) error {
	if p.fn.SetComponentConfig != nil {
		return p.fn.SetComponentConfig(component, data)
	}
	return nil
}

func (p *Adapter) SetComponentEndpointsConfig(component string, endpoints map[string]string) error {
	if p.fn.SetComponentEndpointsConfig != nil {
		return p.fn.SetComponentEndpointsConfig(component, endpoints)
	}
	return nil
}
func (p *Adapter) RenderTerraformStateBackend(site string) (string, error) {
	if p.fn.RenderTerraformStateBackend != nil {
		return p.fn.RenderTerraformStateBackend(site)
	}
	return "", nil
}

func (p *Adapter) RenderTerraformProviders(site string) (string, error) {
	if p.fn.RenderTerraformProviders != nil {
		return p.fn.RenderTerraformProviders(site)
	}
	return "", nil
}

func (p *Adapter) RenderTerraformResources(site string) (string, error) {
	if p.fn.RenderTerraformResources != nil {
		return p.fn.RenderTerraformResources(site)
	}
	return "", nil
}

func (p *Adapter) RenderTerraformComponent(site, component string) (*schema.ComponentSchema, error) {
	if p.fn.RenderTerraformComponent != nil {
		return p.fn.RenderTerraformComponent(site, component)
	}
	return nil, nil
}
