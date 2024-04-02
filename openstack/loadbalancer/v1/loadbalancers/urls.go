package loadbalancers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath       = "?command=listLoadBalancers" // Temperary


	listCommand       = "?command=listLoadBalancers"
	createCommand     = "?command=createLoadBalancer"
	deleteCommand     = "?command=deleteLoadBalancer"
	resourcePath   = ""
	statusPath     = "status"
	statisticsPath = "stats"
	failoverPath   = "failover"
)

// func rootURL(c *gophercloud.ServiceClient) string {
// 	return c.GetServiceURL(rootPath, resourcePath)
// }

func listNlbURL(c *gophercloud.ServiceClient) string {
	return c.GetServiceURL(listCommand, resourcePath)
}

func createNlbURL(c *gophercloud.ServiceClient) string {
	return c.GetServiceURL(createCommand, resourcePath)
}

func deleteNlbURL(c *gophercloud.ServiceClient) string {
	return c.GetServiceURL(deleteCommand, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.GetServiceURL(rootPath, resourcePath, id)
}

func statusRootURL(c *gophercloud.ServiceClient, id string) string {
	return c.GetServiceURL(rootPath, resourcePath, id, statusPath)
}

func statisticsRootURL(c *gophercloud.ServiceClient, id string) string {
	return c.GetServiceURL(rootPath, resourcePath, id, statisticsPath)
}

func failoverRootURL(c *gophercloud.ServiceClient, id string) string {
	return c.GetServiceURL(rootPath, resourcePath, id, failoverPath)
}
