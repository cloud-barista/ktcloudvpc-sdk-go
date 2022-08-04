package remoteconsoles

import "github.com/innodreamer/ktvpc-sdk_poc"

const (
	rootPath = "servers"

	resourcePath = "remote-consoles"
)

func rootURL(c *gophercloud.ServiceClient, serverID string) string {
	return c.ServiceURL(rootPath, serverID, resourcePath)
}

func createURL(c *gophercloud.ServiceClient, serverID string) string {
	return rootURL(c, serverID)
}
