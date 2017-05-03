package main

import (
	"github.com/eugenetaranov/terraform/builtin/provisioners/file"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProvisionerFunc: file.Provisioner,
	})
}
