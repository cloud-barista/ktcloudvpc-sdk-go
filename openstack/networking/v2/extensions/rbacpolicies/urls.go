package rbacpolicies

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("rbac-policies", id)
}

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("rbac-policies")
}

func createURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}
