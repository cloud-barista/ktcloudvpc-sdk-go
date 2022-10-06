package availabilityzones

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}

func listDetailURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("os-availability-zone", "detail")
}
