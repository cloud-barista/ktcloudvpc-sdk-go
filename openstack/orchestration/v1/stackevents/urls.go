package stackevents

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func findURL(c *ktvpcsdk.ServiceClient, stackName string) string {
	return c.ServiceURL("stacks", stackName, "events")
}

func listURL(c *ktvpcsdk.ServiceClient, stackName, stackID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "events")
}

func listResourceEventsURL(c *ktvpcsdk.ServiceClient, stackName, stackID, resourceName string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources", resourceName, "events")
}

func getURL(c *ktvpcsdk.ServiceClient, stackName, stackID, resourceName, eventID string) string {
	return c.ServiceURL("stacks", stackName, stackID, "resources", resourceName, "events", eventID)
}
