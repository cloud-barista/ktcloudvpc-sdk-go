package limits

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

const resourcePath = "limits"

func getURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}
