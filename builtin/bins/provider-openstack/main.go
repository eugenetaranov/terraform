package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/openstack"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: openstack.Provider,
	})
}
