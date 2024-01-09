package protocol

import (
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mach-composer/mach-composer-plugin-sdk/v2/schema"
)

func TestPluginRPC(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.LevelFromString("DEBUG"),
		JSONFormat: true,
	})

	adapter := NewAdapter(&schema.PluginSchema{
		Identifier: "test",
	}, logger)

	var pluginMap = map[string]plugin.Plugin{
		"MachComposerPlugin": &Plugin{Adapter: adapter, Logger: logger},
	}
	client, _ := plugin.TestPluginRPCConn(t, pluginMap, nil)
	defer client.Close()

	raw, err := client.Dispense("MachComposerPlugin")
	require.NoError(t, err)

	p, ok := raw.(schema.MachComposerPlugin)
	require.True(t, ok)

	assert.Equal(t, p.Identifier(), "test")
}
