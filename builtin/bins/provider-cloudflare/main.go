package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/cloudflare"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudflare.Provider,
	})
}
