package azurerm

import (
	"fmt"
	"testing"

	"github.com/eugenetaranov/terraform/helper/acctest"
	"github.com/eugenetaranov/terraform/helper/resource"
)

func TestAccAzureRMDnsZone_importBasic(t *testing.T) {
	resourceName := "azurerm_dns_zone.test"

	ri := acctest.RandInt()
	config := fmt.Sprintf(testAccAzureRMDnsZone_basic, ri, ri)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMDnsZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: config,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
