package volumeactions

import "github.com/innodreamer/ktvpc-sdk_poc"

func actionURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}
