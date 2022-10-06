package services

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("os-services")
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("os-services", id)
}
