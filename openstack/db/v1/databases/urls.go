package databases

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func baseURL(c *ktvpcsdk.ServiceClient, instanceID string) string {
	return c.ServiceURL("instances", instanceID, "databases")
}

func dbURL(c *ktvpcsdk.ServiceClient, instanceID, dbName string) string {
	return c.ServiceURL("instances", instanceID, "databases", dbName)
}
