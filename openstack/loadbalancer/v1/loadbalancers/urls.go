// ### KT Cloud D1 platform > 'NLB (Network Load-Balancer)' handler Go SDK
// Open API Guide : https://cloud.kt.com/docs/open-api-guide/d/network/load-balancer

package loadbalancers

import (
	"strings"

	"github.com/cloud-barista/ktcloudvpc-sdk-go"
)

// The NLBClient endpoint is '.../d1/loadbalancers/' (it ends with a '/').
// See: openstack/client.go -> lbV1Endpoint

// rootURL returns the collection URL '.../d1/loadbalancers' (the trailing '/' is trimmed).
// It is used for the List (GET) and Create (POST) operations.
func rootURL(c *gophercloud.ServiceClient) string {
	return strings.TrimRight(c.ServiceURL(""), "/")
}

// resourceURL returns '.../d1/loadbalancers/{loadbalancerId}'.
// It is used for the Delete (DELETE) and Update (PUT) operations.
func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(id)
}

// serversURL returns '.../d1/loadbalancers/{loadbalancerId}/servers'.
// It is used for the AddServer (POST) and ListServers (GET) operations.
func serversURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(id, "servers")
}

// serverURL returns '.../d1/loadbalancers/{loadbalancerId}/servers/{serviceId}'.
// It is used for the RemoveServer (DELETE) operation.
func serverURL(c *gophercloud.ServiceClient, id, serviceID string) string {
	return c.ServiceURL(id, "servers", serviceID)
}

// tagURL returns '.../d1/loadbalancers/{loadbalancerId}/tag'.
// It is used for the CreateTag (PUT) and DeleteTag (DELETE) operations.
func tagURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(id, "tag")
}
