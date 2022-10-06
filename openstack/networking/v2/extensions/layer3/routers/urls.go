package routers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "routers"

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func addInterfaceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, "add_router_interface")
}

func removeInterfaceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, "remove_router_interface")
}

func listl3AgentsURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, "l3-agents")
}
