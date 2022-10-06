package secrets

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("secrets")
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("secrets", id)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("secrets")
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("secrets", id)
}

func updateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("secrets", id)
}

func payloadURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("secrets", id, "payload")
}

func metadataURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("secrets", id, "metadata")
}

func metadatumURL(client *ktvpcsdk.ServiceClient, id, key string) string {
	return client.ServiceURL("secrets", id, "metadata", key)
}
