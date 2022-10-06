package buildinfo

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func getURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("build_info")
}
