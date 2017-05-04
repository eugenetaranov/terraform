package bootfromvolume

import "github.com/eugenetaranov/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-volumes_boot")
}
