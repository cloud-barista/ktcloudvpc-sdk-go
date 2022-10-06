package zones

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func baseURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("zones")
}

func zoneURL(c *ktvpcsdk.ServiceClient, zoneID string) string {
	return c.ServiceURL("zones", zoneID)
}
