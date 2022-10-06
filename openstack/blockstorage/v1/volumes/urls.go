package volumes

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("volumes")
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return createURL(c)
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id)
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}
