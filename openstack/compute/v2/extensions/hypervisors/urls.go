package hypervisors

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func hypervisorsListDetailURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("os-hypervisors", "detail")
}

func hypervisorsStatisticsURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("os-hypervisors", "statistics")
}

func hypervisorsGetURL(c *ktvpcsdk.ServiceClient, hypervisorID string) string {
	return c.ServiceURL("os-hypervisors", hypervisorID)
}

func hypervisorsUptimeURL(c *ktvpcsdk.ServiceClient, hypervisorID string) string {
	return c.ServiceURL("os-hypervisors", hypervisorID, "uptime")
}
