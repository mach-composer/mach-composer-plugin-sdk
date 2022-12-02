package plugin

import (
	"encoding/gob"

	"github.com/hashicorp/go-plugin"

	"github.com/mach-composer/mach-composer-plugin-sdk/helpers"
	"github.com/mach-composer/mach-composer-plugin-sdk/protocol"
	"github.com/mach-composer/mach-composer-plugin-sdk/schema"
)

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})
}

func ServePlugin(p schema.MachComposerPlugin) {
	logger := helpers.NewLogger(nil)
	if val, ok := p.(*protocol.Adapter); ok {
		val.SetLogger(logger)
	}

	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"MachComposerPlugin": &protocol.Plugin{Impl: p},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: protocol.HandShakeConfig(),
		Plugins:         pluginMap,
		Logger:          logger,
	})
}
