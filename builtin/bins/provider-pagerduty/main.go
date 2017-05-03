package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/pagerduty"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pagerduty.Provider,
	})
}
