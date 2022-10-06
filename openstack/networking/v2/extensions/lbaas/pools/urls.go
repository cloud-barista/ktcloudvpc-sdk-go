package pools

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath     = "lb"
	resourcePath = "pools"
	monitorPath  = "health_monitors"
)

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

func associateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, monitorPath)
}

func disassociateURL(c *ktvpcsdk.ServiceClient, poolID, monitorID string) string {
	return c.ServiceURL(rootPath, resourcePath, poolID, monitorPath, monitorID)
}
