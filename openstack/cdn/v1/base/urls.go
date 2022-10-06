package base

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func getURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL()
}

func pingURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("ping")
}
