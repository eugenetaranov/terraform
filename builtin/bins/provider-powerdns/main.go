package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/powerdns"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: powerdns.Provider,
	})
}
