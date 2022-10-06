package configurations

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func baseURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("configurations")
}

func resourceURL(c *ktvpcsdk.ServiceClient, configID string) string {
	return c.ServiceURL("configurations", configID)
}

func instancesURL(c *ktvpcsdk.ServiceClient, configID string) string {
	return c.ServiceURL("configurations", configID, "instances")
}

func listDSParamsURL(c *ktvpcsdk.ServiceClient, datastoreID, versionID string) string {
	return c.ServiceURL("datastores", datastoreID, "versions", versionID, "parameters")
}

func getDSParamURL(c *ktvpcsdk.ServiceClient, datastoreID, versionID, paramID string) string {
	return c.ServiceURL("datastores", datastoreID, "versions", versionID, "parameters", paramID)
}

func listGlobalParamsURL(c *ktvpcsdk.ServiceClient, versionID string) string {
	return c.ServiceURL("datastores", "versions", versionID, "parameters")
}

func getGlobalParamURL(c *ktvpcsdk.ServiceClient, versionID, paramID string) string {
	return c.ServiceURL("datastores", "versions", versionID, "parameters", paramID)
}
