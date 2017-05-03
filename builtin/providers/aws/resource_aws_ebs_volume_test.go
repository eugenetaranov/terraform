package aws

import (
	"fmt"
	"testing"

	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/eugenetaranov/terraform/helper/acctest"
	"github.com/eugenetaranov/terraform/helper/resource"
	"github.com/eugenetaranov/terraform/terraform"
)

func TestAccAWSEBSVolume_basic(t *testing.T) {
	var v ec2.Volume
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: "aws_ebs_volume.test",
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsEbsVolumeConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVolumeExists("aws_ebs_volume.test", &v),
				),
			},
		},
	})
}

func TestAccAWSEBSVolume_updateSize(t *testing.T) {
	var v ec2.Volume
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: "aws_ebs_volume.test",
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsEbsVolumeConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVolumeExists("aws_ebs_volume.test", &v),
					resource.TestCheckResourceAttr("aws_ebs_volume.test", "size", "1"),
				),
			},
			{
				Config: testAccAwsEbsVolumeConfigUpdateSize,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVolumeExists("aws_ebs_volume.test", &v),
					resource.TestCheckResourceAttr("aws_ebs_volume.test", "size", "10"),
				),
			},
		},
	})
}

func TestAccAWSEBSVolume_updateType(t *testing.T) {
	var v ec2.Volume
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: "aws_ebs_volume.test",
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsEbsVolumeConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVolumeExists("aws_ebs_volume.test", &v),
					resource.TestCheckResourceAttr("aws_ebs_volume.test", "type", "gp2"),
				),
			},
			{
				Config: testAccAwsEbsVolumeConfigUpdateType,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVolumeExists("aws_ebs_volume.test", &v),
					resource.TestCheckResourceAttr("aws_ebs_volume.test", "type", "sc1"),
				),
			},
		},
	})
}

func TestAccAWSEBSVolume_updateIops(t *testing.T) {
	var v ec2.Volume
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: "aws_ebs_volume.test",
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsEbsVolumeConfigWithIops,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVolumeExists("aws_ebs_volume.test", &v),
					resource.TestCheckResourceAttr("aws_ebs_volume.test", "iops", "100"),
				),
			},
			{
				Config: testAccAwsEbsVolumeConfigWithIopsUpdated,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVolumeExists("aws_ebs_volume.test", &v),
					resource.TestCheckResourceAttr("aws_ebs_volume.test", "iops", "200"),
				),
			},
		},
	})
}

func TestAccAWSEBSVolume_kmsKey(t *testing.T) {
	var v ec2.Volume
	ri := acctest.RandInt()
	config := fmt.Sprintf(testAccAwsEbsVolumeConfigWithKmsKey, ri)
	keyRegex := regexp.MustCompile("^arn:aws:([a-zA-Z0-9\\-])+:([a-z]{2}-[a-z]+-\\d{1})?:(\\d{12})?:(.*)$")

	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: "aws_ebs_volume.test",
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVolumeExists("aws_ebs_volume.test", &v),
					resource.TestCheckResourceAttr("aws_ebs_volume.test", "encrypted", "true"),
					resource.TestMatchResourceAttr("aws_ebs_volume.test", "kms_key_id", keyRegex),
				),
			},
		},
	})
}

func TestAccAWSEBSVolume_NoIops(t *testing.T) {
	var v ec2.Volume
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsEbsVolumeConfigWithNoIops,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVolumeExists("aws_ebs_volume.iops_test", &v),
				),
			},
		},
	})
}

func TestAccAWSEBSVolume_withTags(t *testing.T) {
	var v ec2.Volume
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: "aws_ebs_volume.tags_test",
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsEbsVolumeConfigWithTags,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVolumeExists("aws_ebs_volume.tags_test", &v),
				),
			},
		},
	})
}

func testAccCheckVolumeExists(n string, v *ec2.Volume) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		conn := testAccProvider.Meta().(*AWSClient).ec2conn

		request := &ec2.DescribeVolumesInput{
			VolumeIds: []*string{aws.String(rs.Primary.ID)},
		}

		response, err := conn.DescribeVolumes(request)
		if err == nil {
			if response.Volumes != nil && len(response.Volumes) > 0 {
				*v = *response.Volumes[0]
				return nil
			}
		}
		return fmt.Errorf("Error finding EC2 volume %s", rs.Primary.ID)
	}
}

const testAccAwsEbsVolumeConfig = `
resource "aws_ebs_volume" "test" {
  availability_zone = "us-west-2a"
  type = "gp2"
  size = 1
  tags {
    Name = "tf-acc-test-ebs-volume-test"
  }
}
`

const testAccAwsEbsVolumeConfigUpdateSize = `
resource "aws_ebs_volume" "test" {
  availability_zone = "us-west-2a"
  type = "gp2"
  size = 10
  tags {
    Name = "tf-acc-test-ebs-volume-test"
  }
}
`

const testAccAwsEbsVolumeConfigUpdateType = `
resource "aws_ebs_volume" "test" {
  availability_zone = "us-west-2a"
  type = "sc1"
  size = 500
  tags {
    Name = "tf-acc-test-ebs-volume-test"
  }
}
`

const testAccAwsEbsVolumeConfigWithIops = `
resource "aws_ebs_volume" "test" {
  availability_zone = "us-west-2a"
  type = "io1"
  size = 4
  iops = 100
  tags {
    Name = "tf-acc-test-ebs-volume-test"
  }
}
`

const testAccAwsEbsVolumeConfigWithIopsUpdated = `
resource "aws_ebs_volume" "test" {
  availability_zone = "us-west-2a"
  type = "io1"
  size = 4
  iops = 200
  tags {
    Name = "tf-acc-test-ebs-volume-test"
  }
}
`

const testAccAwsEbsVolumeConfigWithKmsKey = `
resource "aws_kms_key" "foo" {
  description = "Terraform acc test %d"
  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "kms-tf-1",
  "Statement": [
    {
      "Sid": "Enable IAM User Permissions",
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": "kms:*",
      "Resource": "*"
    }
  ]
}
POLICY
}

resource "aws_ebs_volume" "test" {
  availability_zone = "us-west-2a"
  size = 1
  encrypted = true
  kms_key_id = "${aws_kms_key.foo.arn}"
}
`

const testAccAwsEbsVolumeConfigWithTags = `
resource "aws_ebs_volume" "tags_test" {
  availability_zone = "us-west-2a"
  size = 1
  tags {
    Name = "TerraformTest"
  }
}
`

const testAccAwsEbsVolumeConfigWithNoIops = `
resource "aws_ebs_volume" "iops_test" {
  availability_zone = "us-west-2a"
  size = 10
  type = "gp2"
  iops = 0
  tags {
    Name = "TerraformTest"
  }
}
`
