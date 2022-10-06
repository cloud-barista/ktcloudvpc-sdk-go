package tokens

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func tokenURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
