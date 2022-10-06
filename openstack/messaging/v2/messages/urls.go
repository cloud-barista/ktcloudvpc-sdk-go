package messages

import (
	"net/url"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

const (
	apiVersion = "v2"
	apiName    = "queues"
)

func createURL(client *ktvpcsdk.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "messages")
}

func listURL(client *ktvpcsdk.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "messages")
}

func getURL(client *ktvpcsdk.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "messages")
}

func deleteURL(client *ktvpcsdk.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "messages")
}

func DeleteMessageURL(client *ktvpcsdk.ServiceClient, queueName string, messageID string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "messages", messageID)
}

func messageURL(client *ktvpcsdk.ServiceClient, queueName string, messageID string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "messages", messageID)
}

// Builds next page full url based on current url.
func nextPageURL(currentURL string, next string) (string, error) {
	base, err := url.Parse(currentURL)
	if err != nil {
		return "", err
	}
	rel, err := url.Parse(next)
	if err != nil {
		return "", err
	}
	return base.ResolveReference(rel).String(), nil
}
