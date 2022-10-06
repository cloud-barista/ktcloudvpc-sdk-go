package servers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("servers")
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return createURL(client)
}

func listDetailURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("servers", "detail")
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return deleteURL(client, id)
}

func updateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return deleteURL(client, id)
}

func actionURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}

func metadatumURL(client *ktvpcsdk.ServiceClient, id, key string) string {
	return client.ServiceURL("servers", id, "metadata", key)
}

func metadataURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "metadata")
}

func listAddressesURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "ips")
}

func listAddressesByNetworkURL(client *ktvpcsdk.ServiceClient, id, network string) string {
	return client.ServiceURL("servers", id, "ips", network)
}

func passwordURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "os-server-password")
}
