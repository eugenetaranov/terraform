package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/datadog"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: datadog.Provider,
	})
}
