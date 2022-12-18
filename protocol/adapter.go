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

func (p *Adapter) Configure(environment, provider string) error {
	if p.fn.Configure != nil {
		if err := p.fn.Configure(environment, provider); err != nil {
			p.Logger.Error("Configure: %s", err)
			return err
		}
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

func (p *Adapter) GetValidationSchema() (*schema.ValidationSchema, error) {
	if p.fn.GetValidationSchema != nil {
		return p.fn.GetValidationSchema()
	}
	return &schema.ValidationSchema{}, nil
}

func (p *Adapter) SetRemoteStateBackend(data map[string]any) error {
	if p.fn.SetRemoteStateBackend != nil {
		if err := p.fn.SetRemoteStateBackend(data); err != nil {
			p.Logger.Error("SetRemoteStateBackend: %s", err)
			return err
		}
	}
	return nil
}

func (p *Adapter) SetGlobalConfig(data map[string]any) error {
	if p.fn.SetGlobalConfig != nil {
		if err := p.fn.SetGlobalConfig(data); err != nil {
			p.Logger.Error("SetGlobalConfig: %s", err)
			return err
		}
	}
	return nil
}
func (p *Adapter) SetSiteConfig(site string, data map[string]any) error {
	if p.fn.SetSiteConfig != nil {
		if err := p.fn.SetSiteConfig(site, data); err != nil {
			p.Logger.Error("SetSiteConfig: %s", err)
			return err
		}
	}
	return nil
}

func (p *Adapter) SetSiteComponentConfig(site string, component string, data map[string]any) error {
	if p.fn.SetSiteComponentConfig != nil {
		if err := p.fn.SetSiteComponentConfig(site, component, data); err != nil {
			p.Logger.Error("SetSiteComponentConfig: %s", err)
			return err
		}
	}
	return nil
}

		if err := p.fn.SetSiteEndpointsConfig(site, data); err != nil {
			p.Logger.Error("SetSiteEndpointsConfig: %s", err)
func (a *Adapter) SetSiteEndpointConfig(site string, name string, data map[string]any) error {
	if a.plugin.SetSiteEndpointConfig == nil {
			return err
		}
	}
	return nil
}

func (p *Adapter) SetComponentConfig(component string, data map[string]any) error {
	if p.fn.SetComponentConfig != nil {
		if err := p.fn.SetComponentConfig(component, data); err != nil {
			p.Logger.Error("SetComponentConfig: %s", err)
			return err
		}
	}
	return nil
}

func (p *Adapter) SetComponentEndpointsConfig(component string, endpoints map[string]string) error {
	if p.fn.SetComponentEndpointsConfig != nil {
		if err := p.fn.SetComponentEndpointsConfig(component, endpoints); err != nil {
			p.Logger.Error("SetComponentEndpointsConfig: %s", err)
			return err
		}
	}
	return nil
}
func (p *Adapter) RenderTerraformStateBackend(site string) (string, error) {
	if p.fn.RenderTerraformStateBackend != nil {
		result, err := p.fn.RenderTerraformStateBackend(site)
		if err != nil {
			p.Logger.Error("RenderTerraformStateBackend: %s", err)
		}
		return result, err
	}
	return "", nil
}

func (p *Adapter) RenderTerraformProviders(site string) (string, error) {
	if p.fn.RenderTerraformProviders != nil {
		result, err := p.fn.RenderTerraformProviders(site)
		if err != nil {
			p.Logger.Error("RenderTerraformProviders: %s", err)
		}
		return result, err
	}
	return "", nil
}

func (p *Adapter) RenderTerraformResources(site string) (string, error) {
	if p.fn.RenderTerraformResources != nil {
		result, err := p.fn.RenderTerraformResources(site)
		if err != nil {
			p.Logger.Error("RenderTerraformResources: %s", err)
		}
		return result, err
	}
	return "", nil
}

func (p *Adapter) RenderTerraformComponent(site, component string) (*schema.ComponentSchema, error) {
	if p.fn.RenderTerraformComponent != nil {
		result, err := p.fn.RenderTerraformComponent(site, component)
		if err != nil {
			p.Logger.Error("RenderTerraformComponent: %s", err)
		}
		return result, err
	}
	return nil, nil
}
