package stacks

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("stacks")
}

func adoptURL(c *ktvpcsdk.ServiceClient) string {
	return createURL(c)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return createURL(c)
}

func getURL(c *ktvpcsdk.ServiceClient, name, id string) string {
	return c.ServiceURL("stacks", name, id)
}

func findURL(c *ktvpcsdk.ServiceClient, identity string) string {
	return c.ServiceURL("stacks", identity)
}

func updateURL(c *ktvpcsdk.ServiceClient, name, id string) string {
	return getURL(c, name, id)
}

func deleteURL(c *ktvpcsdk.ServiceClient, name, id string) string {
	return getURL(c, name, id)
}

func previewURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("stacks", "preview")
}

func abandonURL(c *ktvpcsdk.ServiceClient, name, id string) string {
	return c.ServiceURL("stacks", name, id, "abandon")
}
