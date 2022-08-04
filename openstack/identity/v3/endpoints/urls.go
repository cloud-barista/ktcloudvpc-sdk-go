package endpoints

import "github.com/innodreamer/ktvpc-sdk_poc"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("endpoints")
}

func endpointURL(client *gophercloud.ServiceClient, endpointID string) string {
	return client.ServiceURL("endpoints", endpointID)
}
