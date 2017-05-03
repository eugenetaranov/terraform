package azurerm

import (
	"fmt"
	"testing"

	"github.com/eugenetaranov/terraform/helper/acctest"
	"github.com/eugenetaranov/terraform/helper/resource"
)

func TestAccAzureRMContainerRegistry_importBasic(t *testing.T) {
	resourceName := "azurerm_container_registry.test"

	ri := acctest.RandInt()
	rs := acctest.RandString(4)
	config := fmt.Sprintf(testAccAzureRMContainerRegistry_basic, ri, rs, ri)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMContainerRegistryDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
			},

			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"storage_account"},
			},
		},
	})
}

func TestAccAzureRMContainerRegistry_importComplete(t *testing.T) {
	resourceName := "azurerm_container_registry.test"

	ri := acctest.RandInt()
	rs := acctest.RandString(4)
	config := fmt.Sprintf(testAccAzureRMContainerRegistry_complete, ri, rs, ri)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMContainerRegistryDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
			},

			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"storage_account"},
			},
		},
	})
}
