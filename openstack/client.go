// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// Updated by ETRI, 2022.07.

package openstack

import (
	"fmt"
	"reflect"
	"strings"
	// "github.com/davecgh/go-spew/spew"
	cblog "github.com/cloud-barista/cb-log"

	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	tokens2 "github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/identity/v2/tokens"
	tokens3 "github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/identity/v3/tokens"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/identity/v3/extensions/ec2tokens"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/identity/v3/extensions/oauth1"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/utils"
)

const (
	// v2 represents Keystone v2.
	// It should never increase beyond 2.0.
	v2 = "v2.0"

	// v3 represents Keystone v3.
	// The version can be anything from v3 to v3.x.
	v3 = "v3"

	// Added.
	computeEndpoint  = "https://api.ucloudbiz.olleh.com/d1/server/"   // Caution : Need to Add '/' at the end of the endpoint
	imageEndpoint 	 = "https://api.ucloudbiz.olleh.com/d1/image/"  // Caution : Need to Add '/' at the end of the endpoint
	networkEndpoint  = "https://api.ucloudbiz.olleh.com/d1/nsm/v1/"  // Caution : Need to Add '/' at the end of the endpoint
	volumeV2Endpoint = "https://api.ucloudbiz.olleh.com/d1/volume/"  // Caution : Need to Add '/' at the end of the endpoint
	lbV2Endpoint 	 = "https://api.ucloudbiz.olleh.com/d1/loadbalancer/api"  // Caution : Not need to Add '/' at the end of the endpoint
	// ### KT Cloud LB Info API URL ex) : https://api.ucloudbiz.olleh.com/d1/loadbalancer/api?command=listLoadBalancers&...
)

func init() {
	// cblog is a global variable.
	cblogger = cblog.GetLogger("KTCloud VPC Client")
}

/*
NewClient prepares an unauthenticated ProviderClient instance.
Most users will probably prefer using the AuthenticatedClient function
instead.

This is useful if you wish to explicitly control the version of the identity
service that's used for authentication explicitly, for example.

A basic example of using this would be:

	ao, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.NewClient(ao.IdentityEndpoint)
	client, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
*/
func NewClient(endpoint string) (*gophercloud.ProviderClient, error) {
	base, err := utils.BaseEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	endpoint = gophercloud.NormalizeURL(endpoint)
	base = gophercloud.NormalizeURL(base)

	p := new(gophercloud.ProviderClient)
	p.IdentityBase = base
	p.IdentityEndpoint = endpoint
	p.UseTokenLock()

	return p, nil
}

/*
AuthenticatedClient logs in to an OpenStack cloud found at the identity endpoint
specified by the options, acquires a token, and returns a Provider Client
instance that's ready to operate.

If the full path to a versioned identity endpoint was specified  (example:
http://example.com:5000/v3), that path will be used as the endpoint to query.

If a versionless endpoint was specified (example: http://example.com:5000/),
the endpoint will be queried to determine which versions of the identity service
are available, then chooses the most recent or most supported version.

Example:

	ao, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.AuthenticatedClient(ao)
	client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
*/
func AuthenticatedClient(options gophercloud.AuthOptions) (*gophercloud.ProviderClient, error) {
	client, err := NewClient(options.IdentityEndpoint)
	if err != nil {
		return nil, err
	}

	// cblogger.Info("\n# options : ")
	// spew.Dump(options)
	// cblogger.Info("\n\n")

	err = Authenticate(client, options)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Authenticate or re-authenticate against the most recent identity service
// supported at the provided endpoint.
func Authenticate(client *gophercloud.ProviderClient, options gophercloud.AuthOptions) error {
	versions := []*utils.Version{
		{ID: v2, Priority: 20, Suffix: "/v2.0/"},
		{ID: v3, Priority: 30, Suffix: "/v3/"},
	}

	chosen, endpoint, err := utils.ChooseVersion(client, versions)
	if err != nil {
		return err
	}

	// cblogger.Info("\n# version ID : ")
	// spew.Dump(chosen.ID)
	// cblogger.Info("\n\n")

	// cblogger.Info("\n# endpoint : ")
	// spew.Dump(endpoint)
	// cblogger.Info("\n\n")

	switch chosen.ID {
	case v2:
		return v2auth(client, endpoint, options, gophercloud.EndpointOpts{})
	case v3:
		return v3auth(client, endpoint, &options, gophercloud.EndpointOpts{})
	default:
		// The switch statement must be out of date from the versions list.
		return fmt.Errorf("Unrecognized identity version: %s", chosen.ID)
	}
}

// AuthenticateV2 explicitly authenticates against the identity v2 endpoint.
func AuthenticateV2(client *gophercloud.ProviderClient, options gophercloud.AuthOptions, eo gophercloud.EndpointOpts) error {
	return v2auth(client, "", options, eo)
}

func v2auth(client *gophercloud.ProviderClient, endpoint string, options gophercloud.AuthOptions, eo gophercloud.EndpointOpts) error {
	v2Client, err := NewIdentityV2(client, eo)
	if err != nil {
		return err
	}

	if endpoint != "" {
		v2Client.Endpoint = endpoint
	}

	v2Opts := tokens2.AuthOptions{
		IdentityEndpoint: options.IdentityEndpoint,
		Username:         options.Username,
		Password:         options.Password,
		TenantID:         options.TenantID,
		TenantName:       options.TenantName,
		AllowReauth:      options.AllowReauth,
		TokenID:          options.TokenID,
	}

	result := tokens2.Create(v2Client, v2Opts)

	err = client.SetTokenAndAuthResult(result)
	if err != nil {
		return err
	}

	catalog, err := result.ExtractServiceCatalog()
	if err != nil {
		return err
	}

	if options.AllowReauth {
		// here we're creating a throw-away client (tac). it's a copy of the user's provider client, but
		// with the token and reauth func zeroed out. combined with setting `AllowReauth` to `false`,
		// this should retry authentication only once
		tac := *client
		tac.SetThrowaway(true)
		tac.ReauthFunc = nil
		tac.SetTokenAndAuthResult(nil)
		tao := options
		tao.AllowReauth = false
		client.ReauthFunc = func() error {
			err := v2auth(&tac, endpoint, tao, eo)
			if err != nil {
				return err
			}
			client.CopyTokenFrom(&tac)
			return nil
		}
	}
	client.EndpointLocator = func(opts gophercloud.EndpointOpts) (string, error) {
		return V2EndpointURL(catalog, opts)
	}

	return nil
}

// AuthenticateV3 explicitly authenticates against the identity v3 service.
func AuthenticateV3(client *gophercloud.ProviderClient, options tokens3.AuthOptionsBuilder, eo gophercloud.EndpointOpts) error {
	return v3auth(client, "", options, eo)
}

func v3auth(client *gophercloud.ProviderClient, endpoint string, opts tokens3.AuthOptionsBuilder, eo gophercloud.EndpointOpts) error {
	// Override the generated service endpoint with the one returned by the version endpoint.
	v3Client, err := NewIdentityV3(client, eo)
	if err != nil {
		return err
	}

	// cblogger.Info("\n# v3Client : ")
	// spew.Dump(v3Client)
	// cblogger.Info("\n\n")

	if endpoint != "" {
		v3Client.Endpoint = v3Client.IdentityBase   // Modified.
	}
	// ### Original Code
	// if endpoint != "" {
	// 	v3Client.Endpoint = endpoint
	// }
	
	// cblogger.Info("\n# v3Client.Endpoint : ")
	// spew.Dump(v3Client.Endpoint)
	// cblogger.Info("\n\n")

	var catalog *tokens3.ServiceCatalog

	var tokenID string
	// passthroughToken allows to passthrough the token without a scope
	var passthroughToken bool
	switch v := opts.(type) {
	case *gophercloud.AuthOptions:
		tokenID = v.TokenID
		passthroughToken = (v.Scope == nil || *v.Scope == gophercloud.AuthScope{})
	case *tokens3.AuthOptions:
		tokenID = v.TokenID
		passthroughToken = (v.Scope == tokens3.Scope{})
	}

	if tokenID != "" && passthroughToken {
		// passing through the token ID without requesting a new scope
		if opts.CanReauth() {
			return fmt.Errorf("cannot use AllowReauth, when the token ID is defined and auth scope is not set")
		}

		v3Client.SetToken(tokenID)
		result := tokens3.Get(v3Client, tokenID)
		if result.Err != nil {
			return result.Err
		}

		err = client.SetTokenAndAuthResult(result)
		if err != nil {
			return err
		}

		catalog, err = result.ExtractServiceCatalog()
		if err != nil {
			return err
		}
	} else {
		var result tokens3.CreateResult
		switch opts.(type) {
		case *ec2tokens.AuthOptions:
			result = ec2tokens.Create(v3Client, opts)
		case *oauth1.AuthOptions:
			result = oauth1.Create(v3Client, opts)
		default:
			result = tokens3.Create(v3Client, opts)
		}

		err = client.SetTokenAndAuthResult(result)
		if err != nil {
			return err
		}

		catalog, err = result.ExtractServiceCatalog()
		if err != nil {
			return err
		}
	}

	if opts.CanReauth() {
		// here we're creating a throw-away client (tac). it's a copy of the user's provider client, but
		// with the token and reauth func zeroed out. combined with setting `AllowReauth` to `false`,
		// this should retry authentication only once
		tac := *client
		tac.SetThrowaway(true)
		tac.ReauthFunc = nil
		tac.SetTokenAndAuthResult(nil)
		var tao tokens3.AuthOptionsBuilder
		switch ot := opts.(type) {
		case *gophercloud.AuthOptions:
			o := *ot
			o.AllowReauth = false
			tao = &o
		case *tokens3.AuthOptions:
			o := *ot
			o.AllowReauth = false
			tao = &o
		case *ec2tokens.AuthOptions:
			o := *ot
			o.AllowReauth = false
			tao = &o
		case *oauth1.AuthOptions:
			o := *ot
			o.AllowReauth = false
			tao = &o
		default:
			tao = opts
		}
		client.ReauthFunc = func() error {
			err := v3auth(&tac, endpoint, tao, eo)
			if err != nil {
				return err
			}
			client.CopyTokenFrom(&tac)
			return nil
		}
	}
	client.EndpointLocator = func(opts gophercloud.EndpointOpts) (string, error) {
		return V3EndpointURL(catalog, opts)
	}

	return nil
}

// NewIdentityV2 creates a ServiceClient that may be used to interact with the
// v2 identity service.
func NewIdentityV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	endpoint := client.IdentityBase + "v2.0/"
	clientType := "identity"
	var err error
	if !reflect.DeepEqual(eo, gophercloud.EndpointOpts{}) {
		eo.ApplyDefaults(clientType)
		endpoint, err = client.EndpointLocator(eo)
		if err != nil {
			return nil, err
		}
	}

	return &gophercloud.ServiceClient{
		ProviderClient: client,
		Endpoint:       endpoint,
		Type:           clientType,
	}, nil
}

// NewIdentityV3 creates a ServiceClient that may be used to access the v3
// identity service.
func NewIdentityV3(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {

	endpoint := client.IdentityBase  // Modified.
	clientType := "identity"
	var err error
	if !reflect.DeepEqual(eo, gophercloud.EndpointOpts{}) {
		eo.ApplyDefaults(clientType)
		endpoint, err = client.EndpointLocator(eo)
		if err != nil {
			return nil, err
		}
	}

	// Ensure endpoint still has a suffix of v3.
	// This is because EndpointLocator might have found a versionless
	// endpoint or the published endpoint is still /v2.0. In both
	// cases, we need to fix the endpoint to point to /v3.
	base, err := utils.BaseEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	endpoint = gophercloud.NormalizeURL(base)  // Modified.

	// cblogger.Info("\n")
	// cblogger.Infof("\n# base endpoint : %s", endpoint)
	// cblogger.Info("\n\n")

	return &gophercloud.ServiceClient{
		ProviderClient: client,
		Endpoint:       endpoint,
		Type:           clientType,
	}, nil
}

func initClientOpts(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts, clientType string) (*gophercloud.ServiceClient, error) {
	sc := new(gophercloud.ServiceClient)
	eo.ApplyDefaults(clientType)

	// url, err := client.EndpointLocator(eo)
	// if err != nil {
	// 	return sc, err
	// }
	// sc.Endpoint = url

	// ### To get ProjectID from the ProviderClient Object
	authResult := client.GetAuthResult().(tokens3.CreateResult)
	// Type assert to access the Body map
	body, ok := authResult.Body.(map[string]interface{})
	if !ok {
		newErr := fmt.Errorf("Error: Body type assertion failed")
		return nil, newErr
	}
	// Access the token map
	token, ok := body["token"].(map[string]interface{})
	if !ok {
		newErr := fmt.Errorf("Error: Token type assertion failed")
		return nil, newErr
	}
	// Access the project map
	project, ok := token["project"].(map[string]interface{})
	if !ok {
		newErr := fmt.Errorf("Error: Project type assertion failed")
		return nil, newErr
	}
	// Retrieve the project ID
	projectID, ok := project["id"].(string)
	if !ok {
		newErr := fmt.Errorf("Error: Project ID type assertion failed")
		return nil, newErr
	}
	// fmt.Println("# Project ID:", projectID)

	sc.ProviderClient = client
	sc.Type = clientType

	switch eo.Type {   // Added. since Provider Client of KT Cloud does not provide correct endpoint URL.
    case "compute":
		sc.Endpoint = computeEndpoint
    case "image":
		sc.Endpoint = imageEndpoint
    case "network":
		sc.Endpoint = networkEndpoint
    case "volumev2":
		sc.Endpoint = volumeV2Endpoint + projectID + "/"
	case "load-balancer-v1":
		sc.Endpoint = lbV2Endpoint
    }
	// ### KT Cloud Volume Info API URL : https://api.ucloudbiz.olleh.com/d1/volume/{project_id}/volumes/{volume_id}
	// ### KT Cloud LB Info API URL ex) : https://api.ucloudbiz.olleh.com/d1/loadbalancer/api?command=listLoadBalancers&...

	cblogger.Infof("\n# sc.Type in initClientOpts() : %s", sc.Type)
	cblogger.Infof("\n# sc.Endpoint in initClientOpts() : %s", sc.Endpoint)
	cblogger.Info("\n\n")

	// cblogger.Info("\n# sc in initClientOpts() : ")
	// spew.Dump(sc)
	// cblogger.Info("\n\n")

	return sc, nil
}

// NewBareMetalV1 creates a ServiceClient that may be used with the v1
// bare metal package.
func NewBareMetalV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "baremetal")
}

// NewBareMetalIntrospectionV1 creates a ServiceClient that may be used with the v1
// bare metal introspection package.
func NewBareMetalIntrospectionV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "baremetal-inspector")
}

// NewObjectStorageV1 creates a ServiceClient that may be used with the v1
// object storage package.
func NewObjectStorageV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "object-store")
}

// NewComputeV2 creates a ServiceClient that may be used with the v2 compute
// package.
func NewComputeV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "compute")
}

// NewNetworkV2 creates a ServiceClient that may be used with the v2 network
// package.
func NewNetworkV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc, err
}

// NewBlockStorageV1 creates a ServiceClient that may be used to access the v1
// block storage service.
func NewBlockStorageV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "volume")
}

// NewBlockStorageV2 creates a ServiceClient that may be used to access the v2
// block storage service.
func NewBlockStorageV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "volumev2")
}

// NewBlockStorageV3 creates a ServiceClient that may be used to access the v3 block storage service.
func NewBlockStorageV3(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "volumev3")
}

// NewSharedFileSystemV2 creates a ServiceClient that may be used to access the v2 shared file system service.
func NewSharedFileSystemV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "sharev2")
}

// NewCDNV1 creates a ServiceClient that may be used to access the OpenStack v1
// CDN service.
func NewCDNV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "cdn")
}

// NewOrchestrationV1 creates a ServiceClient that may be used to access the v1
// orchestration service.
func NewOrchestrationV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "orchestration")
}

// NewDBV1 creates a ServiceClient that may be used to access the v1 DB service.
func NewDBV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "database")
}

// NewDNSV2 creates a ServiceClient that may be used to access the v2 DNS
// service.
func NewDNSV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "dns")
	sc.ResourceBase = sc.Endpoint + "v2/"
	return sc, err
}

// NewImageServiceV2 creates a ServiceClient that may be used to access the v2
// image service.
func NewImageServiceV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "image")
	sc.ResourceBase = sc.Endpoint   // Modified
	// sc.ResourceBase = sc.Endpoint + "v2/"
	return sc, err
}

// # For KT Cloud D platform Looad Balancer
func NewLoadBalancerV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) { // Added
	sc, err := initClientOpts(client, eo, "load-balancer-v1")
	sc.ResourceBase = sc.Endpoint

	return sc, err
}

// NewLoadBalancerV2 creates a ServiceClient that may be used to access the v2
// load balancer service.
func NewLoadBalancerV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "load-balancer")
	sc.ResourceBase = sc.Endpoint   // Modified
	
	// Fixes edge case having an OpenStack lb endpoint with trailing version number.
	endpoint := strings.Replace(sc.Endpoint, "v2.0/", "", -1)
	sc.ResourceBase = endpoint + "v2.0/"
	return sc, err
}

// NewClusteringV1 creates a ServiceClient that may be used with the v1 clustering
// package.
func NewClusteringV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "clustering")
}

// NewMessagingV2 creates a ServiceClient that may be used with the v2 messaging
// service.
func NewMessagingV2(client *gophercloud.ProviderClient, clientID string, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "messaging")
	sc.MoreHeaders = map[string]string{"Client-ID": clientID}
	return sc, err
}

// NewContainerV1 creates a ServiceClient that may be used with v1 container package
func NewContainerV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "container")
}

// NewKeyManagerV1 creates a ServiceClient that may be used with the v1 key
// manager service.
func NewKeyManagerV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "key-manager")
	sc.ResourceBase = sc.Endpoint + "v1/"
	return sc, err
}

// NewContainerInfraV1 creates a ServiceClient that may be used with the v1 container infra management
// package.
func NewContainerInfraV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "container-infra")
}

// NewWorkflowV2 creates a ServiceClient that may be used with the v2 workflow management package.
func NewWorkflowV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "workflowv2")
}

// NewPlacementV1 creates a ServiceClient that may be used with the placement package.
func NewPlacementV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "placement")
}
