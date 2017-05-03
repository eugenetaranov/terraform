package aws

import (
	"testing"

	"github.com/eugenetaranov/terraform/helper/resource"
)

func TestAccAWSVpnGateway_importBasic(t *testing.T) {
	resourceName := "aws_vpn_gateway.foo"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnGatewayDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVpnGatewayConfig,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
