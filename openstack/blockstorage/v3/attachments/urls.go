package attachments

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("attachments")
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("attachments", "detail")
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("attachments", id)
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("attachments", id)
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("attachments", id)
}

func completeURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("attachments", id, "action")
}
