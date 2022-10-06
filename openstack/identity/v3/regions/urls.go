package regions

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("regions")
}

func getURL(client *ktvpcsdk.ServiceClient, regionID string) string {
	return client.ServiceURL("regions", regionID)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("regions")
}

func updateURL(client *ktvpcsdk.ServiceClient, regionID string) string {
	return client.ServiceURL("regions", regionID)
}

func deleteURL(client *ktvpcsdk.ServiceClient, regionID string) string {
	return client.ServiceURL("regions", regionID)
}
