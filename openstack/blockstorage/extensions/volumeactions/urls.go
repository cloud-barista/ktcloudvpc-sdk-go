package volumeactions

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func actionURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}
