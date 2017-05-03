package opc

import (
	"testing"

	"github.com/eugenetaranov/terraform/helper/acctest"
	"github.com/eugenetaranov/terraform/helper/resource"
)

func TestAccOPCIPNetwork_importBasic(t *testing.T) {
	resourceName := "opc_compute_ip_network.test"

	rInt := acctest.RandInt()
	config := testAccOPCIPNetworkConfig_Basic(rInt)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: opcResourceCheck(resourceName, testAccOPCCheckIPNetworkDestroyed),
		Steps: []resource.TestStep{
			{
				Config: config,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
