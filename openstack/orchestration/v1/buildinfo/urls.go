package buildinfo

import "github.com/innodreamer/ktvpc-sdk_poc"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
