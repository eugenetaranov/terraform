package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/dyn"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dyn.Provider,
	})
}
