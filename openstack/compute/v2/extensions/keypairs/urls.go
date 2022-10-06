package keypairs

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "os-keypairs"

func resourceURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return resourceURL(c)
}

func createURL(c *ktvpcsdk.ServiceClient) string {
	return resourceURL(c)
}

func getURL(c *ktvpcsdk.ServiceClient, name string) string {
	return c.ServiceURL(resourcePath, name)
}

func deleteURL(c *ktvpcsdk.ServiceClient, name string) string {
	return getURL(c, name)
}
