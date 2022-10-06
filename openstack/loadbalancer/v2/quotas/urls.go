package quotas

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "quotas"

func resourceURL(c *ktvpcsdk.ServiceClient, projectID string) string {
	return c.ServiceURL(resourcePath, projectID)
}

func getURL(c *ktvpcsdk.ServiceClient, projectID string) string {
	return resourceURL(c, projectID)
}

func updateURL(c *ktvpcsdk.ServiceClient, projectID string) string {
	return resourceURL(c, projectID)
}
