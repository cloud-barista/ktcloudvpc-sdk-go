package quotasets

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "os-quota-sets"

func getURL(c *ktvpcsdk.ServiceClient, projectID string) string {
	return c.ServiceURL(resourcePath, projectID)
}

func getDefaultsURL(c *ktvpcsdk.ServiceClient, projectID string) string {
	return c.ServiceURL(resourcePath, projectID, "defaults")
}

func updateURL(c *ktvpcsdk.ServiceClient, projectID string) string {
	return getURL(c, projectID)
}

func deleteURL(c *ktvpcsdk.ServiceClient, projectID string) string {
	return getURL(c, projectID)
}
