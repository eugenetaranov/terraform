package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/ns1"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ns1.Provider,
	})
}
