package quotas

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

var apiName = "quotas"

func commonURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiName)
}

func createURL(client *gophercloud.ServiceClient) string {
	return commonURL(client)
}
