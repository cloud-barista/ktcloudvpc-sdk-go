package policies

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "qos/policies"

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func createURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}
