package clusters

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

var apiVersion = "v1"
var apiName = "clusters"

func commonURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func idURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id)
}

func actionURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id, "actions")
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return commonURL(client)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return commonURL(client)
}

func updateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}

func listPoliciesURL(client *ktvpcsdk.ServiceClient, clusterID string) string {
	return client.ServiceURL(apiVersion, apiName, clusterID, "policies")
}

func getPolicyURL(client *ktvpcsdk.ServiceClient, clusterID string, policyID string) string {
	return client.ServiceURL(apiVersion, apiName, clusterID, "policies", policyID)
}

func nodeURL(client *ktvpcsdk.ServiceClient, id string) string {
	return actionURL(client, id)
}

func collectURL(client *ktvpcsdk.ServiceClient, clusterID string, path string) string {
	return client.ServiceURL(apiVersion, apiName, clusterID, "attrs", path)
}

func opsURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id, "ops")
}
