package stackresources

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func findURL(c *ktvpcsdk.ServiceClient, stackName string) string {
	return c.ServiceURL("stacks", stackName, "resources")
}

func listURL(c *ktvpcsdk.ServiceClient, stackName, stackID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources")
}

func getURL(c *ktvpcsdk.ServiceClient, stackName, stackID, resourceName string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources", resourceName)
}

func metadataURL(c *ktvpcsdk.ServiceClient, stackName, stackID, resourceName string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources", resourceName, "metadata")
}

func listTypesURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("resource_types")
}

func schemaURL(c *ktvpcsdk.ServiceClient, typeName string) string {
	return c.ServiceURL("resource_types", typeName)
}

func templateURL(c *ktvpcsdk.ServiceClient, typeName string) string {
	return c.ServiceURL("resource_types", typeName, "template")
}

func markUnhealthyURL(c *ktvpcsdk.ServiceClient, stackName, stackID, resourceName string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources", resourceName)
}
