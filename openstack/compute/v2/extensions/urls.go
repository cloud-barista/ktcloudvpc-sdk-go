package extensions

import "github.com/innodreamer/ktvpc-sdk_poc"

func ActionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
