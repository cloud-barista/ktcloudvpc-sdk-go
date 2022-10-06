package nodes

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("nodes")
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return createURL(client)
}

func listDetailURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("nodes", "detail")
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("nodes", id)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return deleteURL(client, id)
}

func updateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return deleteURL(client, id)
}

func validateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("nodes", id, "validate")
}

func injectNMIURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("nodes", id, "management", "inject_nmi")
}

func bootDeviceURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("nodes", id, "management", "boot_device")
}

func supportedBootDeviceURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("nodes", id, "management", "boot_device", "supported")
}

func statesResourceURL(client *ktvpcsdk.ServiceClient, id string, state string) string {
	return client.ServiceURL("nodes", id, "states", state)
}

func powerStateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return statesResourceURL(client, id, "power")
}

func provisionStateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return statesResourceURL(client, id, "provision")
}

func raidConfigURL(client *ktvpcsdk.ServiceClient, id string) string {
	return statesResourceURL(client, id, "raid")
}

func biosListSettingsURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("nodes", id, "bios")
}

func biosGetSettingURL(client *ktvpcsdk.ServiceClient, id string, setting string) string {
	return client.ServiceURL("nodes", id, "bios", setting)
}

func vendorPassthruMethodsURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("nodes", id, "vendor_passthru", "methods")
}

func vendorPassthruCallURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("nodes", id, "vendor_passthru")
}

func maintenanceURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("nodes", id, "maintenance")
}
