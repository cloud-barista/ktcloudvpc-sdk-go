package tenantnetworks

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "os-tenant-networks"

func resourceURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return resourceURL(c)
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}
