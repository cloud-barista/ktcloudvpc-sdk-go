package pools

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath     = "lbaas"
	resourcePath = "pools"
	memberPath   = "members"
)

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

func memberRootURL(c *ktvpcsdk.ServiceClient, poolId string) string {
	return c.ServiceURL(rootPath, resourcePath, poolId, memberPath)
}

func memberResourceURL(c *ktvpcsdk.ServiceClient, poolID string, memberID string) string {
	return c.ServiceURL(rootPath, resourcePath, poolID, memberPath, memberID)
}
