package request

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath     = "zones"
	tasksPath    = "tasks"
	resourcePath = "transfer_requests"
)

func baseURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, tasksPath, resourcePath)
}

func createURL(c *ktvpcsdk.ServiceClient, zoneID string) string {
	return c.ServiceURL(rootPath, zoneID, tasksPath, resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, transferID string) string {
	return c.ServiceURL(rootPath, tasksPath, resourcePath, transferID)
}
