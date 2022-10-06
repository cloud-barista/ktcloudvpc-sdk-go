package profiletypes

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	apiVersion = "v1"
	apiName    = "profile-types"
)

func commonURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func profileTypeURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return profileTypeURL(client, id)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return commonURL(client)
}

func listOpsURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id, "ops")
}
