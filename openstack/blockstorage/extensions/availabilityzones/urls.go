package availabilityzones

import "github.com/innodreamer/ktvpc-sdk_poc"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
