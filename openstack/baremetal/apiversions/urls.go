package apiversions

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

func getURL(c *ktvpcsdk.ServiceClient, version string) string {
	return c.ServiceURL(version)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL()
}
