package remoteconsoles

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath = "servers"

	resourcePath = "remote-consoles"
)

func rootURL(c *ktvpcsdk.ServiceClient, serverID string) string {
	return c.ServiceURL(rootPath, serverID, resourcePath)
}

func createURL(c *ktvpcsdk.ServiceClient, serverID string) string {
	return rootURL(c, serverID)
}
