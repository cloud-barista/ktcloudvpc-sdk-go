package attributestags

import "github.com/innodreamer/ktvpc-sdk_poc"

const (
	tagsPath = "tags"
)

func replaceURL(c *gophercloud.ServiceClient, r_type string, id string) string {
	return c.ServiceURL(r_type, id, tagsPath)
}

func listURL(c *gophercloud.ServiceClient, r_type string, id string) string {
	return c.ServiceURL(r_type, id, tagsPath)
}

func deleteAllURL(c *gophercloud.ServiceClient, r_type string, id string) string {
	return c.ServiceURL(r_type, id, tagsPath)
}

func addURL(c *gophercloud.ServiceClient, r_type string, id string, tag string) string {
	return c.ServiceURL(r_type, id, tagsPath, tag)
}

func deleteURL(c *gophercloud.ServiceClient, r_type string, id string, tag string) string {
	return c.ServiceURL(r_type, id, tagsPath, tag)
}

func confirmURL(c *gophercloud.ServiceClient, r_type string, id string, tag string) string {
	return c.ServiceURL(r_type, id, tagsPath, tag)
}
