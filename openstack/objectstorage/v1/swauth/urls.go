package swauth

import "github.com/innodreamer/ktvpc-sdk_poc"

func getURL(c *gophercloud.ProviderClient) string {
	return c.IdentityBase + "auth/v1.0"
}
