package resetnetwork

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/compute/v2/extensions"
)

// ResetNetwork will reset the network of a server
func ResetNetwork(client *ktvpcsdk.ServiceClient, id string) (r ResetResult) {
	b := map[string]interface{}{
		"resetNetwork": nil,
	}
	resp, err := client.Post(extensions.ActionURL(client, id), b, nil, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
