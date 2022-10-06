package schedulerstats

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func storagePoolsListURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}
