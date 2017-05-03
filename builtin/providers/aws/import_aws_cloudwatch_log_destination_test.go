package aws

import (
	"testing"

	"github.com/eugenetaranov/terraform/helper/acctest"
	"github.com/eugenetaranov/terraform/helper/resource"
)

func TestAccAWSCloudwatchLogDestination_importBasic(t *testing.T) {
	resourceName := "aws_cloudwatch_log_destination.test"

	rstring := acctest.RandString(5)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSCloudwatchLogDestinationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSCloudwatchLogDestinationConfig(rstring),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
