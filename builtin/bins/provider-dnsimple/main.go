package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/dnsimple"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dnsimple.Provider,
	})
}
