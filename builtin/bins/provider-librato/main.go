package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/librato"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: librato.Provider,
	})
}
