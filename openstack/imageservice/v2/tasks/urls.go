package tasks

import (
	"net/url"
	"strings"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/utils"
)

const resourcePath = "tasks"

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, taskID string) string {
	return c.ServiceURL(resourcePath, taskID)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *ktvpcsdk.ServiceClient, taskID string) string {
	return resourceURL(c, taskID)
}

func createURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

func nextPageURL(serviceURL, requestedNext string) (string, error) {
	base, err := utils.BaseEndpoint(serviceURL)
	if err != nil {
		return "", err
	}

	requestedNextURL, err := url.Parse(requestedNext)
	if err != nil {
		return "", err
	}

	base = ktvpcsdk.NormalizeURL(base)
	nextPath := base + strings.TrimPrefix(requestedNextURL.Path, "/")

	nextURL, err := url.Parse(nextPath)
	if err != nil {
		return "", err
	}

	nextURL.RawQuery = requestedNextURL.RawQuery

	return nextURL.String(), nil
}
