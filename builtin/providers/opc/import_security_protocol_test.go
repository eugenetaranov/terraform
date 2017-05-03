package opc

import (
	"fmt"
	"testing"

	"github.com/eugenetaranov/terraform/helper/acctest"
	"github.com/eugenetaranov/terraform/helper/resource"
)

func TestAccOPCSecurityProtocol_importBasic(t *testing.T) {
	resourceName := "opc_compute_security_protocol.test"

	ri := acctest.RandInt()
	config := fmt.Sprintf(testAccOPCSecurityProtocolBasic, ri)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityProtocolDestroy,
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
func TestAccOPCSecurityProtocol_importComplete(t *testing.T) {
	resourceName := "opc_compute_security_protocol.test"

	ri := acctest.RandInt()
	config := fmt.Sprintf(testAccOPCSecurityProtocolComplete, ri)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityProtocolDestroy,
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
