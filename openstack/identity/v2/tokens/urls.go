package tokens

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

// CreateURL generates the URL used to create new Tokens.
func CreateURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("tokens")
}

// GetURL generates the URL used to Validate Tokens.
func GetURL(client *ktvpcsdk.ServiceClient, token string) string {
	return client.ServiceURL("tokens", token)
}
