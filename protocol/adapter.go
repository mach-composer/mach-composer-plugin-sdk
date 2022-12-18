package protocol

import (
	"fmt"

	"github.com/hashicorp/go-hclog"

	"github.com/mach-composer/mach-composer-plugin-sdk/schema"
)

// Adapter wraps a PluginSchema
type Adapter struct {
	name   string
	plugin *schema.PluginSchema
	schema *schema.ValidationSchema
	logger hclog.Logger
}

var _ schema.MachComposerPlugin = (*Adapter)(nil)

func NewAdapter(plugin *schema.PluginSchema, logger hclog.Logger) *Adapter {
	var schema *schema.ValidationSchema
	if plugin.GetValidationSchema != nil {
		var err error
		schema, err = plugin.GetValidationSchema()
		if err != nil {
			panic(err)
		}
	}

	return &Adapter{
		name:   plugin.Identifier,
		plugin: plugin,
		schema: schema,
		logger: logger,
	}
}

func (a *Adapter) Configure(environment, provider string) error {
	if a.plugin.Configure != nil {
		if err := a.plugin.Configure(environment, provider); err != nil {
			a.logger.Debug("Configure error: %s", err)
			return err
		}
	}
	return nil
}

func (a *Adapter) Identifier() string {
	return a.plugin.Identifier
}

func (a *Adapter) IsEnabled() bool {
	if a.plugin.IsEnabled != nil {
		return a.plugin.IsEnabled()
	}
	return true
}

func (a *Adapter) GetValidationSchema() (*schema.ValidationSchema, error) {
	if a.plugin.GetValidationSchema != nil {
		return a.plugin.GetValidationSchema()
	}
	return &schema.ValidationSchema{}, nil
}

func (a *Adapter) SetRemoteStateBackend(data map[string]any) error {
	if a.plugin.SetRemoteStateBackend == nil {
		return nil
	}

	if a.schema != nil {
		if err := a.schema.Validate(a.schema.RemoteStateSchema, data); err != nil {
			return err
		}
	}

	if err := a.plugin.SetRemoteStateBackend(data); err != nil {
		a.logger.Debug(fmt.Sprintf("SetRemoteStateBackend error: %s", err))
		return err
	}
	return nil
}

func (a *Adapter) SetGlobalConfig(data map[string]any) error {
	if a.plugin.SetGlobalConfig == nil {
		return nil
	}

	if a.schema != nil && data != nil {
		if err := a.schema.Validate(a.schema.GlobalConfigSchema, data); err != nil {
			return err
		}
	}

	if err := a.plugin.SetGlobalConfig(data); err != nil {
		a.logger.Debug(fmt.Sprintf("SetGlobalConfig error: %s", err))
		return err
	}
	return nil
}
func (a *Adapter) SetSiteConfig(site string, data map[string]any) error {
	if a.plugin.SetSiteConfig == nil {
		return nil
	}

	if a.schema != nil && data != nil {
		if err := a.schema.Validate(a.schema.SiteConfigSchema, data); err != nil {
			return err
		}
	}

	if err := a.plugin.SetSiteConfig(site, data); err != nil {
		a.logger.Debug(fmt.Sprintf("SetSiteConfig error: %s", err))
		return err
	}
	return nil
}

func (a *Adapter) SetSiteComponentConfig(site string, component string, data map[string]any) error {
	if a.plugin.SetSiteComponentConfig == nil {
		return nil
	}

	if a.schema != nil && data != nil {
		if err := a.schema.Validate(a.schema.SiteComponentConfigSchema, data); err != nil {
			return err
		}
	}

	if err := a.plugin.SetSiteComponentConfig(site, component, data); err != nil {
		a.logger.Debug(fmt.Sprintf("SetSiteComponentConfig error: %s", err))
		return err
	}
	return nil
}

func (a *Adapter) SetSiteEndpointConfig(site string, name string, data map[string]any) error {
	if a.plugin.SetSiteEndpointConfig == nil {
		return nil
	}

	if a.schema != nil && data != nil {
		if err := a.schema.Validate(a.schema.SiteEndpointConfig, data); err != nil {
			return err
		}
	}

	if err := a.plugin.SetSiteEndpointConfig(site, name, data); err != nil {
		a.logger.Debug(fmt.Sprintf("SetSiteEndpointConfig error: %s", err))
		return err
	}
	return nil
}

func (a *Adapter) SetComponentConfig(component string, data map[string]any) error {
	if a.plugin.SetComponentConfig == nil {
		return nil
	}

	if a.schema != nil {
		if err := a.schema.Validate(a.schema.ComponentConfigSchema, data); err != nil {
			return err
		}
	}

	if err := a.plugin.SetComponentConfig(component, data); err != nil {
		a.logger.Debug(fmt.Sprintf("SetComponentConfig error: %s", err))
		return err
	}
	return nil
}

func (a *Adapter) SetComponentEndpointsConfig(component string, endpoints map[string]string) error {
	if a.plugin.SetComponentEndpointsConfig != nil {
		if err := a.plugin.SetComponentEndpointsConfig(component, endpoints); err != nil {
			a.logger.Debug(fmt.Sprintf("SetComponentEndpointsConfig error: %s", err))
			return err
		}
	}
	return nil
}
func (a *Adapter) RenderTerraformStateBackend(site string) (string, error) {
	if a.plugin.RenderTerraformStateBackend != nil {
		result, err := a.plugin.RenderTerraformStateBackend(site)
		if err != nil {
			a.logger.Debug(fmt.Sprintf("RenderTerraformStateBackend error: %s", err))
		}
		return result, err
	}
	return "", nil
}

func (a *Adapter) RenderTerraformProviders(site string) (string, error) {
	if a.plugin.RenderTerraformProviders != nil {
		result, err := a.plugin.RenderTerraformProviders(site)
		if err != nil {
			a.logger.Debug(fmt.Sprintf("RenderTerraformProviders error: %s", err))
		}
		return result, err
	}
	return "", nil
}

func (a *Adapter) RenderTerraformResources(site string) (string, error) {
	if a.plugin.RenderTerraformResources != nil {
		result, err := a.plugin.RenderTerraformResources(site)
		if err != nil {
			a.logger.Debug("RenderTerraformResources error: %s", err)
		}
		return result, err
	}
	return "", nil
}

func (a *Adapter) RenderTerraformComponent(site, component string) (*schema.ComponentSchema, error) {
	if a.plugin.RenderTerraformComponent != nil {
		result, err := a.plugin.RenderTerraformComponent(site, component)
		if err != nil {
			a.logger.Debug("RenderTerraformComponent error: %s", err)
		}
		return result, err
	}
	return nil, nil
}
