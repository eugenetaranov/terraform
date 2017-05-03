package main

import (
	"github.com/eugenetaranov/terraform/builtin/providers/heroku"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: heroku.Provider,
	})
}
