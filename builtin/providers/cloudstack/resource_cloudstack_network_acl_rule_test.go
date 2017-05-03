package cloudstack

import (
	"fmt"
	"strings"
	"testing"

	"github.com/eugenetaranov/terraform/helper/resource"
	"github.com/eugenetaranov/terraform/terraform"
	"github.com/xanzy/go-cloudstack/cloudstack"
)

func TestAccCloudStackNetworkACLRule_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudStackNetworkACLRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCloudStackNetworkACLRule_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudStackNetworkACLRulesExist("cloudstack_network_acl.foo"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.#", "3"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.action", "allow"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.cidr_list.2835005819", "172.16.100.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.ports.#", "2"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.ports.1889509032", "80"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.ports.3638101695", "443"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.traffic_type", "ingress"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1480917538.action", "allow"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1480917538.cidr_list.#", "1"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1480917538.cidr_list.3056857544", "172.18.100.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1480917538.icmp_code", "-1"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1480917538.icmp_type", "-1"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1480917538.traffic_type", "ingress"),
				),
			},
		},
	})
}

func TestAccCloudStackNetworkACLRule_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudStackNetworkACLRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCloudStackNetworkACLRule_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudStackNetworkACLRulesExist("cloudstack_network_acl.foo"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.#", "3"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.action", "allow"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.cidr_list.2835005819", "172.16.100.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.ports.#", "2"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.ports.1889509032", "80"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.ports.3638101695", "443"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2898748868.traffic_type", "ingress"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1480917538.action", "allow"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1480917538.cidr_list.#", "1"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1480917538.cidr_list.3056857544", "172.18.100.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1480917538.icmp_code", "-1"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1480917538.icmp_type", "-1"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1480917538.traffic_type", "ingress"),
				),
			},

			resource.TestStep{
				Config: testAccCloudStackNetworkACLRule_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudStackNetworkACLRulesExist("cloudstack_network_acl.foo"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.#", "4"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1724235854.action", "deny"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1724235854.cidr_list.3482919157", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1724235854.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1724235854.ports.#", "2"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1724235854.ports.1209010669", "1000-2000"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1724235854.ports.1889509032", "80"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.1724235854.traffic_type", "egress"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2090315355.action", "deny"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2090315355.cidr_list.#", "2"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2090315355.cidr_list.2104435309", "172.18.101.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2090315355.cidr_list.3056857544", "172.18.100.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2090315355.icmp_code", "-1"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2090315355.icmp_type", "-1"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2090315355.traffic_type", "ingress"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2576683033.action", "allow"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2576683033.cidr_list.3056857544", "172.18.100.0/24"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2576683033.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2576683033.ports.#", "2"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2576683033.ports.1889509032", "80"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2576683033.ports.3638101695", "443"),
					resource.TestCheckResourceAttr(
						"cloudstack_network_acl_rule.foo", "rule.2576683033.traffic_type", "ingress"),
				),
			},
		},
	})
}

func testAccCheckCloudStackNetworkACLRulesExist(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No network ACL rule ID is set")
		}

		for k, id := range rs.Primary.Attributes {
			if !strings.Contains(k, ".uuids.") || strings.HasSuffix(k, ".uuids.%") {
				continue
			}

			cs := testAccProvider.Meta().(*cloudstack.CloudStackClient)
			_, count, err := cs.NetworkACL.GetNetworkACLByID(id)

			if err != nil {
				return err
			}

			if count == 0 {
				return fmt.Errorf("Network ACL rule %s not found", k)
			}
		}

		return nil
	}
}

func testAccCheckCloudStackNetworkACLRuleDestroy(s *terraform.State) error {
	cs := testAccProvider.Meta().(*cloudstack.CloudStackClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudstack_network_acl_rule" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No network ACL rule ID is set")
		}

		for k, id := range rs.Primary.Attributes {
			if !strings.Contains(k, ".uuids.") || strings.HasSuffix(k, ".uuids.%") {
				continue
			}

			_, _, err := cs.NetworkACL.GetNetworkACLByID(id)
			if err == nil {
				return fmt.Errorf("Network ACL rule %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

var testAccCloudStackNetworkACLRule_basic = fmt.Sprintf(`
resource "cloudstack_vpc" "foobar" {
  name = "terraform-vpc"
  cidr = "%s"
  vpc_offering = "%s"
  zone = "%s"
}

resource "cloudstack_network_acl" "foo" {
  name = "terraform-acl"
  description = "terraform-acl-text"
  vpc_id = "${cloudstack_vpc.foobar.id}"
}

resource "cloudstack_network_acl_rule" "foo" {
  acl_id = "${cloudstack_network_acl.foo.id}"

  rule {
  	action = "allow"
    cidr_list = ["172.18.100.0/24"]
    protocol = "all"
    traffic_type = "ingress"
  }

  rule {
  	action = "allow"
    cidr_list = ["172.18.100.0/24"]
    protocol = "icmp"
    icmp_type = "-1"
    icmp_code = "-1"
    traffic_type = "ingress"
  }

  rule {
    cidr_list = ["172.16.100.0/24"]
    protocol = "tcp"
    ports = ["80", "443"]
    traffic_type = "ingress"
  }
}`,
	CLOUDSTACK_VPC_CIDR_1,
	CLOUDSTACK_VPC_OFFERING,
	CLOUDSTACK_ZONE)

var testAccCloudStackNetworkACLRule_update = fmt.Sprintf(`
resource "cloudstack_vpc" "foobar" {
  name = "terraform-vpc"
  cidr = "%s"
  vpc_offering = "%s"
  zone = "%s"
}

resource "cloudstack_network_acl" "foo" {
  name = "terraform-acl"
  description = "terraform-acl-text"
  vpc_id = "${cloudstack_vpc.foobar.id}"
}

resource "cloudstack_network_acl_rule" "foo" {
  acl_id = "${cloudstack_network_acl.foo.id}"

  rule {
  	action = "deny"
    cidr_list = ["172.18.100.0/24"]
    protocol = "all"
    traffic_type = "ingress"
  }

  rule {
  	action = "deny"
		cidr_list = ["172.18.100.0/24", "172.18.101.0/24"]
    protocol = "icmp"
    icmp_type = "-1"
    icmp_code = "-1"
    traffic_type = "ingress"
  }

  rule {
	  action = "allow"
    cidr_list = ["172.18.100.0/24"]
    protocol = "tcp"
    ports = ["80", "443"]
    traffic_type = "ingress"
  }

  rule {
	  action = "deny"
    cidr_list = ["10.0.0.0/24"]
    protocol = "tcp"
    ports = ["80", "1000-2000"]
    traffic_type = "egress"
  }
}`,
	CLOUDSTACK_VPC_CIDR_1,
	CLOUDSTACK_VPC_OFFERING,
	CLOUDSTACK_ZONE)
