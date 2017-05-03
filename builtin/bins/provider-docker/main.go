package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/docker"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: docker.Provider,
	})
}
