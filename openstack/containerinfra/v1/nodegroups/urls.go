package nodegroups

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

func getURL(c *ktvpcsdk.ServiceClient, clusterID, nodeGroupID string) string {
	return c.ServiceURL("clusters", clusterID, "nodegroups", nodeGroupID)
}

func listURL(c *ktvpcsdk.ServiceClient, clusterID string) string {
	return c.ServiceURL("clusters", clusterID, "nodegroups")
}

func createURL(c *ktvpcsdk.ServiceClient, clusterID string) string {
	return c.ServiceURL("clusters", clusterID, "nodegroups")
}

func updateURL(c *ktvpcsdk.ServiceClient, clusterID, nodeGroupID string) string {
	return c.ServiceURL("clusters", clusterID, "nodegroups", nodeGroupID)
}

func deleteURL(c *ktvpcsdk.ServiceClient, clusterID, nodeGroupID string) string {
	return c.ServiceURL("clusters", clusterID, "nodegroups", nodeGroupID)
}
