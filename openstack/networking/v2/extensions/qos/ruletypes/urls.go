package ruletypes

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listRuleTypesURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("qos", "rule-types")
}

func getRuleTypeURL(c *ktvpcsdk.ServiceClient, name string) string {
	return c.ServiceURL("qos", "rule-types", name)
}
