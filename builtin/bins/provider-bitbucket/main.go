package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/bitbucket"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bitbucket.Provider,
	})
}
