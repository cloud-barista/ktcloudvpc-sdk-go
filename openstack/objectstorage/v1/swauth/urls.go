package swauth

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func getURL(c *gophercloud.ProviderClient) string {
	return c.IdentityBase + "auth/v1.0"
}
