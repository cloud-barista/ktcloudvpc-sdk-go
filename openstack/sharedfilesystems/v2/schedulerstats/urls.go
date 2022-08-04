package schedulerstats

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func poolsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "pools")
}

func poolsListDetailURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "pools", "detail")
}
