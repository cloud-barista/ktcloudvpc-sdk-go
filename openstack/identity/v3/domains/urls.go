package domains

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("domains")
}

func getURL(client *ktvpcsdk.ServiceClient, domainID string) string {
	return client.ServiceURL("domains", domainID)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("domains")
}

func deleteURL(client *ktvpcsdk.ServiceClient, domainID string) string {
	return client.ServiceURL("domains", domainID)
}

func updateURL(client *ktvpcsdk.ServiceClient, domainID string) string {
	return client.ServiceURL("domains", domainID)
}
