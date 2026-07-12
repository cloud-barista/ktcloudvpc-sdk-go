// ### KT Cloud D1 platform > 'NLB (Network Load-Balancer)' handler Go SDK
// Open API Guide : https://cloud.kt.com/docs/open-api-guide/d/network/load-balancer

package loadbalancers

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

// ############################## List ##############################

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToLoadBalancerListQuery() (string, error)
}

// ListOpts allows filtering of the paginated load balancer collection.
// GET /loadbalancers?loadbalancerName=&loadbalancerId=&page=&size=
type ListOpts struct {
	Name  string `q:"loadbalancerName"` // [Optional] Load Balancer name
	NlbID string `q:"loadbalancerId"`   // [Optional] Load Balancer id
	Page  int    `q:"page"`             // [Optional] [default: 1] page number
	Size  int    `q:"size"`             // [Optional] [default: 20] page size
}

// ToLoadBalancerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToLoadBalancerListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows iterating over a collection of load balancers.
//
//	GET /loadbalancers -> { "pagination": {...}, "data": [ {...} ] }
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToLoadBalancerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return LoadBalancerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// ############################## Create ##############################

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToLoadBalancerCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the request body for creating a Load Balancer.
//
//	POST /loadbalancers
type CreateOpts struct {
	Name            string `json:"loadbalancerName"`             // Required. Load Balancer name
	NetworkID       string `json:"networkId"`                    // Required. Tier network id
	ServiceIP       string `json:"serviceIp,omitempty"`          // [Optional] KT Cloud Virtual IP. Only when a new IP allocation is needed.
	ServicePort     int    `json:"servicePort"`                  // Required. Service port
	ServiceType     string `json:"serviceType"`                  // Required. HTTP | HTTPS | SSLBRIDGE | TCP | FTP | REDIRECT | DSR
	NlbOption       string `json:"loadbalancerOption,omitempty"` // [Optional] ROUNDROBIN | LEASTCONNECTION | LEASTRESPONSE | SOURCEIPHASH | SRCIPSRCPORTHASH
	HealthCheckType string `json:"healthCheckType,omitempty"`    // [Optional] HTTP | HTTPS | TCP
	HealthCheckURL  string `json:"healthCheckUrl,omitempty"`     // [Conditional] Required when healthCheckType is HTTP/HTTPS
	CipherGroupName string `json:"cipherGroupName,omitempty"`    // [Conditional] Required when serviceType is HTTPS
	TLSv1           string `json:"tlsv1,omitempty"`              // [Conditional] DISABLED | ENABLED (HTTPS)
	TLSv11          string `json:"tlsv11,omitempty"`             // [Conditional] DISABLED | ENABLED (HTTPS)
	TLSv12          string `json:"tlsv12,omitempty"`             // [Conditional] DISABLED | ENABLED (HTTPS)
	RedirectURL     string `json:"redirectUrl,omitempty"`        // [Conditional] Required when serviceType is REDIRECT
	IsWebSocket     bool   `json:"isWebSocket,omitempty"`        // [Optional] WebSocket support (true | false)
	IsxForwardProto bool   `json:"isXForwardProto,omitempty"`    // [Optional] Add X-Forwarded-Proto header (true | false)
	Comment         string `json:"comment,omitempty"`            // [Optional] Administrative comment
}

// ToLoadBalancerCreateMap builds the request body from CreateOpts.
func (opts CreateOpts) ToLoadBalancerCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create accepts a CreateOpts and uses the values provided to create a new Load Balancer.
//
//	POST /loadbalancers -> { "data": { "loadbalancerId": "<uuid>" } }
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToLoadBalancerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(rootURL(c), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ############################## Delete ##############################

// DeleteOpts identifies the Load Balancer to delete.
type DeleteOpts struct {
	NlbID string // Required. Load Balancer id (URL path)
}

// Delete permanently deletes a particular Load Balancer based on its unique id.
//
//	DELETE /loadbalancers/{loadbalancerId} -> 204 No Content
func Delete(c *gophercloud.ServiceClient, opts DeleteOpts) (r DeleteResult) {
	resp, err := c.Delete(resourceURL(c, opts.NlbID), &gophercloud.RequestOpts{
		OkCodes: []int{200, 204},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ############################## Add Server ##############################

// AddServerOpts is the request for adding a server (VM) to a Load Balancer.
//
//	POST /loadbalancers/{loadbalancerId}/servers
type AddServerOpts struct {
	NlbID       string `json:"-"`                // Required. Load Balancer id (URL path, not sent in the body)
	VMID        string `json:"vmId"`             // Required. VirtualMachine id
	IPAddress   string `json:"ipAddress"`        // Required. VirtualMachine ip
	PublicPort  string `json:"publicPort"`       // Required. Port
	ServiceType string `json:"serviceType"`      // Required. HTTP | HTTPS
	PortID      string `json:"portId,omitempty"` // [Conditional] Required when the Load Balancer serviceType is DSR
}

// ToLoadBalancerAddServerMap builds the request body from AddServerOpts.
func (opts AddServerOpts) ToLoadBalancerAddServerMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// AddServer adds a server (VM) to a Load Balancer.
//
//	POST /loadbalancers/{loadbalancerId}/servers
//	-> { "data": { "loadbalancerId": "<uuid>", "webServiceId": "<uuid>" } }
func AddServer(c *gophercloud.ServiceClient, opts AddServerOpts) (r AddServerResult) {
	b, err := opts.ToLoadBalancerAddServerMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(serversURL(c, opts.NlbID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ############################## Remove Server ##############################

// RemoveServerOpts identifies the server to remove from a Load Balancer.
type RemoveServerOpts struct {
	NlbID     string // Required. Load Balancer id (URL path)
	ServiceID string // Required. service id of the server to remove (URL path)
}

// RemoveServer removes a server (VM) from a Load Balancer.
//
//	DELETE /loadbalancers/{loadbalancerId}/servers/{serviceId} -> 204 No Content
func RemoveServer(c *gophercloud.ServiceClient, opts RemoveServerOpts) (r RemoveServerResult) {
	resp, err := c.Delete(serverURL(c, opts.NlbID, opts.ServiceID), &gophercloud.RequestOpts{
		OkCodes: []int{200, 204},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ############################## List Servers ##############################

// ListLbServer returns a Pager which allows iterating over the servers (VMs)
// of a Load Balancer.
//
//	GET /loadbalancers/{loadbalancerId}/servers -> { "data": { ..., "vm": [ {...} ] } }
func ListLbServer(c *gophercloud.ServiceClient, opts ListOpts) pagination.Pager {
	url := serversURL(c, opts.NlbID)
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return ServerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// ############################## Tag ##############################

// CreateTagOpts is the request for creating (registering) a Load Balancer tag.
//
//	PUT /loadbalancers/{loadbalancerId}/tag
type CreateTagOpts struct {
	NlbID string `json:"-"`   // Required. Load Balancer id (URL path, not sent in the body)
	Tag   string `json:"tag"` // Required. Tag content
}

// ToLoadBalancerCreateTagMap builds the request body from CreateTagOpts.
func (opts CreateTagOpts) ToLoadBalancerCreateTagMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// CreateTag creates (registers) a tag for a Load Balancer.
//
//	PUT /loadbalancers/{loadbalancerId}/tag -> 200 OK
func CreateTag(c *gophercloud.ServiceClient, opts CreateTagOpts) (r CreateTagResult) {
	b, err := opts.ToLoadBalancerCreateTagMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(tagURL(c, opts.NlbID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// DeleteTag deletes the tag of a Load Balancer.
//
//	DELETE /loadbalancers/{loadbalancerId}/tag -> 200 OK
func DeleteTag(c *gophercloud.ServiceClient, opts DeleteOpts) (r DeleteTagResult) {
	resp, err := c.Delete(tagURL(c, opts.NlbID), &gophercloud.RequestOpts{
		OkCodes: []int{200, 204},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
