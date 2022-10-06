package tenants

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("tenants")
}

func getURL(client *ktvpcsdk.ServiceClient, tenantID string) string {
	return client.ServiceURL("tenants", tenantID)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("tenants")
}

func deleteURL(client *ktvpcsdk.ServiceClient, tenantID string) string {
	return client.ServiceURL("tenants", tenantID)
}

func updateURL(client *ktvpcsdk.ServiceClient, tenantID string) string {
	return client.ServiceURL("tenants", tenantID)
}
