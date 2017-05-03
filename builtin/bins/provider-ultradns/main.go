package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/ultradns"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ultradns.Provider,
	})
}
