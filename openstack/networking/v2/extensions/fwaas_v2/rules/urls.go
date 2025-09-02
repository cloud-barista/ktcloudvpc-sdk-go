package rules

import "github.com/cloud-barista/ktcloudvpc-sdk-go"

const (
	rootPath     = "firewall"
	resourcePath = "policy"
)

// rootURL generates the URL for the firewall rules resource collection.
func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

// resourceURL generates the URL for a specific firewall rule resource.
func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
