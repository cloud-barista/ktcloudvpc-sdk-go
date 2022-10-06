package resourceproviders

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	apiName = "resource_providers"
)

func resourceProvidersListURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiName)
}

func deleteURL(client *ktvpcsdk.ServiceClient, resourceProviderID string) string {
	return client.ServiceURL(apiName, resourceProviderID)
}

func getURL(client *ktvpcsdk.ServiceClient, resourceProviderID string) string {
	return client.ServiceURL(apiName, resourceProviderID)
}

func updateURL(client *ktvpcsdk.ServiceClient, resourceProviderID string) string {
	return client.ServiceURL(apiName, resourceProviderID)
}

func getResourceProviderUsagesURL(client *ktvpcsdk.ServiceClient, resourceProviderID string) string {
	return client.ServiceURL(apiName, resourceProviderID, "usages")
}

func getResourceProviderInventoriesURL(client *ktvpcsdk.ServiceClient, resourceProviderID string) string {
	return client.ServiceURL(apiName, resourceProviderID, "inventories")
}

func getResourceProviderAllocationsURL(client *ktvpcsdk.ServiceClient, resourceProviderID string) string {
	return client.ServiceURL(apiName, resourceProviderID, "allocations")
}

func getResourceProviderTraitsURL(client *ktvpcsdk.ServiceClient, resourceProviderID string) string {
	return client.ServiceURL(apiName, resourceProviderID, "traits")
}
