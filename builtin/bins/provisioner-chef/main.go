package main

import (
	"github.com/eugenetaranov/terraform/builtin/provisioners/chef"
	"github.com/eugenetaranov/terraform/plugin"
	"github.com/eugenetaranov/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProvisionerFunc: func() terraform.ResourceProvisioner {
			return new(chef.ResourceProvisioner)
		},
	})
}
