package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/rancher"
	"github.com/eugenetaranov/terraform/plugin"
	"github.com/eugenetaranov/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return rancher.Provider()
		},
	})
}
