package roles

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rolePath = "roles"
)

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(rolePath)
}

func getURL(client *ktvpcsdk.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(rolePath)
}

func updateURL(client *ktvpcsdk.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func deleteURL(client *ktvpcsdk.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func listAssignmentsURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("role_assignments")
}

func listAssignmentsOnResourceURL(client *ktvpcsdk.ServiceClient, targetType, targetID, actorType, actorID string) string {
	return client.ServiceURL(targetType, targetID, actorType, actorID, rolePath)
}

func assignURL(client *ktvpcsdk.ServiceClient, targetType, targetID, actorType, actorID, roleID string) string {
	return client.ServiceURL(targetType, targetID, actorType, actorID, rolePath, roleID)
}
