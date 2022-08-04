package serviceassets

import "github.com/innodreamer/ktvpc-sdk_poc"

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}
