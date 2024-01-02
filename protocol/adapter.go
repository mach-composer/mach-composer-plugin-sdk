package protocol

import (
	"fmt"

	"github.com/hashicorp/go-hclog"

	"github.com/mach-composer/mach-composer-plugin-sdk/v2/schema"
)

// Adapter wraps a PluginSchema. It allows for omitting implementing methods as it will check if the method is set before
// executing it
type Adapter struct {
	name             string
	pluginSchema     *schema.PluginSchema
	validationSchema *schema.ValidationSchema
	logger           hclog.Logger
}

var _ schema.MachComposerPlugin = (*Adapter)(nil)

// NewAdapter creates a new Adapter
func NewAdapter(plugin *schema.PluginSchema, logger hclog.Logger) *Adapter {
	var validationSchema *schema.ValidationSchema
	if plugin.GetValidationSchema != nil {
		var err error
		validationSchema, err = plugin.GetValidationSchema()
		if err != nil {
			panic(err)
		}
	}

	return &Adapter{
		name:             plugin.Identifier,
		pluginSchema:     plugin,
		validationSchema: validationSchema,
		logger:           logger,
	}
}

func (a *Adapter) Configure(environment, provider string) error {
	if a.pluginSchema.Configure != nil {
		if err := a.pluginSchema.Configure(environment, provider); err != nil {
			a.logger.Debug("Configure error: %s", err)
			return err
		}
	}
	return nil
}

func (a *Adapter) Identifier() string {
	return a.pluginSchema.Identifier
}

func (a *Adapter) GetValidationSchema() (*schema.ValidationSchema, error) {
	if a.pluginSchema.GetValidationSchema != nil {
		return a.pluginSchema.GetValidationSchema()
	}
	return &schema.ValidationSchema{}, nil
}

func (a *Adapter) SetGlobalConfig(data map[string]any) error {
	if a.pluginSchema.SetGlobalConfig == nil {
		return nil
	}

	if a.validationSchema != nil && data != nil {
		if err := a.validationSchema.Validate(a.validationSchema.GlobalConfigSchema, data); err != nil {
			return err
		}
	}

	if err := a.pluginSchema.SetGlobalConfig(data); err != nil {
		a.logger.Debug(fmt.Sprintf("SetGlobalConfig error: %s", err))
		return err
	}
	return nil
}
func (a *Adapter) SetSiteConfig(site string, data map[string]any) error {
	if a.pluginSchema.SetSiteConfig == nil {
		return nil
	}

	if a.validationSchema != nil && data != nil {
		if err := a.validationSchema.Validate(a.validationSchema.SiteConfigSchema, data); err != nil {
			return err
		}
	}

	if err := a.pluginSchema.SetSiteConfig(site, data); err != nil {
		a.logger.Debug(fmt.Sprintf("SetSiteConfig error: %s", err))
		return err
	}
	return nil
}

func (a *Adapter) SetSiteComponentConfig(site string, component string, data map[string]any) error {
	if a.pluginSchema.SetSiteComponentConfig == nil {
		return nil
	}

	if a.validationSchema != nil && data != nil {
		if err := a.validationSchema.Validate(a.validationSchema.SiteComponentConfigSchema, data); err != nil {
			return err
		}
	}

	if err := a.pluginSchema.SetSiteComponentConfig(site, component, data); err != nil {
		a.logger.Debug(fmt.Sprintf("SetSiteComponentConfig error: %s", err))
		return err
	}
	return nil
}

func (a *Adapter) SetSiteEndpointConfig(site string, name string, data map[string]any) error {
	if a.pluginSchema.SetSiteEndpointConfig == nil {
		return nil
	}

	if a.validationSchema != nil && data != nil {
		if err := a.validationSchema.Validate(a.validationSchema.SiteEndpointConfig, data); err != nil {
			return err
		}
	}

	if err := a.pluginSchema.SetSiteEndpointConfig(site, name, data); err != nil {
		a.logger.Debug(fmt.Sprintf("SetSiteEndpointConfig error: %s", err))
		return err
	}
	return nil
}

func (a *Adapter) SetComponentConfig(component string, version string, data map[string]any) error {
	if a.pluginSchema.SetComponentConfig == nil {
		return nil
	}

	if a.validationSchema != nil {
		if err := a.validationSchema.Validate(a.validationSchema.ComponentConfigSchema, data); err != nil {
			return err
		}
	}

	if err := a.pluginSchema.SetComponentConfig(component, version, data); err != nil {
		a.logger.Debug(fmt.Sprintf("SetComponentConfig error: %s", err))
		return err
	}
	return nil
}

func (a *Adapter) SetComponentEndpointsConfig(component string, endpoints map[string]string) error {
	if a.pluginSchema.SetComponentEndpointsConfig != nil {
		if err := a.pluginSchema.SetComponentEndpointsConfig(component, endpoints); err != nil {
			a.logger.Debug(fmt.Sprintf("SetComponentEndpointsConfig error: %s", err))
			return err
		}
	}
	return nil
}

func (a *Adapter) RenderTerraformProviders(site string) (string, error) {
	if a.pluginSchema.RenderTerraformProviders != nil {
		result, err := a.pluginSchema.RenderTerraformProviders(site)
		if err != nil {
			a.logger.Debug(fmt.Sprintf("RenderTerraformProviders error: %s", err))
		}
		return result, err
	}
	return "", nil
}

func (a *Adapter) RenderTerraformResources(site string) (string, error) {
	if a.pluginSchema.RenderTerraformResources != nil {
		result, err := a.pluginSchema.RenderTerraformResources(site)
		if err != nil {
			a.logger.Debug("RenderTerraformResources error: %s", err)
		}
		return result, err
	}
	return "", nil
}

func (a *Adapter) RenderTerraformComponent(site string, component string) (*schema.ComponentSchema, error) {
	if a.pluginSchema.RenderTerraformComponent != nil {
		result, err := a.pluginSchema.RenderTerraformComponent(site, component)
		if err != nil {
			a.logger.Debug("RenderTerraformComponent error: %s", err)
		}
		return result, err
	}
	return nil, nil
}
