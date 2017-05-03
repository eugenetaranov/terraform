package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/external"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: external.Provider,
	})
}
