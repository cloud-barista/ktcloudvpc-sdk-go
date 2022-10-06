package groups

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("groups")
}

func getURL(client *ktvpcsdk.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("groups")
}

func updateURL(client *ktvpcsdk.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

func deleteURL(client *ktvpcsdk.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}
