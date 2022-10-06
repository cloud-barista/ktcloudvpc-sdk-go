package flavors

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("flavors")
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("flavors", id)
}
