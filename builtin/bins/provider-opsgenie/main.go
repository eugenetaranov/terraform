package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/opsgenie"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: opsgenie.Provider,
	})
}
