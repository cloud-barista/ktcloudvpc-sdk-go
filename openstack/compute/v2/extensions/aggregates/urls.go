package aggregates

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func aggregatesListURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("os-aggregates")
}

func aggregatesCreateURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("os-aggregates")
}

func aggregatesDeleteURL(c *ktvpcsdk.ServiceClient, aggregateID string) string {
	return c.ServiceURL("os-aggregates", aggregateID)
}

func aggregatesGetURL(c *ktvpcsdk.ServiceClient, aggregateID string) string {
	return c.ServiceURL("os-aggregates", aggregateID)
}

func aggregatesUpdateURL(c *ktvpcsdk.ServiceClient, aggregateID string) string {
	return c.ServiceURL("os-aggregates", aggregateID)
}

func aggregatesAddHostURL(c *ktvpcsdk.ServiceClient, aggregateID string) string {
	return c.ServiceURL("os-aggregates", aggregateID, "action")
}

func aggregatesRemoveHostURL(c *ktvpcsdk.ServiceClient, aggregateID string) string {
	return c.ServiceURL("os-aggregates", aggregateID, "action")
}

func aggregatesSetMetadataURL(c *ktvpcsdk.ServiceClient, aggregateID string) string {
	return c.ServiceURL("os-aggregates", aggregateID, "action")
}
