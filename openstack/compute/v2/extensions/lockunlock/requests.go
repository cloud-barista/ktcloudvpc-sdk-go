package lockunlock

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/compute/v2/extensions"
)

// Lock is the operation responsible for locking a Compute server.
func Lock(client *ktvpcsdk.ServiceClient, id string) (r LockResult) {
	resp, err := client.Post(extensions.ActionURL(client, id), map[string]interface{}{"lock": nil}, nil, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// Unlock is the operation responsible for unlocking a Compute server.
func Unlock(client *ktvpcsdk.ServiceClient, id string) (r UnlockResult) {
	resp, err := client.Post(extensions.ActionURL(client, id), map[string]interface{}{"unlock": nil}, nil, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
