package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/template"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: template.Provider,
	})
}
