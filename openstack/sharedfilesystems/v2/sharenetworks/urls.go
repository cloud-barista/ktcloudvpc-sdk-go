package sharenetworks

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("share-networks")
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("share-networks", id)
}

func listDetailURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("share-networks", "detail")
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func addSecurityServiceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("share-networks", id, "action")
}

func removeSecurityServiceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("share-networks", id, "action")
}
