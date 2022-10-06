package policytypes

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	apiVersion = "v1"
	apiName    = "policy-types"
)

func policyTypeListURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func policyTypeGetURL(client *ktvpcsdk.ServiceClient, policyTypeName string) string {
	return client.ServiceURL(apiVersion, apiName, policyTypeName)
}
