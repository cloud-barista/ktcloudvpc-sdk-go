package resourcetypes

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	resTypesPath = "resource_types"
)

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resTypesPath)
}

func getSchemaURL(c *ktvpcsdk.ServiceClient, resourceType string) string {
	return c.ServiceURL(resTypesPath, resourceType)
}

func generateTemplateURL(c *ktvpcsdk.ServiceClient, resourceType string) string {
	return c.ServiceURL(resTypesPath, resourceType, "template")
}
