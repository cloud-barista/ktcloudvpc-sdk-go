package queues

import (
	"net/url"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

const (
	apiVersion = "v2"
	apiName    = "queues"
)

func commonURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func createURL(client *ktvpcsdk.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return commonURL(client)
}

func updateURL(client *ktvpcsdk.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName)
}

func getURL(client *ktvpcsdk.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName)
}

func deleteURL(client *ktvpcsdk.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName)
}

func statURL(client *ktvpcsdk.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "stats")
}

func shareURL(client *ktvpcsdk.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "share")
}

func purgeURL(client *ktvpcsdk.ServiceClient, queueName string) string {
	return client.ServiceURL(apiVersion, apiName, queueName, "purge")
}

// builds next page full url based on current url
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
