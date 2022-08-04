package base

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL()
}

func pingURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("ping")
}
