package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/local"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: local.Provider,
	})
}
