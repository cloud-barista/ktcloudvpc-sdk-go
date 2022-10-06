package containers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("containers")
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("containers", id)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("containers")
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("containers", id)
}

func listConsumersURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("containers", id, "consumers")
}

func createConsumerURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("containers", id, "consumers")
}

func deleteConsumerURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("containers", id, "consumers")
}

func createSecretRefURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("containers", id, "secrets")
}

func deleteSecretRefURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("containers", id, "secrets")
}
