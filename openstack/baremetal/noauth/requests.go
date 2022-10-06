package noauth

import (
	"fmt"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

// EndpointOpts specifies a "noauth" Ironic Endpoint.
type EndpointOpts struct {
	// IronicEndpoint [required] is currently only used with "noauth" Ironic.
	// An Ironic endpoint with "auth_strategy=noauth" is necessary, for example:
	// http://ironic.example.com:6385/v1.
	IronicEndpoint string
}

func initClientOpts(client *ktvpcsdk.ProviderClient, eo EndpointOpts) (*ktvpcsdk.ServiceClient, error) {
	sc := new(ktvpcsdk.ServiceClient)
	if eo.IronicEndpoint == "" {
		return nil, fmt.Errorf("IronicEndpoint is required")
	}

	sc.Endpoint = ktvpcsdk.NormalizeURL(eo.IronicEndpoint)
	sc.ProviderClient = client
	return sc, nil
}

// NewBareMetalNoAuth creates a ServiceClient that may be used to access a
// "noauth" bare metal service.
func NewBareMetalNoAuth(eo EndpointOpts) (*ktvpcsdk.ServiceClient, error) {
	sc, err := initClientOpts(&ktvpcsdk.ProviderClient{}, eo)
	if err != nil {
		return nil, err
	}

	sc.Type = "baremetal"

	return sc, nil
}
