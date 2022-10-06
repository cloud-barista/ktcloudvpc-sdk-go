package objects

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

func listURL(c *ktvpcsdk.ServiceClient, container string) string {
	return c.ServiceURL(container)
}

func copyURL(c *ktvpcsdk.ServiceClient, container, object string) string {
	return c.ServiceURL(container, object)
}

func createURL(c *ktvpcsdk.ServiceClient, container, object string) string {
	return copyURL(c, container, object)
}

func getURL(c *ktvpcsdk.ServiceClient, container, object string) string {
	return copyURL(c, container, object)
}

func deleteURL(c *ktvpcsdk.ServiceClient, container, object string) string {
	return copyURL(c, container, object)
}

func downloadURL(c *ktvpcsdk.ServiceClient, container, object string) string {
	return copyURL(c, container, object)
}

func updateURL(c *ktvpcsdk.ServiceClient, container, object string) string {
	return copyURL(c, container, object)
}

func bulkDeleteURL(c *ktvpcsdk.ServiceClient) string {
	return c.Endpoint + "?bulk-delete=true"
}
