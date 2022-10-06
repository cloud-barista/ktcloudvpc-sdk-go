package trusts

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "OS-TRUST/trusts"

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func createURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func listRolesURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, "roles")
}

func getRoleURL(c *ktvpcsdk.ServiceClient, id, roleID string) string {
	return c.ServiceURL(resourcePath, id, "roles", roleID)
}
