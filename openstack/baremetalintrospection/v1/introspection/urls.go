package introspection

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listIntrospectionsURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("introspection")
}

func introspectionURL(client *ktvpcsdk.ServiceClient, nodeID string) string {
	return client.ServiceURL("introspection", nodeID)
}

func abortIntrospectionURL(client *ktvpcsdk.ServiceClient, nodeID string) string {
	return client.ServiceURL("introspection", nodeID, "abort")
}

func introspectionDataURL(client *ktvpcsdk.ServiceClient, nodeID string) string {
	return client.ServiceURL("introspection", nodeID, "data")
}

func introspectionUnprocessedDataURL(client *ktvpcsdk.ServiceClient, nodeID string) string {
	return client.ServiceURL("introspection", nodeID, "data", "unprocessed")
}
