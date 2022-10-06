package oauth1

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func consumersURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("OS-OAUTH1", "consumers")
}

func consumerURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("OS-OAUTH1", "consumers", id)
}

func requestTokenURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("OS-OAUTH1", "request_token")
}

func authorizeTokenURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("OS-OAUTH1", "authorize", id)
}

func createAccessTokenURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("OS-OAUTH1", "access_token")
}

func userAccessTokensURL(c *ktvpcsdk.ServiceClient, userID string) string {
	return c.ServiceURL("users", userID, "OS-OAUTH1", "access_tokens")
}

func userAccessTokenURL(c *ktvpcsdk.ServiceClient, userID string, id string) string {
	return c.ServiceURL("users", userID, "OS-OAUTH1", "access_tokens", id)
}

func userAccessTokenRolesURL(c *ktvpcsdk.ServiceClient, userID string, id string) string {
	return c.ServiceURL("users", userID, "OS-OAUTH1", "access_tokens", id, "roles")
}

func userAccessTokenRoleURL(c *ktvpcsdk.ServiceClient, userID string, id string, roleID string) string {
	return c.ServiceURL("users", userID, "OS-OAUTH1", "access_tokens", id, "roles", roleID)
}

func authURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
