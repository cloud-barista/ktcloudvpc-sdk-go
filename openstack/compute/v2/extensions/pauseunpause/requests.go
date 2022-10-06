package pauseunpause

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/compute/v2/extensions"
)

// Pause is the operation responsible for pausing a Compute server.
func Pause(client *ktvpcsdk.ServiceClient, id string) (r PauseResult) {
	resp, err := client.Post(extensions.ActionURL(client, id), map[string]interface{}{"pause": nil}, nil, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}

// Unpause is the operation responsible for unpausing a Compute server.
func Unpause(client *ktvpcsdk.ServiceClient, id string) (r UnpauseResult) {
	resp, err := client.Post(extensions.ActionURL(client, id), map[string]interface{}{"unpause": nil}, nil, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
