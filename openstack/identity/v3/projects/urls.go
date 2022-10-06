package projects

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listAvailableURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("auth", "projects")
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("projects")
}

func getURL(client *ktvpcsdk.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("projects")
}

func deleteURL(client *ktvpcsdk.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

func updateURL(client *ktvpcsdk.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}
