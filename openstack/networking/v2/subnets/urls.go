package subnets

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("Network", id)  // Modified by B.T. Oh
}

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("Network")  // Modified by B.T. Oh
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
