package workflows

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("workflows")
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("workflows", id)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("workflows", id)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("workflows")
}
