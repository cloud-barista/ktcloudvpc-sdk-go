package portforwarding

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "floatingips"
const portForwardingPath = "port_forwardings"

func portForwardingUrl(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, portForwardingPath)
}

func singlePortForwardingUrl(c *ktvpcsdk.ServiceClient, id string, portForwardingID string) string {
	return c.ServiceURL(resourcePath, id, portForwardingPath, portForwardingID)
}
