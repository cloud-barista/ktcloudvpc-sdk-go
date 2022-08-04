package webhooks

import "github.com/innodreamer/ktvpc-sdk_poc"

func triggerURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("v1", "webhooks", id, "trigger")
}
