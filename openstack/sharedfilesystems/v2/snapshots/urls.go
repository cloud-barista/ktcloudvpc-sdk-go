package snapshots

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("snapshots")
}

func listDetailURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("snapshots", "detail")
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id)
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id)
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id)
}
