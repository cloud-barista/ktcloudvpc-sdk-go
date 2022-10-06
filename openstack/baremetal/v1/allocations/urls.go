package allocations

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("allocations")
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return createURL(client)
}

func resourceURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("allocations", id)
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(client, id)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(client, id)
}
