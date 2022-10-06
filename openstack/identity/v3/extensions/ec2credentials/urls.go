package ec2credentials

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "credentials", "OS-EC2")
}

func getURL(client *ktvpcsdk.ServiceClient, userID string, id string) string {
	return client.ServiceURL("users", userID, "credentials", "OS-EC2", id)
}

func createURL(client *ktvpcsdk.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "credentials", "OS-EC2")
}

func deleteURL(client *ktvpcsdk.ServiceClient, userID string, id string) string {
	return client.ServiceURL("users", userID, "credentials", "OS-EC2", id)
}
