package rules

import "github.com/cloud-barista/ktcloudvpc-sdk-go"

const (
	rootPath = "firewall/policy"  							 		// Modified
)

func rootURL(c *gophercloud.ServiceClient) string {				// Modified
	return c.ServiceURL(rootPath)	
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {	// Modified
	return c.ServiceURL(rootPath, id)
}
