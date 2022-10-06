package drivers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func driversURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("drivers")
}

func driverDetailsURL(client *ktvpcsdk.ServiceClient, driverName string) string {
	return client.ServiceURL("drivers", driverName)
}

func driverPropertiesURL(client *ktvpcsdk.ServiceClient, driverName string) string {
	return client.ServiceURL("drivers", driverName, "properties")
}

func driverDiskPropertiesURL(client *ktvpcsdk.ServiceClient, driverName string) string {
	return client.ServiceURL("drivers", driverName, "raid", "logical_disk_properties")
}
