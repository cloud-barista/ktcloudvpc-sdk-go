package amphorae

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath     = "octavia"
	resourcePath = "amphorae"
	failoverPath = "failover"
)

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

func failoverRootURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, failoverPath)
}
