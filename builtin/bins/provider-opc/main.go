package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/opc"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: opc.Provider,
	})
}
