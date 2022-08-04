package floatingips

import "github.com/innodreamer/ktvpc-sdk_poc"

const resourcePath = "floatingips"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}
