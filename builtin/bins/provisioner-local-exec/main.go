package main

import (
	"github.com/eugenetaranov/terraform/builtin/provisioners/local-exec"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProvisionerFunc: localexec.Provisioner,
	})
}
