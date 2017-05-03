package main

import (
	"github.com/eugenetaranov/terraform/builtin/provisioners/remote-exec"
	"github.com/eugenetaranov/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProvisionerFunc: remoteexec.Provisioner,
	})
}
