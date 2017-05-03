package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/logentries"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: logentries.Provider,
	})
}
