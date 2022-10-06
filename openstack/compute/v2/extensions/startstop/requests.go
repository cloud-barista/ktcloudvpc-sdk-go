package startstop

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/compute/v2/extensions"
)

// Start is the operation responsible for starting a Compute server.
func Start(client *ktvpcsdk.ServiceClient, id string) (r StartResult) {
	resp, err := client.Post(extensions.ActionURL(client, id), map[string]interface{}{"os-start": nil}, nil, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// Stop is the operation responsible for stopping a Compute server.
func Stop(client *ktvpcsdk.ServiceClient, id string) (r StopResult) {
	resp, err := client.Post(extensions.ActionURL(client, id), map[string]interface{}{"os-stop": nil}, nil, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
