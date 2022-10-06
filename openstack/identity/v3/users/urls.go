package users

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("users")
}

func getURL(client *ktvpcsdk.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("users")
}

func updateURL(client *ktvpcsdk.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID)
}

func changePasswordURL(client *ktvpcsdk.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "password")
}

func deleteURL(client *ktvpcsdk.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID)
}

func listGroupsURL(client *ktvpcsdk.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "groups")
}

func addToGroupURL(client *ktvpcsdk.ServiceClient, groupID, userID string) string {
	return client.ServiceURL("groups", groupID, "users", userID)
}

func isMemberOfGroupURL(client *ktvpcsdk.ServiceClient, groupID, userID string) string {
	return client.ServiceURL("groups", groupID, "users", userID)
}

func removeFromGroupURL(client *ktvpcsdk.ServiceClient, groupID, userID string) string {
	return client.ServiceURL("groups", groupID, "users", userID)
}

func listProjectsURL(client *ktvpcsdk.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "projects")
}

func listInGroupURL(client *ktvpcsdk.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID, "users")
}
