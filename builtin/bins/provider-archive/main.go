package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/archive"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: archive.Provider,
	})
}
