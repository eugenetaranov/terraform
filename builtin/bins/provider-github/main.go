package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/github"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: github.Provider,
	})
}
