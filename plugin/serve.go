package plugin

import (
	"encoding/gob"
	"log"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/mach-composer/mach-composer-plugin-sdk/protocol"
	"github.com/mach-composer/mach-composer-plugin-sdk/schema"
)

func init() {
	gob.Register(map[string]any{})
	gob.Register([]any{})
	gob.Register(protocol.PluginError{})
}

func ServePlugin(a schema.MachComposerPlugin) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       a.Identifier(),
		Level:      hclog.LevelFromString("DEBUG"),
		JSONFormat: true,
	})
	log.SetOutput(logger.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true}))
	hclog.SetDefault(logger)

	adapter, ok := a.(*protocol.Adapter)
	if !ok {
		panic("plugin should use protocol.Adapter")
	}

	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"MachComposerPlugin": &protocol.Plugin{Adapter: adapter, Logger: logger},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: protocol.HandShakeConfig(),
		Plugins:         pluginMap,
		Logger:          logger,
	})
}
