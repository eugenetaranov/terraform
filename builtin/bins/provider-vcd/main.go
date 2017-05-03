package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/vcd"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vcd.Provider,
	})
}
