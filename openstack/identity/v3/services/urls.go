package services

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("services")
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("services")
}

func serviceURL(client *ktvpcsdk.ServiceClient, serviceID string) string {
	return client.ServiceURL("services", serviceID)
}

func updateURL(client *ktvpcsdk.ServiceClient, serviceID string) string {
	return client.ServiceURL("services", serviceID)
}
