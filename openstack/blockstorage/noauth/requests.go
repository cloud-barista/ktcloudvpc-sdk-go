package noauth

import (
	"fmt"
	"strings"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

// EndpointOpts specifies a "noauth" Cinder Endpoint.
type EndpointOpts struct {
	// CinderEndpoint [required] is currently only used with "noauth" Cinder.
	// A cinder endpoint with "auth_strategy=noauth" is necessary, for example:
	// http://example.com:8776/v2.
	CinderEndpoint string
}

// NewClient prepares an unauthenticated ProviderClient instance.
func NewClient(options ktvpcsdk.AuthOptions) (*ktvpcsdk.ProviderClient, error) {
	if options.Username == "" {
		options.Username = "admin"
	}
	if options.TenantName == "" {
		options.TenantName = "admin"
	}

	client := &ktvpcsdk.ProviderClient{
		TokenID: fmt.Sprintf("%s:%s", options.Username, options.TenantName),
	}

	return client, nil
}

func initClientOpts(client *ktvpcsdk.ProviderClient, eo EndpointOpts, clientType string) (*ktvpcsdk.ServiceClient, error) {
	sc := new(ktvpcsdk.ServiceClient)
	if eo.CinderEndpoint == "" {
		return nil, fmt.Errorf("CinderEndpoint is required")
	}

	token := strings.Split(client.TokenID, ":")
	if len(token) != 2 {
		return nil, fmt.Errorf("Malformed noauth token")
	}

	endpoint := fmt.Sprintf("%s%s", ktvpcsdk.NormalizeURL(eo.CinderEndpoint), token[1])
	sc.Endpoint = ktvpcsdk.NormalizeURL(endpoint)
	sc.ProviderClient = client
	sc.Type = clientType
	return sc, nil
}

// NewBlockStorageNoAuthV2 creates a ServiceClient that may be used to access "noauth" v2 block storage service.
func NewBlockStorageNoAuthV2(client *ktvpcsdk.ProviderClient, eo EndpointOpts) (*ktvpcsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "volumev2")
}

// NewBlockStorageNoAuthV3 creates a ServiceClient that may be used to access "noauth" v3 block storage service.
func NewBlockStorageNoAuthV3(client *ktvpcsdk.ProviderClient, eo EndpointOpts) (*ktvpcsdk.ServiceClient, error) {
	return initClientOpts(client, eo, "volumev3")
}
