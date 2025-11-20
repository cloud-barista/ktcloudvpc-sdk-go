package staticnat

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-go"
)

const resourcePath = "staticNat"									// Added

func staticNatURL(c *gophercloud.ServiceClient) string {			// Added
	return c.ServiceURL(resourcePath)
}

func singleStaticNatURL(c *gophercloud.ServiceClient, staticNatID string) string {	// Added
	return c.ServiceURL(resourcePath, staticNatID)
}
