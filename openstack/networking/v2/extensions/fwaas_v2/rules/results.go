package rules

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/pagination"
)

// Interface represents network interface information
type Interface struct {
	NetworkID string `json:"networkId"`
	Name      string `json:"name"`
	VlanName  string `json:"vlanName"`
}

// Address represents address information with name
type Address struct {
	Name string `json:"name"`
}

// Service represents service/port information
type Service struct {
	Protocol  string `json:"protocol"`
	StartPort string `json:"startPort"`
	EndPort   string `json:"endPort"`
}

// FirewallRule represents a firewall rule structure.
type FirewallRule struct {
	// PolicyId policy id
	PolicyID string `json:"policyId"`

	// Action allow or deny value
	Action string `json:"action"` // Not 'bool' type

	// Schedule firewall rule schedule time
	Schedule string `json:"schedule"`

	// Priority firewall rule priority order
	Priority int `json:"priority"`

	// SrcAddress src ip address information
	SrcAddress []Address `json:"srcAddress"`

	// DstAddress destination ip address information
	DstAddress []Address `json:"dstAddress"`

	// SrcInterface src network information
	SrcInterface []Interface `json:"srcInterface"`

	// DstInterface destination network information
	DstInterface []Interface `json:"dstInterface"`

	// Services destination port information
	Services []Service `json:"services"`

	// Comment rule description
	Comment string `json:"comment"`

	// Mkey fortigate schedule ID
	Mkey string `json:"mkey"`

	// Nat poolname creation value
	Nat string `json:"nat"`

	// VpcId vpc id
	VpcID string `json:"vpcId"`

	// Risk risk level
	Risk interface{} `json:"risk"`

	// IpPool poolname creation value
	IpPool string `json:"ipPool"`

	// PoolName creation CIDR value
	PoolName []interface{} `json:"poolName"`

	// Status rule status
	Status string `json:"status"`

	// CreateDate creation time
	CreateDate string `json:"createDate"`
}


type commonResult struct {
	gophercloud.Result
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a FirewallRule.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a FirewallRule.
type GetResult struct {
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
func (r commonResult) Extract() (*FirewallRule, error) {
	var s FirewallRule
	err := r.ExtractInto(&s)
	return &s, err
}


// func (r commonResult) ExtractInto(v interface{}) error {
// 	return r.Result.ExtractIntoStructPtr(v, "job_id")
// }


// RuleResponse represents the full response structure for P/F list.
type RuleResponse struct {                    // Modified
    HTTPStatus int              `json:"httpStatus"`
    Meta       interface{}      `json:"meta"`
    Pagination PaginationInfo   `json:"pagination"`
    Data       []FirewallRule `json:"data"`
}

// PaginationInfo represents pagination information in API responses.
type PaginationInfo struct {                // Added
    Size   int `json:"size"`
    Total  int `json:"total"`
    Offset int `json:"offset"`
}

// FirewallRulePage is the page returned by a pager when traversing over a
// collection of port forwardings.
type FirewallRulePage struct {
	pagination.LinkedPageBase
}

type CreateRuleJobResponse struct {
    HTTPStatus int    `json:"httpStatus"`
    JobID      string `json:"jobId"`
}

func ExtractJobID(result CreateResult) (string, error) {
    var resp CreateRuleJobResponse
    err := result.ExtractInto(&resp)
    if err != nil {
        return "", err
    }
    return resp.JobID, nil
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

// IsEmpty checks whether a FirewallRulePage struct is empty.
func (r FirewallRulePage) IsEmpty() (bool, error) {
	is, err := ExtractRules(r)
	return len(is) == 0, err
}

// NextPageURL returns the next page URL for traversing over VPC pages.
func (r FirewallRulePage) NextPageURL() (string, error) {
    // Pagination URL generation logic
    // Currently, the API does not provide pagination links, so manual implementation is required.
    return "", nil
}


// ExtractRules extracts 'a slice of FirewallRule' from a FirewallRulePage.
func ExtractRules(r pagination.Page) ([]FirewallRule, error) {
    var s RuleResponse
    err := r.(FirewallRulePage).ExtractInto(&s)
    return s.Data, err
}

// ExtractFirewallRule extracts 'a FirewallRule' from a GetResult.
func (r GetResult) ExtractFirewallRule() (*FirewallRule, error) { // Added
    var response RuleResponse
    err := r.ExtractInto(&response)
    if err != nil {
        return nil, err
    }
    
    if len(response.Data) == 0 {
        return nil, gophercloud.ErrDefault404{}
    }
    
    return &response.Data[0], nil
}

// ExtractRuleList extracts a RuleResponse from the result.
func (r commonResult) ExtractRuleList() (*RuleResponse, error) { // Added
    var s RuleResponse
    err := r.Result.ExtractInto(&s)
    return &s, err
}

// ExtractRuleResponse extracts the full RuleResponse from a result.
func (r commonResult) ExtractRuleResponse() (*RuleResponse, error) { // Added
    var s RuleResponse
    err := r.ExtractInto(&s)
    return &s, err
}

// ErrorResponse represents the error response structure for port forwarding API.
type ErrorResponse struct {
    HTTPStatus int    `json:"httpStatus"`
    Title      string `json:"title"`
    Detail     string `json:"detail"`
    Instance   string `json:"instance"`
    ServiceID  string `json:"serviceId"`
    Status     int    `json:"status"`
}

// ExtractErrorResponse extracts ErrorResponse from a CreateResult, UpdateResult, or commonResult.
func (r commonResult) ExtractErrorResponse() (*ErrorResponse, error) {
    var e ErrorResponse
    err := r.Result.ExtractInto(&e)
    if err != nil {
        return nil, err
    }
    return &e, nil
}

// For DeleteResult (which embeds ErrResult), add similar helper:
func (r DeleteResult) ExtractErrorResponse() (*ErrorResponse, error) {
    var e ErrorResponse
    err := r.ErrResult.ExtractInto(&e)
    if err != nil {
        return nil, err
    }
    return &e, nil
}
