package secgroups

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	secgrouppath = "os-security-groups"
	rulepath     = "os-security-group-rules"
)

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(secgrouppath, id)
}

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(secgrouppath)
}

func listByServerURL(c *ktvpcsdk.ServiceClient, serverID string) string {
	return c.ServiceURL("servers", serverID, secgrouppath)
}

func rootRuleURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(rulepath)
}

func resourceRuleURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rulepath, id)
}

func serverActionURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("servers", id, "action")
}
