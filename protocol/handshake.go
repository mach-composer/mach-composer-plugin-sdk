package protocol

import "github.com/hashicorp/go-plugin"

// HandShakeConfig is used to just do a basic handshake between a plugin and host. If the handshake fails, a user-friendly
// error is shown. This prevents users from executing bad plugins or executing a plugin directory. It is a UX feature,
// not a security feature.
func HandShakeConfig() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion: 2,

		//Don't change these as it will mean the plugin cannot be loaded by mach-composer-cli!
		MagicCookieKey:   "MACH_COMPOSER",
		MagicCookieValue: "plugin",
	}
}
