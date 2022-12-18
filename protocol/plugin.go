package protocol

import (
	"net/rpc"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type Plugin struct {
	Adapter    *Adapter
	Logger     hclog.Logger
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
	result := &PluginRPC{
		client: c,
		name:   p.Identifier,
		logger: p.Logger,
	}
	return result, nil
}
