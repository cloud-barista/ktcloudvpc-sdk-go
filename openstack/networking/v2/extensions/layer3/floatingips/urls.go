package floatingips

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "floatingips"

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}
