package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/vault"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vault.Provider,
	})
}
