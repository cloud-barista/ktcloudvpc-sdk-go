package serviceassets

import "github.com/cloud-barista/ktcloudvpc-sdk-go"

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}
