package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/packet"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: packet.Provider,
	})
}
