package policies

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const policyPath = "policies"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(policyPath)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(policyPath)
}

func getURL(client *ktvpcsdk.ServiceClient, policyID string) string {
	return client.ServiceURL(policyPath, policyID)
}

func updateURL(client *ktvpcsdk.ServiceClient, policyID string) string {
	return client.ServiceURL(policyPath, policyID)
}

func deleteURL(client *ktvpcsdk.ServiceClient, policyID string) string {
	return client.ServiceURL(policyPath, policyID)
}
