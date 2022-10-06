package policies

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath     = "fwaas"
	resourcePath = "firewall_policies"
	insertPath   = "insert_rule"
	removePath   = "remove_rule"
)

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

func insertURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, insertPath)
}

func removeURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, removePath)
}
