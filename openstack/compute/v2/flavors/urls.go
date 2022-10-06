package flavors

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("flavors", "detail")
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("flavors")
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}

func accessURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "os-flavor-access")
}

func accessActionURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "action")
}

func extraSpecsListURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs")
}

func extraSpecsGetURL(client *ktvpcsdk.ServiceClient, id, key string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs", key)
}

func extraSpecsCreateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs")
}

func extraSpecUpdateURL(client *ktvpcsdk.ServiceClient, id, key string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs", key)
}

func extraSpecDeleteURL(client *ktvpcsdk.ServiceClient, id, key string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs", key)
}
