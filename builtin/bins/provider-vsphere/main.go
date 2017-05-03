package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/vsphere"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vsphere.Provider,
	})
}
