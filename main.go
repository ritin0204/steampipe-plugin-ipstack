package main

import (
	"github.com/turbot/steampipe-plugin-ipstack/ipstack"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: ipstack.Plugin})
}
