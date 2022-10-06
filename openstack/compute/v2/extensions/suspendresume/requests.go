package suspendresume

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/compute/v2/extensions"
)

// Suspend is the operation responsible for suspending a Compute server.
func Suspend(client *ktvpcsdk.ServiceClient, id string) (r SuspendResult) {
	resp, err := client.Post(extensions.ActionURL(client, id), map[string]interface{}{"suspend": nil}, nil, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// Resume is the operation responsible for resuming a Compute server.
func Resume(client *ktvpcsdk.ServiceClient, id string) (r UnsuspendResult) {
	resp, err := client.Post(extensions.ActionURL(client, id), map[string]interface{}{"resume": nil}, nil, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
