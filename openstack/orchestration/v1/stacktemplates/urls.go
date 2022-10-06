package stacktemplates

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func getURL(c *ktvpcsdk.ServiceClient, stackName, stackID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "template")
}

func validateURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("validate")
}
