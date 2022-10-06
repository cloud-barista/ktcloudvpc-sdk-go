package ports

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("ports")
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return createURL(client)
}

func listDetailURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("ports", "detail")
}

func resourceURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("ports", id)
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(client, id)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(client, id)
}

func updateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(client, id)
}
