package endpoints

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("endpoints")
}

func endpointURL(client *ktvpcsdk.ServiceClient, endpointID string) string {
	return client.ServiceURL("endpoints", endpointID)
}
