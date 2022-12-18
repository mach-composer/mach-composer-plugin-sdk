package schema

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/xeipuuv/gojsonschema"
)

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
	SiteEndpointConfig             map[string]any
	ComponentConfigSchema          map[string]any
	ComponentEndpointsConfigSchema map[string]any
}

func (v *ValidationSchema) Validate(schema map[string]any, data map[string]any) error {
	sl := gojsonschema.NewRawLoader(schema)
	dl := gojsonschema.NewRawLoader(data)

	result, err := gojsonschema.Validate(sl, dl)
	if err != nil {
		return err
	}
	if !result.Valid() {
		details := strings.Builder{}
		for _, e := range result.Errors() {
			details.WriteString(fmt.Sprintf(" - %s\n", e))
		}
		hclog.Default().Debug("plugin received invalid data", "details", details.String())
		return fmt.Errorf("plugin received invalid data")
	}
	return nil
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
	SetSiteEndpointConfig       func(site string, name string, data map[string]any) error
	SetComponentConfig          func(component string, data map[string]any) error
	SetComponentEndpointsConfig func(component string, endpoints map[string]string) error

	RenderTerraformStateBackend func(site string) (string, error)
	RenderTerraformProviders    func(site string) (string, error)
	RenderTerraformResources    func(site string) (string, error)
	RenderTerraformComponent    func(site string, component string) (*ComponentSchema, error)
}
