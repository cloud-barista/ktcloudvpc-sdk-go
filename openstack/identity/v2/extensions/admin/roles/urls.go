package roles

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	ExtPath  = "OS-KSADM"
	RolePath = "roles"
	UserPath = "users"
)

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(ExtPath, RolePath, id)
}

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(ExtPath, RolePath)
}

func userRoleURL(c *ktvpcsdk.ServiceClient, tenantID, userID, roleID string) string {
	return c.ServiceURL("tenants", tenantID, UserPath, userID, RolePath, ExtPath, roleID)
}
