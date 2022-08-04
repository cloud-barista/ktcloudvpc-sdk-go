package tokens

import "github.com/innodreamer/ktvpc-sdk_poc"

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
