package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/digitalocean"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: digitalocean.Provider,
	})
}
