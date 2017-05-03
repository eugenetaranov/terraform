package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/dns"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dns.Provider,
	})
}
