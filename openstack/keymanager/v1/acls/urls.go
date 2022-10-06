package acls

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func containerURL(client *ktvpcsdk.ServiceClient, containerID string) string {
	return client.ServiceURL("containers", containerID, "acl")
}

func secretURL(client *ktvpcsdk.ServiceClient, secretID string) string {
	return client.ServiceURL("secrets", secretID, "acl")
}
