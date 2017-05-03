package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/azurerm"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: azurerm.Provider,
	})
}
