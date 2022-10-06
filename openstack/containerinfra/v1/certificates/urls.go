package certificates

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

var apiName = "certificates"

func commonURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiName)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiName, id)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return commonURL(client)
}

func updateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiName, id)
}
