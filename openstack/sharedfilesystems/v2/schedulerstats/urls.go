package schedulerstats

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func poolsListURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "pools")
}

func poolsListDetailURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "pools", "detail")
}
