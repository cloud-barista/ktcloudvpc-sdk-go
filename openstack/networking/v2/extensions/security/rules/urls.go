package rules

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const rootPath = "security-group-rules"

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(rootPath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, id)
}
