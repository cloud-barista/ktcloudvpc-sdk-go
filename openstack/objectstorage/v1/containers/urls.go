package containers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.Endpoint
}

func createURL(c *ktvpcsdk.ServiceClient, container string) string {
	return c.ServiceURL(container)
}

func getURL(c *ktvpcsdk.ServiceClient, container string) string {
	return createURL(c, container)
}

func deleteURL(c *ktvpcsdk.ServiceClient, container string) string {
	return createURL(c, container)
}

func updateURL(c *ktvpcsdk.ServiceClient, container string) string {
	return createURL(c, container)
}

func bulkDeleteURL(c *ktvpcsdk.ServiceClient) string {
	return c.Endpoint + "?bulk-delete=true"
}
