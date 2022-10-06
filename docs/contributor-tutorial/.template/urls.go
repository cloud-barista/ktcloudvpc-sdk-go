package RESOURCE

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("resource")
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("resource", id)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("resource")
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("resource", id)
}

func updateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("resource", id)
}
