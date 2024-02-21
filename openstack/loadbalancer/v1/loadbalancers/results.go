package loadbalancers 

import (
	"encoding/json"
	"time"
	// "fmt"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	// "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/loadbalancer/v2/listeners"
	// "github.com/cloud-barista/ktcloudvpc-sdk-for-drv/openstack/loadbalancer/v2/pools"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/pagination"
)

// LoadBalancer is the primary load balancing configuration object that
// specifies the virtual IP address on which client traffic is received, as well
// as other details such as the load balancing method to be use, protocol, etc.

type LoadBalancer struct {
	CertificateName    string `json:"certificatename"`
	CipherGroupName    string `json:"cipherGroupName"`
	ClientIpYn         string `json:"clientIpYn"`
	EstablishedConn    int    `json:"establishedconn"` // Caution!!) int
	HealthCheckType    string `json:"healthchecktype"` // Health CheckType : http / https / tcp
	HealthCheckURL     string `json:"healthcheckurl"`
	NlbID     		   int    `json:"loadbalancerid"`  // Caution!!) int
	NlbOption 		   string `json:"loadbalanceroption"`
	Name               string `json:"name"`
	NetworkID          string `json:"networkid"`
	RequestsRate       int 	`  json:"requestsrate"` // Caution!!) int
	ServiceIP          string `json:"serviceip"`
	ServicePort        string `json:"serviceport"`
	ServiceType        string `json:"servicetype"` // NLB Service Type : https / http / sslbridge / tcp / ftp
	SSLv2              string `json:"sslv2"`
	SSLv3              string `json:"sslv3"`
	State              string `json:"state"`
	Tag                string `json:"tag"`
	TLSv1              string `json:"tlsv1"`
	TLSv11             string `json:"tlsv11"`
	TLSv12             string `json:"tlsv12"`
	ZoneID             string `json:"zoneid"`
	ZoneName           string `json:"zonename"`
}

func (r *LoadBalancer) UnmarshalJSON(b []byte) error {
	type tmp LoadBalancer

	// Support for older neutron time format
	var s1 struct {
		tmp
		CreatedAt gophercloud.JSONRFC3339NoZ `json:"created_at"`
		UpdatedAt gophercloud.JSONRFC3339NoZ `json:"updated_at"`
	}

	err := json.Unmarshal(b, &s1)
	if err == nil {
		*r = LoadBalancer(s1.tmp)
		// r.CreatedAt = time.Time(s1.CreatedAt)
		// r.UpdatedAt = time.Time(s1.UpdatedAt)

		return nil
	}

	// Support for newer neutron time format
	var s2 struct {
		tmp
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	err = json.Unmarshal(b, &s2)
	if err != nil {
		return err
	}

	*r = LoadBalancer(s2.tmp)
	// r.CreatedAt = time.Time(s2.CreatedAt)
	// r.UpdatedAt = time.Time(s2.UpdatedAt)

	return nil
}

// StatusTree represents the status of a loadbalancer.
type StatusTree struct {
	Loadbalancer *LoadBalancer `json:"loadbalancer"`
}

type Stats struct {
	// The currently active connections.
	ActiveConnections int `json:"active_connections"`

	// The total bytes received.
	BytesIn int `json:"bytes_in"`

	// The total bytes sent.
	BytesOut int `json:"bytes_out"`

	// The total requests that were unable to be fulfilled.
	RequestErrors int `json:"request_errors"`

	// The total connections handled.
	TotalConnections int `json:"total_connections"`
}

// LoadBalancerPage is the page returned by a pager when traversing over a
// collection of load balancers.
type LoadBalancerPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of load balancers has
// reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r LoadBalancerPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"loadbalancers_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a LoadBalancerPage struct is empty.
func (r LoadBalancerPage) IsEmpty() (bool, error) {
	is, err := ExtractLoadBalancers(r)
	return len(is) == 0, err
}

// ExtractLoadBalancers accepts a Page struct, specifically a LoadbalancerPage
// struct, and extracts the elements into a slice of LoadBalancer structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractLoadBalancers(r pagination.Page) ([]LoadBalancer, error) { 		// Modified
	var s struct {
			ListLoadBalancersResponse struct {
			Count        	int `json:"count"`
			LoadBalancers 	[]LoadBalancer `json:"loadbalancer"`
		} `json:"listloadbalancersresponse"`
	}
	err := (r.(LoadBalancerPage)).ExtractInto(&s)
	return s.ListLoadBalancersResponse.LoadBalancers, err
}


// func ExtractLoadBalancers(r pagination.Page) ([]LoadBalancer, error) { 		// Modified
// 	// var s struct {
// 	// 	LoadBalancers []LoadBalancer `json:"loadbalancers"`
// 	// }

// 	var s struct {
// 			ListLoadBalancersResponse struct {
// 			Count        	int `json:"count"`
// 			LoadBalancers 	[]LoadBalancer `json:"loadbalancer"`
// 		} `json:"listloadbalancersresponse"`
// 	}
// 	err := ExtractLoadBalancersInto(r, &s.ListLoadBalancersResponse.LoadBalancers)
// 	return s.ListLoadBalancersResponse.LoadBalancers, err
// }

// func ExtractLoadBalancersInto(r pagination.Page, v interface{}) error { // Added
// 	return r.(LoadBalancerPage).Result.ExtractIntoSlicePtr(v, "loadbalancer")
// }

type commonResult struct {
	gophercloud.Result
}

type CreateNLBResponse struct {
	Createnlbresponse struct {
		NLBId    			string `json:"loadbalancerid"`
		ZoneId            	string `json:"zoneid"`
		ZoneName          	string `json:"zonename"`
		ServiceIP         	string `json:"serviceip"`
		ServicePort       	string `json:"serviceport"`
		ServiceType       	string `json:"servicetype"`
		Name              	string `json:"name"`
		NLBOption 			string `json:"loadbalanceroption"`
		HealthCheckType   	string `json:"healthchecktype"`
		HealthCheckURL    	string `json:"healthcheckurl"`
		ErrorCode    		string `json:"errorcode"`
		ErrorText    		string `json:"errortext"`

		DisplayText    		string `json:"displaytext"`		
	} `json:"createloadbalancerresponse"`
}
// Caution!!) Not `json:"createLoadBalancerresponse"` : API manul is incorrect!!


// Extract is a function that accepts a result and extracts a loadbalancer.
func (r commonResult) Extract() (*CreateNLBResponse, error) {
	var resp CreateNLBResponse

	err := r.ExtractInto(&resp.Createnlbresponse)
	return &resp, err
}
/*
// Extract is a function that accepts a result and extracts a loadbalancer.
func (r commonResult) Extract() (*LoadBalancer, error) {
	var s struct {
		LoadBalancer *LoadBalancer `json:"loadbalancer"`
	}
	err := r.ExtractInto(&s)
	return s.LoadBalancer, err
}
*/

// GetStatusesResult represents the result of a GetStatuses operation.
// Call its Extract method to interpret it as a StatusTree.
type GetStatusesResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts the status of
// a Loadbalancer.
func (r GetStatusesResult) Extract() (*StatusTree, error) {
	var s struct {
		Statuses *StatusTree `json:"statuses"`
	}
	err := r.ExtractInto(&s)
	return s.Statuses, err
}

// StatsResult represents the result of a GetStats operation.
// Call its Extract method to interpret it as a Stats.
type StatsResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts the status of
// a Loadbalancer.
func (r StatsResult) Extract() (*Stats, error) {
	var s struct {
		Stats *Stats `json:"stats"`
	}
	err := r.ExtractInto(&s)
	return s.Stats, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a LoadBalancer.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a LoadBalancer.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a LoadBalancer.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// FailoverResult represents the result of a failover operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type FailoverResult struct {
	gophercloud.ErrResult
}
