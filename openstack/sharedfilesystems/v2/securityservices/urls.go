package securityservices

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("security-services")
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("security-services", id)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("security-services", "detail")
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}
