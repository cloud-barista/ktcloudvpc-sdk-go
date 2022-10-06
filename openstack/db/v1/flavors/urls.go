package flavors

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("flavors")
}
