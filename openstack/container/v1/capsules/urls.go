package capsules

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("capsules", id)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("capsules")
}

// `listURL` is a pure function. `listURL(c)` is a URL for which a GET
// request will respond with a list of capsules in the service `c`.
func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("capsules")
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("capsules", id)
}
