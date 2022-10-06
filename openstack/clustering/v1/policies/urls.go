package policies

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	apiVersion = "v1"
	apiName    = "policies"
)

func policyListURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func policyCreateURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func policyDeleteURL(client *ktvpcsdk.ServiceClient, policyID string) string {
	return client.ServiceURL(apiVersion, apiName, policyID)
}

func policyGetURL(client *ktvpcsdk.ServiceClient, policyID string) string {
	return client.ServiceURL(apiVersion, apiName, policyID)
}

func updateURL(client *ktvpcsdk.ServiceClient, policyID string) string {
	return client.ServiceURL(apiVersion, apiName, policyID)
}

func validateURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName, "validate")
}
