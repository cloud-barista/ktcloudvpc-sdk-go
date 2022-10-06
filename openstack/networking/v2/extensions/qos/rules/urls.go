package rules

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath = "qos/policies"

	bandwidthLimitRulesResourcePath   = "bandwidth_limit_rules"
	dscpMarkingRulesResourcePath      = "dscp_marking_rules"
	minimumBandwidthRulesResourcePath = "minimum_bandwidth_rules"
)

func bandwidthLimitRulesRootURL(c *ktvpcsdk.ServiceClient, policyID string) string {
	return c.ServiceURL(rootPath, policyID, bandwidthLimitRulesResourcePath)
}

func bandwidthLimitRulesResourceURL(c *ktvpcsdk.ServiceClient, policyID, ruleID string) string {
	return c.ServiceURL(rootPath, policyID, bandwidthLimitRulesResourcePath, ruleID)
}

func listBandwidthLimitRulesURL(c *ktvpcsdk.ServiceClient, policyID string) string {
	return bandwidthLimitRulesRootURL(c, policyID)
}

func getBandwidthLimitRuleURL(c *ktvpcsdk.ServiceClient, policyID, ruleID string) string {
	return bandwidthLimitRulesResourceURL(c, policyID, ruleID)
}

func createBandwidthLimitRuleURL(c *ktvpcsdk.ServiceClient, policyID string) string {
	return bandwidthLimitRulesRootURL(c, policyID)
}

func updateBandwidthLimitRuleURL(c *ktvpcsdk.ServiceClient, policyID, ruleID string) string {
	return bandwidthLimitRulesResourceURL(c, policyID, ruleID)
}

func deleteBandwidthLimitRuleURL(c *ktvpcsdk.ServiceClient, policyID, ruleID string) string {
	return bandwidthLimitRulesResourceURL(c, policyID, ruleID)
}

func dscpMarkingRulesRootURL(c *ktvpcsdk.ServiceClient, policyID string) string {
	return c.ServiceURL(rootPath, policyID, dscpMarkingRulesResourcePath)
}

func dscpMarkingRulesResourceURL(c *ktvpcsdk.ServiceClient, policyID, ruleID string) string {
	return c.ServiceURL(rootPath, policyID, dscpMarkingRulesResourcePath, ruleID)
}

func listDSCPMarkingRulesURL(c *ktvpcsdk.ServiceClient, policyID string) string {
	return dscpMarkingRulesRootURL(c, policyID)
}

func getDSCPMarkingRuleURL(c *ktvpcsdk.ServiceClient, policyID, ruleID string) string {
	return dscpMarkingRulesResourceURL(c, policyID, ruleID)
}

func createDSCPMarkingRuleURL(c *ktvpcsdk.ServiceClient, policyID string) string {
	return dscpMarkingRulesRootURL(c, policyID)
}

func updateDSCPMarkingRuleURL(c *ktvpcsdk.ServiceClient, policyID, ruleID string) string {
	return dscpMarkingRulesResourceURL(c, policyID, ruleID)
}

func deleteDSCPMarkingRuleURL(c *ktvpcsdk.ServiceClient, policyID, ruleID string) string {
	return dscpMarkingRulesResourceURL(c, policyID, ruleID)
}

func minimumBandwidthRulesRootURL(c *ktvpcsdk.ServiceClient, policyID string) string {
	return c.ServiceURL(rootPath, policyID, minimumBandwidthRulesResourcePath)
}

func minimumBandwidthRulesResourceURL(c *ktvpcsdk.ServiceClient, policyID, ruleID string) string {
	return c.ServiceURL(rootPath, policyID, minimumBandwidthRulesResourcePath, ruleID)
}

func listMinimumBandwidthRulesURL(c *ktvpcsdk.ServiceClient, policyID string) string {
	return minimumBandwidthRulesRootURL(c, policyID)
}

func getMinimumBandwidthRuleURL(c *ktvpcsdk.ServiceClient, policyID, ruleID string) string {
	return minimumBandwidthRulesResourceURL(c, policyID, ruleID)
}

func createMinimumBandwidthRuleURL(c *ktvpcsdk.ServiceClient, policyID string) string {
	return minimumBandwidthRulesRootURL(c, policyID)
}

func updateMinimumBandwidthRuleURL(c *ktvpcsdk.ServiceClient, policyID, ruleID string) string {
	return minimumBandwidthRulesResourceURL(c, policyID, ruleID)
}

func deleteMinimumBandwidthRuleURL(c *ktvpcsdk.ServiceClient, policyID, ruleID string) string {
	return minimumBandwidthRulesResourceURL(c, policyID, ruleID)
}
