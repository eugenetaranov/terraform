package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/cobbler"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cobbler.Provider,
	})
}
