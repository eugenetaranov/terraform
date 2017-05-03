package aws

import (
	"fmt"
	"testing"

	"github.com/eugenetaranov/terraform/helper/acctest"
	"github.com/eugenetaranov/terraform/helper/resource"
)

func TestAWSElasticBeanstalkApplication_importBasic(t *testing.T) {
	resourceName := "aws_elastic_beanstalk_application.tftest"
	config := fmt.Sprintf("tf-test-name-%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBeanstalkAppDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBeanstalkAppImportConfig(config),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBeanstalkAppImportConfig(name string) string {
	return fmt.Sprintf(`resource "aws_elastic_beanstalk_application" "tftest" {
	  name = "%s"
	  description = "tf-test-desc"
	}`, name)
}
