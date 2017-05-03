package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/grafana"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: grafana.Provider,
	})
}
