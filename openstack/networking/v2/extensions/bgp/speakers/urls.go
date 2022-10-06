package speakers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const urlBase = "bgp-speakers"

// return /v2.0/bgp-speakers/{bgp-speaker-id}
func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(urlBase, id)
}

// return /v2.0/bgp-speakers
func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(urlBase)
}

// return /v2.0/bgp-speakers/{bgp-speaker-id}
func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

// return /v2.0/bgp-speakers
func listURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

// return /v2.0/bgp-speakers
func createURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

// return /v2.0/bgp-speakers/{bgp-peer-id}
func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

// return /v2.0/bgp-speakers/{bgp-peer-id}
func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

// return /v2.0/bgp-speakers/{bgp-speaker-id}/add_bgp_peer
func addBGPPeerURL(c *ktvpcsdk.ServiceClient, speakerID string) string {
	return c.ServiceURL(urlBase, speakerID, "add_bgp_peer")
}

// return /v2.0/bgp-speakers/{bgp-speaker-id}/remove_bgp_peer
func removeBGPPeerURL(c *ktvpcsdk.ServiceClient, speakerID string) string {
	return c.ServiceURL(urlBase, speakerID, "remove_bgp_peer")
}

// return /v2.0/bgp-speakers/{bgp-speaker-id}/get_advertised_routes
func getAdvertisedRoutesURL(c *ktvpcsdk.ServiceClient, speakerID string) string {
	return c.ServiceURL(urlBase, speakerID, "get_advertised_routes")
}

// return /v2.0/bgp-speakers/{bgp-speaker-id}/add_gateway_network
func addGatewayNetworkURL(c *ktvpcsdk.ServiceClient, speakerID string) string {
	return c.ServiceURL(urlBase, speakerID, "add_gateway_network")
}

// return /v2.0/bgp-speakers/{bgp-speaker-id}/remove_gateway_network
func removeGatewayNetworkURL(c *ktvpcsdk.ServiceClient, speakerID string) string {
	return c.ServiceURL(urlBase, speakerID, "remove_gateway_network")
}
