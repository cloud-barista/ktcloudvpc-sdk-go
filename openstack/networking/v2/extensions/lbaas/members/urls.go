package members

import "github.com/innodreamer/ktvpc-sdk_poc"

const (
	rootPath     = "lb"
	resourcePath = "members"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
