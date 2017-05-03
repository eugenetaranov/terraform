package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/rundeck"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: rundeck.Provider,
	})
}
