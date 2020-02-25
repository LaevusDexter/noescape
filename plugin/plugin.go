// This file has automatically been generated on Wed Feb 26 02:10:10 +05 2020.
// DO NOT EDIT.
package plugin

import (
	"plugin"
	_ "unsafe"
)

//go:linkname Open plugin.Open
//go:noescape
func Open(path string) (*plugin.Plugin, error)

//go:linkname pluginlookup plugin.sub_pluginlookup
func pluginlookup(p *plugin.Plugin, symName string) (plugin.Symbol, error) {
	return p.Lookup(symName)
}

//go:linkname PluginLookup plugin.sub_pluginlookup
//go:noescape
func PluginLookup(p *plugin.Plugin, symName string) (plugin.Symbol, error)
