package tags

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootResourcePath = "servers"
	resourcePath     = "tags"
)

func rootURL(c *ktvpcsdk.ServiceClient, serverID string) string {
	return c.ServiceURL(rootResourcePath, serverID, resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, serverID, tag string) string {
	return c.ServiceURL(rootResourcePath, serverID, resourcePath, tag)
}

func listURL(c *ktvpcsdk.ServiceClient, serverID string) string {
	return rootURL(c, serverID)
}

func checkURL(c *ktvpcsdk.ServiceClient, serverID, tag string) string {
	return resourceURL(c, serverID, tag)
}

func replaceAllURL(c *ktvpcsdk.ServiceClient, serverID string) string {
	return rootURL(c, serverID)
}

func addURL(c *ktvpcsdk.ServiceClient, serverID, tag string) string {
	return resourceURL(c, serverID, tag)
}

func deleteURL(c *ktvpcsdk.ServiceClient, serverID, tag string) string {
	return resourceURL(c, serverID, tag)
}

func deleteAllURL(c *ktvpcsdk.ServiceClient, serverID string) string {
	return rootURL(c, serverID)
}
