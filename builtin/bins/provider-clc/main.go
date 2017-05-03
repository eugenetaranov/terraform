package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/clc"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: clc.Provider,
	})
}
