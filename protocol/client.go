package protocol

// This is the client interface to a pluginSchema. This will be used in mach-composer
// to call the functions provided by the plugins (servers)

import (
	"net/rpc"

	"github.com/hashicorp/go-hclog"

	"github.com/mach-composer/mach-composer-plugin-sdk/v1/schema"
)

type PluginRPC struct {
	client *rpc.Client
	name   string
	logger hclog.Logger
}

var _ schema.MachComposerPlugin = (*PluginRPC)(nil)

func (p *PluginRPC) Identifier() string {
	// return p.identifier
	var resp string
	err := p.client.Call("Plugin.Identifier", new(any), &resp)
	if err != nil {
		p.logger.Error(err.Error(), "error", err)
		return ""
	}
	return resp
}

func (p *PluginRPC) Configure(environment string, provider string) error {
	param := ConfigureInput{
		Environment: environment,
		Provider:    provider,
	}
	resp := ErrorOutput{}
	err := p.client.Call("Plugin.Configure", param, &resp)
	if err != nil {
		return err
	}
	return unwrapError(p.name, resp.Err)
}

func (p *PluginRPC) GetValidationSchema() (*schema.ValidationSchema, error) {
	resp := GetValidationSchemaOutput{}
	err := p.client.Call("Plugin.GetValidationSchema", new(any), &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

func (p *PluginRPC) SetGlobalConfig(meta schema.Meta, data map[string]any) error {
	param := SetGlobalConfigInput{
		Data: makeNil(data),
	}
	resp := ErrorOutput{}
	err := p.client.Call("Plugin.SetGlobalConfig", param, &resp)
	if err != nil {
		return err
	}
	return unwrapError(p.name, resp.Err)
}

func (p *PluginRPC) SetSiteConfig(name string, meta schema.Meta, data map[string]any) error {
	param := SetSiteConfigInput{
		Name: name,
		Data: makeNil(data),
	}
	resp := ErrorOutput{}
	err := p.client.Call("Plugin.SetSiteConfig", param, &resp)
	if err != nil {
		return err
	}
	return unwrapError(p.name, resp.Err)
}

func (p *PluginRPC) SetSiteComponentConfig(site string, component string, meta schema.Meta, data map[string]any) error {
	param := SetSiteComponentConfigInput{
		Site:      site,
		Component: component,
		Data:      makeNil(data),
	}
	resp := SetSiteComponentConfigOutput{}
	err := p.client.Call("Plugin.SetSiteComponentConfig", param, &resp)
	if err != nil {
		return err
	}
	return unwrapError(p.name, resp.Err)
}

func (p *PluginRPC) SetSiteEndpointConfig(site string, name string, meta schema.Meta, data map[string]any) error {
	param := SetSiteEndpointsConfigInput{
		Site: site,
		Name: name,
		Data: makeNil(data),
	}
	resp := SetSiteEndpointsConfigOutput{}
	err := p.client.Call("Plugin.SetSiteEndpointConfig", param, &resp)
	if err != nil {
		return err
	}
	return unwrapError(p.name, resp.Err)
}

func (p *PluginRPC) SetComponentConfig(component string, version string, meta schema.Meta, data map[string]any) error {
	param := SetComponentConfigInput{
		Component: component,
		Data:      makeNil(data),
	}
	resp := SetComponentConfigOutput{}
	err := p.client.Call("Plugin.SetComponentConfig", param, &resp)
	if err != nil {
		return err
	}
	return unwrapError(p.name, resp.Err)
}

func (p *PluginRPC) SetComponentEndpointsConfig(component string, meta schema.Meta, endpoints map[string]string) error {
	param := SetComponentEndpointsConfigInput{
		Component: component,
		Endpoints: endpoints,
	}
	resp := SetComponentEndpointsConfigOutput{}
	err := p.client.Call("Plugin.SetComponentEndpointsConfig", param, &resp)
	if err != nil {
		return err
	}
	return unwrapError(p.name, resp.Err)
}

func (p *PluginRPC) RenderTerraformProviders(site string) (string, error) {
	param := RenderTerraformProvidersInput{
		Site: site,
	}
	resp := RenderTerraformProvidersOutput{}
	err := p.client.Call("Plugin.RenderTerraformProviders", param, &resp)
	if err != nil {
		return "", err
	}
	return resp.Result, unwrapError(p.name, resp.Err)
}

func (p *PluginRPC) RenderTerraformResources(site string) (string, error) {
	param := RenderTerraformResourcesInput{
		Site: site,
	}
	resp := RenderTerraformResourcesOutput{}
	err := p.client.Call("Plugin.RenderTerraformResources", param, &resp)
	if err != nil {
		return "", err
	}
	return resp.Result, unwrapError(p.name, resp.Err)
}

func (p *PluginRPC) RenderTerraformComponent(site string, component string) (*schema.ComponentSchema, error) {
	param := RenderTerraformComponentInput{
		Site:      site,
		Component: component,
	}
	resp := RenderTerraformComponentOutput{}
	err := p.client.Call("Plugin.RenderTerraformComponent", param, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Result, unwrapError(p.name, resp.Err)
}

func makeNil[T any](data map[string]T) map[string]T {
	if len(data) == 0 {
		return nil
	}
	return data
}
