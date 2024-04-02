package projectendpoints

import "github.com/cloud-barista/ktcloudvpc-sdk-go"

func listURL(client *gophercloud.ServiceClient, projectID string) string {
	return client.ServiceURL("OS-EP-FILTER", "projects", projectID, "endpoints")
}

func createURL(client *gophercloud.ServiceClient, projectID, endpointID string) string {
	return client.ServiceURL("OS-EP-FILTER", "projects", projectID, "endpoints", endpointID)
}

func deleteURL(client *gophercloud.ServiceClient, projectID, endpointID string) string {
	return client.ServiceURL("OS-EP-FILTER", "projects", projectID, "endpoints", endpointID)
}
