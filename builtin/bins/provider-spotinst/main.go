package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/spotinst"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: spotinst.Provider,
	})
}
