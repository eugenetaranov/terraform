package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/chef"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: chef.Provider,
	})
}
