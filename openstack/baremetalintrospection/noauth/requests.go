package noauth

import (
	"fmt"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

// EndpointOpts specifies a "noauth" Ironic Inspector Endpoint.
type EndpointOpts struct {
	// IronicInspectorEndpoint [required] is currently only used with "noauth" Ironic introspection.
	// An Ironic inspector endpoint with "auth_strategy=noauth" is necessary, for example:
	// http://ironic.example.com:5050/v1.
	IronicInspectorEndpoint string
}

func initClientOpts(client *ktvpcsdk.ProviderClient, eo EndpointOpts) (*ktvpcsdk.ServiceClient, error) {
	sc := new(ktvpcsdk.ServiceClient)
	if eo.IronicInspectorEndpoint == "" {
		return nil, fmt.Errorf("IronicInspectorEndpoint is required")
	}

	sc.Endpoint = ktvpcsdk.NormalizeURL(eo.IronicInspectorEndpoint)
	sc.ProviderClient = client
	return sc, nil
}

// NewBareMetalIntrospectionNoAuth creates a ServiceClient that may be used to access a
// "noauth" bare metal introspection service.
func NewBareMetalIntrospectionNoAuth(eo EndpointOpts) (*ktvpcsdk.ServiceClient, error) {
	sc, err := initClientOpts(&ktvpcsdk.ProviderClient{}, eo)
	if err != nil {
		return nil, err
	}

	sc.Type = "baremetal-inspector"

	return sc, nil
}
