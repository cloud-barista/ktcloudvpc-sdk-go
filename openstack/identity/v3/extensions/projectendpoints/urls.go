package projectendpoints

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient, projectID string) string {
	return client.ServiceURL("OS-EP-FILTER", "projects", projectID, "endpoints")
}

func createURL(client *ktvpcsdk.ServiceClient, projectID, endpointID string) string {
	return client.ServiceURL("OS-EP-FILTER", "projects", projectID, "endpoints", endpointID)
}

func deleteURL(client *ktvpcsdk.ServiceClient, projectID, endpointID string) string {
	return client.ServiceURL("OS-EP-FILTER", "projects", projectID, "endpoints", endpointID)
}
