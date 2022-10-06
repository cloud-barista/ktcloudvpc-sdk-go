package events

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

var apiVersion = "v1"
var apiName = "events"

func commonURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return commonURL(client)
}

func idURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}
