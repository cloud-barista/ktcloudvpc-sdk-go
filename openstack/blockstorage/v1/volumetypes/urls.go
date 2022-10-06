package volumetypes

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("types")
}

func createURL(c *ktvpcsdk.ServiceClient) string {
	return listURL(c)
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return getURL(c, id)
}
