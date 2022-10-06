package volumetypes

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("types")
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}

func createURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("types")
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}

func extraSpecsListURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("types", id, "extra_specs")
}

func extraSpecsGetURL(client *ktvpcsdk.ServiceClient, id, key string) string {
	return client.ServiceURL("types", id, "extra_specs", key)
}

func extraSpecsCreateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("types", id, "extra_specs")
}

func extraSpecUpdateURL(client *ktvpcsdk.ServiceClient, id, key string) string {
	return client.ServiceURL("types", id, "extra_specs", key)
}

func extraSpecDeleteURL(client *ktvpcsdk.ServiceClient, id, key string) string {
	return client.ServiceURL("types", id, "extra_specs", key)
}

func accessURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("types", id, "os-volume-type-access")
}

func accessActionURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("types", id, "action")
}
