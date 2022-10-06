package volumetransfers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func transferURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("os-volume-transfer")
}

func acceptURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("os-volume-transfer", id, "accept")
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("os-volume-transfer", id)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("os-volume-transfer", "detail")
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("os-volume-transfer", id)
}
