package policies

import "github.com/innodreamer/ktvpc-sdk_poc"

const (
	apiVersion = "v1"
	apiName    = "policies"
)

func policyListURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func policyCreateURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func policyDeleteURL(client *gophercloud.ServiceClient, policyID string) string {
	return client.ServiceURL(apiVersion, apiName, policyID)
}

func policyGetURL(client *gophercloud.ServiceClient, policyID string) string {
	return client.ServiceURL(apiVersion, apiName, policyID)
}

func updateURL(client *gophercloud.ServiceClient, policyID string) string {
	return client.ServiceURL(apiVersion, apiName, policyID)
}

func validateURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName, "validate")
}
