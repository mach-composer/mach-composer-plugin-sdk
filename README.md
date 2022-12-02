# mach-composer-plugin-sdk
SDK to create plugins for Mach Composer


## Writing plugins
Mach Composer utiltizes the [go-plugin](https://github.com/hashicorp/go-plugin)
package for implementing plugins.  This means that mach-composer runs each
plugin in a separate process whereby communication between mach-composer and the
plugin happens via `net/rpc`.

Mach Composer automatically finds plugins defined in the config file by looking
for an executable named `mach-composer-plugin-<name>`.

A plugin exists out of three major parts:

 1. Setup functionality
 2. Processing settings (global, site or component)
 3. Rendering of terraform snippets


## Example

Below is a very minimal implementation of a plugin.

```go

import (
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"
	"github.com/mach-composer/mach-composer-plugin-sdk/schema"
)


func MyPlugin() schema.MachComposerPlugin {
	return plugin.NewPlugin(&schema.PluginSchema{
		Identifier: "<name of plugin>",
	})
}

func main() {
	p := MyPlugin()
	plugin.ServePlugin(p)
}
```
