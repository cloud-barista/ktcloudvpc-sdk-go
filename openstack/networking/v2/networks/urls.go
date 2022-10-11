package networks

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("VPC", id)  // Modified by B.T. Oh
}

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("VPC")  // Modified by B.T. Oh
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
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
