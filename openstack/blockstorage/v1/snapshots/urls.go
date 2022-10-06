package snapshots

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("snapshots")
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id)
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return createURL(c)
}

func metadataURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id, "metadata")
}

func updateMetadataURL(c *ktvpcsdk.ServiceClient, id string) string {
	return metadataURL(c, id)
}
