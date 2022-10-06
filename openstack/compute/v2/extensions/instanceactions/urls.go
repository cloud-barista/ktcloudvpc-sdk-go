package instanceactions

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "os-instance-actions")
}

func instanceActionsURL(client *ktvpcsdk.ServiceClient, serverID, requestID string) string {
	return client.ServiceURL("servers", serverID, "os-instance-actions", requestID)
}
