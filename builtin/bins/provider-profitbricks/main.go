package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/profitbricks"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: profitbricks.Provider,
	})
}
