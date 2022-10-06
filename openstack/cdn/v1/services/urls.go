package services

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("services")
}

func createURL(c *ktvpcsdk.ServiceClient) string {
	return listURL(c)
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("services", id)
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return getURL(c, id)
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return getURL(c, id)
}
