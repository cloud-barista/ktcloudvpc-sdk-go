package portforwarding

import (
	"strings"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

const resourcePath = "Portforwarding"
const portForwardingPath = "port_forwardings"

func portForwardingUrl(c *gophercloud.ServiceClient) string {
	return strings.Replace(c.ServiceURL(resourcePath, "", ""), "Portforwarding//", "Portforwarding", 1) 
}
// func portForwardingUrl(c *gophercloud.ServiceClient, id string) string {
// 	return c.ServiceURL(resourcePath, id, portForwardingPath)
// }

func singlePortForwardingUrl(c *gophercloud.ServiceClient, id string, portForwardingID string) string {
	return c.ServiceURL(resourcePath, id, portForwardingPath, portForwardingID)
}
