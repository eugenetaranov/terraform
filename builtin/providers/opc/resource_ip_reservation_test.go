package opc

import (
	"fmt"
	"testing"

	"github.com/eugenetaranov/terraform/helper/acctest"
	"github.com/eugenetaranov/terraform/helper/resource"
	"github.com/eugenetaranov/terraform/terraform"
	"github.com/hashicorp/go-oracle-terraform/compute"
)

func TestAccOPCIPReservation_Basic(t *testing.T) {

	ri := acctest.RandInt()
	config := fmt.Sprintf(testAccIPReservationBasic, ri)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPReservationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  testAccCheckIPReservationExists,
			},
		},
	})
}

func testAccCheckIPReservationExists(s *terraform.State) error {
	client := testAccProvider.Meta().(*compute.Client).IPReservations()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opc_compute_ip_reservation" {
			continue
		}

		input := compute.GetIPReservationInput{
			Name: rs.Primary.Attributes["name"],
		}
		if _, err := client.GetIPReservation(&input); err != nil {
			return fmt.Errorf("Error retrieving state of IP Reservation %s: %s", input.Name, err)
		}
	}

	return nil
}

func testAccCheckIPReservationDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*compute.Client).IPReservations()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opc_compute_ip_reservation" {
			continue
		}

		input := compute.GetIPReservationInput{
			Name: rs.Primary.Attributes["name"],
		}
		if info, err := client.GetIPReservation(&input); err == nil {
			return fmt.Errorf("IP Reservation %s still exists: %#v", input.Name, info)
		}
	}

	return nil
}

var testAccIPReservationBasic = `
resource "opc_compute_ip_reservation" "test" {
  name        = "acc-test-ip-reservation-%d"
  parent_pool = "/oracle/public/ippool"
  permanent   = true
}
`
