package extensions

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

// ExtensionURL generates the URL for an extension resource by name.
func ExtensionURL(c *ktvpcsdk.ServiceClient, name string) string {
	return c.ServiceURL("extensions", name)
}

// ListExtensionURL generates the URL for the extensions resource collection.
func ListExtensionURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL("extensions")
}
