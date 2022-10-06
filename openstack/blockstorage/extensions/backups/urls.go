package backups

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("backups")
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("backups", id)
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("backups", id)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("backups")
}

func listDetailURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("backups", "detail")
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("backups", id)
}

func restoreURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("backups", id, "restore")
}

func exportURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL("backups", id, "export_record")
}

func importURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("backups", "import_record")
}
