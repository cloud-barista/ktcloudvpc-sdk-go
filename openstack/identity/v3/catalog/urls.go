package catalog

import "github.com/innodreamer/ktvpc-sdk_poc"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("auth", "catalog")
}
