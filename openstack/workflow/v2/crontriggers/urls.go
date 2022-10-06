package crontriggers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("cron_triggers")
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("cron_triggers", id)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("cron_triggers", id)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("cron_triggers")
}
