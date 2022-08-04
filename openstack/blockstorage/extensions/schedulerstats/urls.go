package schedulerstats

import "github.com/innodreamer/ktvpc-sdk_poc"

func storagePoolsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}
