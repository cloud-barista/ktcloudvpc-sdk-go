package networks

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("networks", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("VPC")  // Modified by B.T. Oh
	// return c.ServiceURL("Network")  // Modified by B.T. Oh

	// return c.ServiceURL("networks") // Origin
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
