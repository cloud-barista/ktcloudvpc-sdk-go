package applicationcredentials

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "application_credentials")
}

func getURL(client *ktvpcsdk.ServiceClient, userID string, id string) string {
	return client.ServiceURL("users", userID, "application_credentials", id)
}

func createURL(client *ktvpcsdk.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "application_credentials")
}

func deleteURL(client *ktvpcsdk.ServiceClient, userID string, id string) string {
	return client.ServiceURL("users", userID, "application_credentials", id)
}

func listAccessRulesURL(client *ktvpcsdk.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "access_rules")
}

func getAccessRuleURL(client *ktvpcsdk.ServiceClient, userID string, id string) string {
	return client.ServiceURL("users", userID, "access_rules", id)
}

func deleteAccessRuleURL(client *ktvpcsdk.ServiceClient, userID string, id string) string {
	return client.ServiceURL("users", userID, "access_rules", id)
}
