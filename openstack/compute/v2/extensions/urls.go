package extensions

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func ActionURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
