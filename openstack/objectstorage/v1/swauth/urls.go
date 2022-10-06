package swauth

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func getURL(c *ktvpcsdk.ProviderClient) string {
	return c.IdentityBase + "auth/v1.0"
}
