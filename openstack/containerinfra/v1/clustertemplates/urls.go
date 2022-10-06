package clustertemplates

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

var apiName = "clustertemplates"

func commonURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiName)
}

func idURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiName, id)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return commonURL(client)
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return commonURL(client)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}

func updateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}
