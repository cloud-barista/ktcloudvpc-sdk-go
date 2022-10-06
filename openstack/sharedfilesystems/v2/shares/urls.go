package shares

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("shares")
}

func listDetailURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("shares", "detail")
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id)
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id)
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id)
}

func listExportLocationsURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "export_locations")
}

func getExportLocationURL(c *ktvpcsdk.ServiceClient, shareID, id string) string {
	return c.ServiceURL("shares", shareID, "export_locations", id)
}

func grantAccessURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "action")
}

func revokeAccessURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "action")
}

func listAccessRightsURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "action")
}

func extendURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "action")
}

func shrinkURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "action")
}

func revertURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "action")
}

func resetStatusURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "action")
}

func forceDeleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "action")
}

func unmanageURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "action")
}

func getMetadataURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "metadata")
}

func getMetadatumURL(c *ktvpcsdk.ServiceClient, id, key string) string {
	return c.ServiceURL("shares", id, "metadata", key)
}

func setMetadataURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "metadata")
}

func updateMetadataURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("shares", id, "metadata")
}

func deleteMetadatumURL(c *ktvpcsdk.ServiceClient, id, key string) string {
	return c.ServiceURL("shares", id, "metadata", key)
}
