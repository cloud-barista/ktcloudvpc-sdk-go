package portforwarding

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

const resourcePath = "Portforwarding"									// Modified

func portForwardingUrl(c *gophercloud.ServiceClient) string {			// Modified
	return c.ServiceURL(resourcePath)
}

func singlePortForwardingUrl(c *gophercloud.ServiceClient, portForwardingID string) string {	// Modified
	return c.ServiceURL(resourcePath, portForwardingID)
}
