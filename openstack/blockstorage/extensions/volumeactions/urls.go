package volumeactions

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func actionURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}
