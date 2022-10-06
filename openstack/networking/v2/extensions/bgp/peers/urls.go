package peers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const urlBase = "bgp-peers"

// return /v2.0/bgp-peers/{bgp-peer-id}
func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(urlBase, id)
}

// return /v2.0/bgp-peers
func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(urlBase)
}

// return /v2.0/bgp-peers/{bgp-peer-id}
func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

// return /v2.0/bgp-peers
func listURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

// return /v2.0/bgp-peers
func createURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

// return /v2.0/bgp-peers/{bgp-peer-id}
func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

// return /v2.0/bgp-peers/{bgp-peer-id}
func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}
