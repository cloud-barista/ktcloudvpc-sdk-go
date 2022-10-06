package quotas

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "quotas"
const resourcePathDetail = "details.json"

func resourceURL(c *ktvpcsdk.ServiceClient, projectID string) string {
	return c.ServiceURL(resourcePath, projectID)
}

func resourceDetailURL(c *ktvpcsdk.ServiceClient, projectID string) string {
	return c.ServiceURL(resourcePath, projectID, resourcePathDetail)
}

func getURL(c *ktvpcsdk.ServiceClient, projectID string) string {
	return resourceURL(c, projectID)
}

func getDetailURL(c *ktvpcsdk.ServiceClient, projectID string) string {
	return resourceDetailURL(c, projectID)
}

func updateURL(c *ktvpcsdk.ServiceClient, projectID string) string {
	return resourceURL(c, projectID)
}
