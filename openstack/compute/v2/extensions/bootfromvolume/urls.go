package bootfromvolume

import "github.com/innodreamer/ktvpc-sdk_poc"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("servers")
}
