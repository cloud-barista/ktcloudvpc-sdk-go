package usage

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const resourcePath = "os-simple-tenant-usage"

func allTenantsURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(resourcePath)
}

func getTenantURL(client *ktvpcsdk.ServiceClient, tenantID string) string {
	return client.ServiceURL(resourcePath, tenantID)
}
