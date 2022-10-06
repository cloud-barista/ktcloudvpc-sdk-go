package instances

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func baseURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("instances")
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("instances", id)
}

func userRootURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("instances", id, "root")
}

func actionURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("instances", id, "action")
}
