package messages

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("messages")
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("messages", id)
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return getURL(c, id)
}
