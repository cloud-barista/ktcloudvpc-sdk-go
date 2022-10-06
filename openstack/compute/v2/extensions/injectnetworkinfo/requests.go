package injectnetworkinfo

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/compute/v2/extensions"
)

// InjectNetworkInfo will inject the network info into a server
func InjectNetworkInfo(client *ktvpcsdk.ServiceClient, id string) (r InjectNetworkResult) {
	b := map[string]interface{}{
		"injectNetworkInfo": nil,
	}
	resp, err := client.Post(extensions.ActionURL(client, id), b, nil, nil)
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
