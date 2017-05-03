package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/tls"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: tls.Provider,
	})
}
