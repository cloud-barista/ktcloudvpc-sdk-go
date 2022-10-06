package attachinterfaces

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listInterfaceURL(client *ktvpcsdk.ServiceClient, serverID string) string {
	return client.ServiceURL("servers", serverID, "os-interface")
}

func getInterfaceURL(client *ktvpcsdk.ServiceClient, serverID, portID string) string {
	return client.ServiceURL("servers", serverID, "os-interface", portID)
}

func createInterfaceURL(client *ktvpcsdk.ServiceClient, serverID string) string {
	return client.ServiceURL("servers", serverID, "os-interface")
}
func deleteInterfaceURL(client *ktvpcsdk.ServiceClient, serverID, portID string) string {
	return client.ServiceURL("servers", serverID, "os-interface", portID)
}
