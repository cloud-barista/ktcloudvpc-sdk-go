package claims

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	apiVersion = "v2"
	apiName    = "queues"
)

func createURL(client *ktvpcsdk.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "claims")
}

func getURL(client *ktvpcsdk.ServiceClient, queueName string, claimID string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "claims", claimID)
}

func updateURL(client *ktvpcsdk.ServiceClient, queueName string, claimID string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "claims", claimID)
}

func deleteURL(client *ktvpcsdk.ServiceClient, queueName string, claimID string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "claims", claimID)
}
