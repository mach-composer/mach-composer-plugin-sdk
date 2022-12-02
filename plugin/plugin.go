package plugin

import (
	"github.com/mach-composer/mach-composer-plugin-sdk/helpers"
	"github.com/mach-composer/mach-composer-plugin-sdk/protocol"
	"github.com/mach-composer/mach-composer-plugin-sdk/schema"
)

func NewPlugin(fn *schema.PluginSchema) schema.MachComposerPlugin {
	logger := helpers.NewLogger(nil)
	return protocol.NewAdapter(fn, logger)
}
