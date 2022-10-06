package nodes

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

var apiVersion = "v1"
var apiName = "nodes"

func commonURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func actionURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id, "actions")
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return commonURL(client)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return commonURL(client)
}

func idURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id)
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}

func updateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}

func opsURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id, "ops")
}
