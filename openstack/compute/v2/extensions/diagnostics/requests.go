package diagnostics

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

// Diagnostics
func Get(client *ktvpcsdk.ServiceClient, serverId string) (r serverDiagnosticsResult) {
	resp, err := client.Get(serverDiagnosticsURL(client, serverId), &r.Body, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
