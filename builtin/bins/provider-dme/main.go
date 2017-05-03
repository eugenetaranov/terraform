package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/dme"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dme.Provider,
	})
}
