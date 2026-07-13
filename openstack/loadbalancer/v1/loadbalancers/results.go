// ### KT Cloud D1 platform > 'NLB (Network Load-Balancer)' handler Go SDK
// Open API Guide : https://cloud.kt.com/docs/open-api-guide/d/network/load-balancer

package loadbalancers

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

// LoadBalancer represents a KT Cloud (Platform @D) load balancer, as returned by
// the 'GET /loadbalancers' (list) API within the 'data' array.
type LoadBalancer struct {
	AccountID       string        `json:"accountId"`
	ZoneID          string        `json:"zoneId"`
	ZoneName        string        `json:"zoneName"`
	NetworkID       string        `json:"networkId"`
	NlbID           string        `json:"loadbalancerId"`
	Name            string        `json:"loadbalancerName"`
	DeviceID        string        `json:"deviceId"`
	IPPoolID        string        `json:"ippoolId"`
	ServiceType     string        `json:"serviceType"`
	ServiceIP       string        `json:"serviceIp"`
	ServicePort     int           `json:"servicePort"`
	NlbOption       string        `json:"loadbalancerOption"`
	Certificates    []interface{} `json:"certificates"`
	CipherGroupName string        `json:"cipherGroupName"`
	HealthCheckType string        `json:"healthCheckType"`
	HealthCheckURL  string        `json:"healthCheckUrl"`
	State           string        `json:"state"`
	IsWebSocket     bool          `json:"isWebSocket"`
	IsXForwardProto bool          `json:"isXForwardProto"`
	RedirectURL     string        `json:"redirectUrl"`
	Tag             string        `json:"tag"`
	EstablishedConn int           `json:"establishedConn"`
	RequestsRate    int           `json:"requestsRate"`
	TLSv1           string        `json:"tlsv1"`
	TLSv11          string        `json:"tlsv11"`
	TLSv12          string        `json:"tlsv12"`
	Comment         string        `json:"comment"`
}

// LbServer represents a server (VM) serviced by a load balancer. It corresponds to
// an item of the 'data.vm' array in the 'GET /loadbalancers/{loadbalancerId}/servers'
// response.
type LbServer struct {
	ServiceID           string `json:"webServiceId"`   // Load Balancer web service id
	Name                string `json:"name"`
	PortID              string `json:"portId"`
	VmID                string `json:"vmId"`
	ServiceType         string `json:"serviceType"` // HTTP | HTTPS
	IPAddress           string `json:"ipAddress"`
	PublicPort          int    `json:"publicPort"`
	MonitorState        string `json:"monitorState"`
	MonitorStatusMsg    string `json:"monitorStatusMessage"`
	VSvrServiceHitsRate int    `json:"vSvrServiceHitsRate"`
	AvgSvrTTFB          int    `json:"avgSvrTtfb"`
	CurCLntConnections  int    `json:"curCLntConnections"`
	RequestByteRate     int    `json:"requestByteRate"`
	ResponseByteRate    int    `json:"responseByteRate"`
}

// PaginationInfo represents the 'pagination' object included in list responses.
type PaginationInfo struct {
	Size   int `json:"size"`
	Total  int `json:"total"`
	Offset int `json:"offset"`
}

type commonResult struct {
	gophercloud.Result
}

// CreateResult represents the result of a Create operation.
// Call its ExtractCreate method to interpret it as a CreateNLBResponse.
type CreateResult struct {
	commonResult
}

// AddServerResult represents the result of an AddServer operation.
// Call its ExtractAddServer method to interpret it as an AddServerResponse.
type AddServerResult struct {
	commonResult
}

// RemoveServerResult represents the result of a RemoveServer operation.
type RemoveServerResult struct {
	commonResult
}

// CreateTagResult represents the result of a CreateTag operation.
type CreateTagResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// DeleteTagResult represents the result of a DeleteTag operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteTagResult struct {
	gophercloud.ErrResult
}

// CreateNLBResponse is the response body of the Create (POST /loadbalancers) API:
//
//	{ "data": { "loadbalancerId": "<uuid>" } }
type CreateNLBResponse struct {
	Data struct {
		LoadBalancerID string `json:"loadbalancerId"`
	} `json:"data"`
}

// AddServerResponse is the response body of the AddServer
// (POST /loadbalancers/{loadbalancerId}/servers) API:
//
//	{ "data": { "loadbalancerId": "<uuid>", "webServiceId": "<uuid>" } }
type AddServerResponse struct {
	Data struct {
		LoadBalancerID string `json:"loadbalancerId"`
		WebServiceID   string `json:"webServiceId"`
	} `json:"data"`
}

// Extract interprets any commonResult as a CreateNLBResponse.
func (r commonResult) Extract() (*CreateNLBResponse, error) {
	var s CreateNLBResponse
	err := r.ExtractInto(&s)
	return &s, err
}

// ExtractCreate interprets a CreateResult as a CreateNLBResponse.
func (r CreateResult) ExtractCreate() (*CreateNLBResponse, error) {
	var s CreateNLBResponse
	err := r.ExtractInto(&s)
	return &s, err
}

// ExtractAddServer interprets an AddServerResult as an AddServerResponse.
func (r AddServerResult) ExtractAddServer() (*AddServerResponse, error) {
	var s AddServerResponse
	err := r.ExtractInto(&s)
	return &s, err
}

// LoadBalancerPage is the page returned by a pager when traversing over a
// collection of load balancers.
type LoadBalancerPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a LoadBalancerPage struct is empty.
func (r LoadBalancerPage) IsEmpty() (bool, error) {
	is, err := ExtractLoadBalancers(r)
	return len(is) == 0, err
}

// NextPageURL returns the next page URL. The API uses 'page'/'size' query
// parameters rather than link-based pagination, so no next URL is provided.
func (r LoadBalancerPage) NextPageURL() (string, error) {
	return "", nil
}

// ServerPage is the page returned by a pager when traversing over a collection
// of servers within a load balancer.
type ServerPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a ServerPage struct is empty.
func (r ServerPage) IsEmpty() (bool, error) {
	servers, err := ExtractLbServers(r)
	return len(servers) == 0, err
}

// NextPageURL returns the next page URL. The servers API does not provide
// pagination links.
func (r ServerPage) NextPageURL() (string, error) {
	return "", nil
}

// ExtractLoadBalancers accepts a LoadBalancerPage and extracts the load balancers
// from the 'data' array.
func ExtractLoadBalancers(r pagination.Page) ([]LoadBalancer, error) {
	var s struct {
		Pagination PaginationInfo `json:"pagination"`
		Data       []LoadBalancer `json:"data"`
	}
	err := (r.(LoadBalancerPage)).ExtractInto(&s)
	return s.Data, err
}

// ExtractLbServers accepts a ServerPage and extracts the servers from the
// 'data.vm' array.
func ExtractLbServers(r pagination.Page) ([]LbServer, error) {
	var s struct {
		Data struct {
			VMs []LbServer `json:"vm"`
		} `json:"data"`
	}
	err := (r.(ServerPage)).ExtractInto(&s)
	return s.Data.VMs, err
}
