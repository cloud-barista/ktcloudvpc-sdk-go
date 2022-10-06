package agents

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "agents"
const dhcpNetworksResourcePath = "dhcp-networks"
const bgpSpeakersResourcePath = "bgp-drinstances"

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func listURL(c *ktvpcsdk.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func updateURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *ktvpcsdk.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func dhcpNetworksURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, dhcpNetworksResourcePath)
}

func listDHCPNetworksURL(c *ktvpcsdk.ServiceClient, id string) string {
	return dhcpNetworksURL(c, id)
}

func scheduleDHCPNetworkURL(c *ktvpcsdk.ServiceClient, id string) string {
	return dhcpNetworksURL(c, id)
}

func removeDHCPNetworkURL(c *ktvpcsdk.ServiceClient, id string, networkID string) string {
	return c.ServiceURL(resourcePath, id, dhcpNetworksResourcePath, networkID)
}

// return /v2.0/agents/{agent-id}/bgp-drinstances
func listBGPSpeakersURL(c *ktvpcsdk.ServiceClient, agentID string) string {
	return c.ServiceURL(resourcePath, agentID, bgpSpeakersResourcePath)
}
