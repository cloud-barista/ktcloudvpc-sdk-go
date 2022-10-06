package qos

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id)
}

func createURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("qos-specs")
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("qos-specs")
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("qos-specs", id)
}

func updateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id)
}

func deleteKeysURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id, "delete_keys")
}

func associateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id, "associate")
}

func disassociateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id, "disassociate")
}

func disassociateAllURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id, "disassociate_all")
}

func listAssociationsURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("qos-specs", id, "associations")
}
