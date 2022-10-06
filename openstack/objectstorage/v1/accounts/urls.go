package accounts

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func getURL(c *ktvpcsdk.ServiceClient) string {
	return c.Endpoint
}

func updateURL(c *ktvpcsdk.ServiceClient) string {
	return getURL(c)
}
