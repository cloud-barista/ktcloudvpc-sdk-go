package accept

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath     = "zones"
	tasksPath    = "tasks"
	resourcePath = "transfer_accepts"
)

func baseURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, tasksPath, resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, transferAcceptID string) string {
	return c.ServiceURL(rootPath, tasksPath, resourcePath, transferAcceptID)
}
