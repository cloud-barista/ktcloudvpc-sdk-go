package recordsets

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func baseURL(c *ktvpcsdk.ServiceClient, zoneID string) string {
	return c.ServiceURL("zones", zoneID, "recordsets")
}

func rrsetURL(c *ktvpcsdk.ServiceClient, zoneID string, rrsetID string) string {
	return c.ServiceURL("zones", zoneID, "recordsets", rrsetID)
}
