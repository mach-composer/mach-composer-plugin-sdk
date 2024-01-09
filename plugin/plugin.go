package plugin

import (
	"github.com/hashicorp/go-hclog"
	"github.com/mach-composer/mach-composer-plugin-sdk/v2/protocol"

	"github.com/mach-composer/mach-composer-plugin-sdk/v2/schema"
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
