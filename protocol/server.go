package protocol

// This is the server interface for a plugin. This will be used in the plugin
// to make the functions available to the client (mach-composer)

import (
	"github.com/hashicorp/go-hclog"
)

type PluginRPCServer struct {
	adapter *Adapter
	name    string
	logger  hclog.Logger
}

func (s *PluginRPCServer) Identifier(args any, resp *string) error {
	*resp = s.adapter.Identifier()
	return nil
}

func (s *PluginRPCServer) IsEnabled(args any, resp *bool) error {
	*resp = s.adapter.IsEnabled()
	return nil
}

func (s *PluginRPCServer) Configure(args ConfigureInput, resp *ErrorOutput) error {
	err := s.adapter.Configure(args.Environment, args.Provider)
	resp.Err = wrapError(err)
	return nil
}

func (s *PluginRPCServer) GetValidationSchema(args any, resp *GetValidationSchemaOutput) error {
	result, err := s.adapter.GetValidationSchema()
	resp.Result = *result
	resp.Err = wrapError(err)
	return nil
}

func (s *PluginRPCServer) SetRemoteStateBackend(args SetRemoteStateBackendInput, resp *ErrorOutput) error {
	err := s.adapter.SetRemoteStateBackend(args.Data)
	resp.Err = wrapError(err)
	return nil
}

func (s *PluginRPCServer) SetGlobalConfig(args SetGlobalConfigInput, resp *ErrorOutput) error {
	err := s.adapter.SetGlobalConfig(args.Data)
	resp.Err = wrapError(err)
	return nil
}

func (s *PluginRPCServer) SetSiteConfig(args SetSiteConfigInput, resp *ErrorOutput) error {
	err := s.adapter.SetSiteConfig(args.Name, args.Data)
	resp.Err = wrapError(err)
	return nil
}

func (s *PluginRPCServer) SetSiteComponentConfig(args SetSiteComponentConfigInput, resp *SetSiteComponentConfigOutput) error {
	err := s.adapter.SetSiteComponentConfig(args.Site, args.Component, args.Data)
	resp.Err = wrapError(err)
	return nil
}

func (s *PluginRPCServer) SetSiteEndpointConfig(args SetSiteEndpointsConfigInput, resp *SetSiteEndpointsConfigOutput) error {
	err := s.adapter.SetSiteEndpointConfig(args.Site, args.Name, args.Data)
	resp.Err = wrapError(err)
	return nil
}

func (s *PluginRPCServer) SetComponentConfig(args SetSiteComponentConfigInput, resp *SetSiteComponentConfigOutput) error {
	err := s.adapter.SetComponentConfig(args.Component, args.Data)
	resp.Err = wrapError(err)
	return nil
}

func (s *PluginRPCServer) SetComponentEndpointsConfig(args SetComponentEndpointsConfigInput, resp *SetSiteComponentConfigOutput) error {
	err := s.adapter.SetComponentEndpointsConfig(args.Component, args.Endpoints)
	resp.Err = wrapError(err)
	return nil
}

func (s *PluginRPCServer) RenderTerraformStateBackend(
	args RenderTerraformStateBackendInput,
	resp *RenderTerraformStateBackendOutput) error {
	result, err := s.adapter.RenderTerraformStateBackend(args.Site)
	resp.Result = result
	resp.Err = wrapError(err)
	return nil
}

func (s *PluginRPCServer) RenderTerraformProviders(
	args RenderTerraformProvidersInput,
	resp *RenderTerraformProvidersOutput) error {
	result, err := s.adapter.RenderTerraformProviders(args.Site)
	resp.Result = result
	resp.Err = wrapError(err)
	return nil
}

func (s *PluginRPCServer) RenderTerraformResources(
	args RenderTerraformResourcesInput,
	resp *RenderTerraformResourcesOutput) error {
	result, err := s.adapter.RenderTerraformResources(args.Site)
	resp.Result = result
	resp.Err = wrapError(err)
	return nil
}

func (s *PluginRPCServer) RenderTerraformComponent(
	args RenderTerraformComponentInput,
	resp *RenderTerraformComponentOutput) error {
	result, err := s.adapter.RenderTerraformComponent(args.Site, args.Component)
	resp.Result = *result
	resp.Err = wrapError(err)
	return nil
}
