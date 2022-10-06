package sharetypes

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("types")
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return createURL(c)
}

func getDefaultURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("types", "default")
}

func getExtraSpecsURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id, "extra_specs")
}

func setExtraSpecsURL(c *ktvpcsdk.ServiceClient, id string) string {
	return getExtraSpecsURL(c, id)
}

func unsetExtraSpecsURL(c *ktvpcsdk.ServiceClient, id string, key string) string {
	return c.ServiceURL("types", id, "extra_specs", key)
}

func showAccessURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id, "share_type_access")
}

func addAccessURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id, "action")
}

func removeAccessURL(c *ktvpcsdk.ServiceClient, id string) string {
	return addAccessURL(c, id)
}
