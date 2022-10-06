package clusters

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

var apiName = "clusters"

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

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("clusters", id)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("clusters")
}

func listDetailURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("clusters", "detail")
}

func updateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}

func upgradeURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("clusters", id, "actions/upgrade")
}

func resizeURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("clusters", id, "actions/resize")
}
