package loadbalancers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath       = "lbaas"
	resourcePath   = "loadbalancers"
	statusPath     = "status"
	statisticsPath = "stats"
	failoverPath   = "failover"
)

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

func statusRootURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, statusPath)
}

func statisticsRootURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, statisticsPath)
}

func failoverRootURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, failoverPath)
}
