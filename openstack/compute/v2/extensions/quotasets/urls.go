package quotasets

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "os-quota-sets"

func resourceURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func getURL(c *ktvpcsdk.ServiceClient, tenantID string) string {
	return c.ServiceURL(resourcePath, tenantID)
}

func getDetailURL(c *ktvpcsdk.ServiceClient, tenantID string) string {
	return c.ServiceURL(resourcePath, tenantID, "detail")
}

func updateURL(c *ktvpcsdk.ServiceClient, tenantID string) string {
	return getURL(c, tenantID)
}

func deleteURL(c *ktvpcsdk.ServiceClient, tenantID string) string {
	return getURL(c, tenantID)
}
