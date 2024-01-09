package schema

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/xeipuuv/gojsonschema"
)

type ComponentSchema struct {
	Resources string
	Variables string
	DependsOn []string
	Providers []string
}

type ValidationSchema struct {
	GlobalConfigSchema             map[string]any
	RemoteStateSchema              map[string]any
	SiteConfigSchema               map[string]any
	SiteComponentConfigSchema      map[string]any
	SiteEndpointConfig             map[string]any
	ComponentConfigSchema          map[string]any
	ComponentEndpointsConfigSchema map[string]any
}

func (v *ValidationSchema) Validate(schema map[string]any, data map[string]any) error {
	sl := gojsonschema.NewRawLoader(schema)
	dl := gojsonschema.NewRawLoader(data)

	result, err := gojsonschema.Validate(sl, dl)
	if err != nil {
		return err
	}
	if !result.Valid() {
		details := strings.Builder{}
		for _, e := range result.Errors() {
			details.WriteString(fmt.Sprintf(" - %s\n", e))
		}
		hclog.Default().Debug("plugin received invalid data", "details", details.String())
		return fmt.Errorf("plugin received invalid data")
	}
	return nil
}
