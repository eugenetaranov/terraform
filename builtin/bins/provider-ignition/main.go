package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/ignition"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ignition.Provider,
	})
}
