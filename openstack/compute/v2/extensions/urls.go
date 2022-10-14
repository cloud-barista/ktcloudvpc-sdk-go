package extensions

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func ActionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
