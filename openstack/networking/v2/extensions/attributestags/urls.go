package attributestags

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	tagsPath = "tags"
)

func replaceURL(c *ktvpcsdk.ServiceClient, r_type string, id string) string {
	return c.ServiceURL(r_type, id, tagsPath)
}

func listURL(c *ktvpcsdk.ServiceClient, r_type string, id string) string {
	return c.ServiceURL(r_type, id, tagsPath)
}

func deleteAllURL(c *ktvpcsdk.ServiceClient, r_type string, id string) string {
	return c.ServiceURL(r_type, id, tagsPath)
}

func addURL(c *ktvpcsdk.ServiceClient, r_type string, id string, tag string) string {
	return c.ServiceURL(r_type, id, tagsPath, tag)
}

func deleteURL(c *ktvpcsdk.ServiceClient, r_type string, id string, tag string) string {
	return c.ServiceURL(r_type, id, tagsPath, tag)
}

func confirmURL(c *ktvpcsdk.ServiceClient, r_type string, id string, tag string) string {
	return c.ServiceURL(r_type, id, tagsPath, tag)
}
