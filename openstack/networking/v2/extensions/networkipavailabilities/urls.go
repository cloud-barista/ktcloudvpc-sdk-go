package networkipavailabilities

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "network-ip-availabilities"

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, networkIPAvailabilityID string) string {
	return c.ServiceURL(resourcePath, networkIPAvailabilityID)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *ktvpcsdk.ServiceClient, networkIPAvailabilityID string) string {
	return resourceURL(c, networkIPAvailabilityID)
}
