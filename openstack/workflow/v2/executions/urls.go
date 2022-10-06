package executions

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("executions")
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("executions", id)
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("executions", id)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("executions")
}
