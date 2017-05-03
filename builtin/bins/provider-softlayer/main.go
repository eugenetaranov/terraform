package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/softlayer"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: softlayer.Provider,
	})
}
