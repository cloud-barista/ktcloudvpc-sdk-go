package testing

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/testhelper"
)

func createClient() *ktvpcsdk.ServiceClient {
	return &ktvpcsdk.ServiceClient{
		ProviderClient: &ktvpcsdk.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
