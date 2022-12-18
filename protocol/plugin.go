package protocol

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"

	"github.com/mach-composer/mach-composer-plugin-sdk/schema"
)

type Plugin struct {
	Impl       schema.MachComposerPlugin
	Identifier string
}

func (p *Plugin) Server(*plugin.MuxBroker) (any, error) {
	result := &PluginRPCServer{
		adapter: p.Adapter,
		name:    p.Identifier,
		logger:  p.Logger,
	}
	return result, nil
}

func (p *Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (any, error) {
	return &PluginRPC{client: c, identifier: p.Identifier}, nil
}
