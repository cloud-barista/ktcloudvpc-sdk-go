package diagnostics

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

// serverDiagnosticsURL returns the diagnostics url for a nova instance/server
func serverDiagnosticsURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "diagnostics")
}
