package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/consul"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: consul.Provider,
	})
}
