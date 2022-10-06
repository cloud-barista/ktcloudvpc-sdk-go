package l7policies

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath     = "lbaas"
	resourcePath = "l7policies"
	rulePath     = "rules"
)

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

func ruleRootURL(c *ktvpcsdk.ServiceClient, policyID string) string {
	return c.ServiceURL(rootPath, resourcePath, policyID, rulePath)
}

func ruleResourceURL(c *ktvpcsdk.ServiceClient, policyID string, ruleID string) string {
	return c.ServiceURL(rootPath, resourcePath, policyID, rulePath, ruleID)
}
