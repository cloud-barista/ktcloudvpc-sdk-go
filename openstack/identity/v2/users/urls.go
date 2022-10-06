package users

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	tenantPath = "tenants"
	userPath   = "users"
	rolePath   = "roles"
)

func ResourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(userPath, id)
}

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(userPath)
}

func listRolesURL(c *ktvpcsdk.ServiceClient, tenantID, userID string) string {
	return c.ServiceURL(tenantPath, tenantID, userPath, userID, rolePath)
}
