package tenants

import "github.com/eugenetaranov/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("tenants")
}
