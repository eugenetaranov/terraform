package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/google"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: google.Provider,
	})
}
