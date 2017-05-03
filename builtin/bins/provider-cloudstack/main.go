package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/cloudstack"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudstack.Provider,
	})
}
