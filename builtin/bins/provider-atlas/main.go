package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/atlas"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: atlas.Provider,
	})
}
