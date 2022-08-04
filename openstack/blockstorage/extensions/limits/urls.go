package limits

import (
	"github.com/innodreamer/ktvpc-sdk_poc"
)

const resourcePath = "limits"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}
