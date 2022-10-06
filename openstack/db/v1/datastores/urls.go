package datastores

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func baseURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("datastores")
}

func resourceURL(c *ktvpcsdk.ServiceClient, dsID string) string {
	return c.ServiceURL("datastores", dsID)
}

func versionsURL(c *ktvpcsdk.ServiceClient, dsID string) string {
	return c.ServiceURL("datastores", dsID, "versions")
}

func versionURL(c *ktvpcsdk.ServiceClient, dsID, versionID string) string {
	return c.ServiceURL("datastores", dsID, "versions", versionID)
}
