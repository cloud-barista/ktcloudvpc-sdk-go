package credentials

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("credentials")
}

func getURL(client *ktvpcsdk.ServiceClient, credentialID string) string {
	return client.ServiceURL("credentials", credentialID)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("credentials")
}

func deleteURL(client *ktvpcsdk.ServiceClient, credentialID string) string {
	return client.ServiceURL("credentials", credentialID)
}

func updateURL(client *ktvpcsdk.ServiceClient, credentialID string) string {
	return client.ServiceURL("credentials", credentialID)
}
