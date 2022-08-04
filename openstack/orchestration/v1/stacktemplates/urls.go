package stacktemplates

import "github.com/innodreamer/ktvpc-sdk_poc"

func getURL(c *gophercloud.ServiceClient, stackName, stackID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "template")
}

func validateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("validate")
}
