package aws

import (
	"testing"

	"github.com/eugenetaranov/terraform/helper/acctest"
	"github.com/eugenetaranov/terraform/helper/resource"
)

func TestAccAWSIAMAccountAlias_importBasic(t *testing.T) {
	resourceName := "aws_iam_account_alias.test"

	rstring := acctest.RandString(5)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSIAMAccountAliasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSIAMAccountAliasConfig(rstring),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
