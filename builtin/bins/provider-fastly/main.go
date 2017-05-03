package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/fastly"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fastly.Provider,
	})
}
