package env

import (
	"github.com/dan-compton/exile/pkg/plugins"
)

// PluginNamespace is the function namespace used for the plugin call function.
// It is the same as the filename by convention.
const PluginNamespace = "env"

// PluginMappers is used to register functions in template.FuncMap.
var PluginMappers = plugins.NewMappers(PluginNamespace)

func init() {
	PluginMappers.Register(new(GetCaller), new(SetCaller))
}
