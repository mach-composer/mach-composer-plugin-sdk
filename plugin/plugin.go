package plugin

import (
	"github.com/hashicorp/go-hclog"
	"github.com/mach-composer/mach-composer-plugin-sdk/v1/protocol"

	"github.com/mach-composer/mach-composer-plugin-sdk/v1/schema"
)

// NewPlugin creates a new plugin instance
func NewPlugin(s *schema.PluginSchema) schema.MachComposerPlugin {
	logger := hclog.New(&hclog.LoggerOptions{
		Output: hclog.DefaultOutput,
		Level:  hclog.Trace,
		Name:   "plugin",
	})
	return protocol.NewAdapter(s, logger)
}
