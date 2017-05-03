package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/statuscake"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: statuscake.Provider,
	})
}
