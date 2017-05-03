package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/terraform"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: terraform.Provider,
	})
}
