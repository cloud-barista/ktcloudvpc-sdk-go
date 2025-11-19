package staticnat

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

type StaticNAT struct {
	PublicIpID   string  `json:"publicIpId"`	// Public IP ID to set up Static NAT
	MappedIP     string  `json:"mappedIp"`
	SrcAddress   *string `json:"srcAddress"` 	// nullable field
	StaticNatID  string  `json:"staticNatId"`
	VpcID        string  `json:"vpcId"`
	CreateDate   string  `json:"createDate"`
	Name         string  `json:"name"`
	VlanName     string  `json:"vlanName"`
	PublicIP     string  `json:"publicIp"`		// Public IP to set up Static NAT
	NatSourceVip string  `json:"natSourceVip"`
}

type commonResult struct {
	gophercloud.Result
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a PortForwarding.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a PortForwarding.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a PortForwarding.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
// type DeleteResult struct {															// Modified
// 	commonResult
// }

// Extract will extract a Port Forwarding resource from a result.
func (r commonResult) Extract() (*StaticNAT, error) {
	var s StaticNAT
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "nc_enablestaticnatresponse")
}

type Pagination struct {															// Added
	Size   int `json:"size"`
	Total  int `json:"total"`
	Offset int `json:"offset"`
}

type CreateStaticNatResponse struct {												// Added
	HTTPStatus int          `json:"httpStatus"`
	Meta       interface{}  `json:"meta"` 		// nullable, use interface{} since it is unknown structure
	Pagination Pagination   `json:"pagination"`
	Data       []StaticNAT  `json:"data"`
}

func (r commonResult) ExtractInfo() (*CreateStaticNatResponse, error) {   			// Added
	var s CreateStaticNatResponse
	err := r.ExtractInto(&s)
	return &s, err
}

// type DellPortforwardingResponse struct {											// Added
// 	JopID string `json:"job_id"`
// }

// func (r commonResult) ExtractDelJobInfo() (*DellPortforwardingResponse, error) {   	// Added
// 	var s struct {
// 		DellPortforwardingResponse *DellPortforwardingResponse `json:"nc_deleteportforwardingruleresponse"`
// 	}
// 	err := r.ExtractInto(&s)
// 	return s.DellPortforwardingResponse, err
// }

// PortForwardingPage is the page returned by a pager when traversing over a
// collection of port forwardings.
type StaticNatPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of port forwardings has
// reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r StaticNatPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"static_net_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a PortForwardingPage struct is empty.
func (r StaticNatPage) IsEmpty() (bool, error) {
	is, err := ExtractStaticNats(r)
	return len(is) == 0, err
}

type SNATs struct {											  // Added
	StaticNATs []StaticNAT `json:"staticnats"`
}

// ExtractPortForwardings accepts a Page struct, specifically a PortForwardingPage
// struct, and extracts the elements into a slice of PortForwarding structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractStaticNats(r pagination.Page) ([]StaticNAT, error) {		// Added
	var s struct {
		SNAT SNATs `json:"nc_liststaticnatsresponse"`
	}
	err := (r.(StaticNatPage)).ExtractInto(&s)
	return s.SNAT.StaticNATs, err
}

// ExtractCreate extracts a CreateStaticNatResponse from a CreateResult.
func (r CreateResult) ExtractCreate() (*CreateStaticNatResponse, error) { // Modified
    var s CreateStaticNatResponse
    err := r.ExtractInto(&s)
    return &s, err
}
